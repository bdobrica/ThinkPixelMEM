package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	stdhttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bdobrica/ThinkPixelMEM/internal/config"
	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
	clockport "github.com/bdobrica/ThinkPixelMEM/internal/ports/clock"
	"github.com/bdobrica/ThinkPixelMEM/internal/telemetry"
)

type fixedClock struct{ at time.Time }

func (c fixedClock) Now() time.Time { return c.at }

type memoryListener struct{ closed chan struct{} }

func (l *memoryListener) Accept() (net.Conn, error) { <-l.closed; return nil, net.ErrClosed }
func (l *memoryListener) Close() error {
	select {
	case <-l.closed:
	default:
		close(l.closed)
	}
	return nil
}
func (l *memoryListener) Addr() net.Addr { return memoryAddr("memory") }

type memoryAddr string

func (a memoryAddr) Network() string { return string(a) }
func (a memoryAddr) String() string  { return string(a) }

func testServer(t *testing.T, ready ReadinessCheck, app stdhttp.Handler) *Server {
	t.Helper()
	runtime, err := telemetry.New(context.Background(), config.TelemetryConfig{Mode: "noop", ServiceName: "http-test"})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = runtime.Shutdown(context.Background()) })
	ids, err := domain.NewIDGenerator(clockport.Clock(fixedClock{time.UnixMilli(1720000000000)}), bytes.NewReader(make([]byte, 10000)))
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.Defaults().HTTP
	cfg.MaxBodyBytes = 8
	server, err := NewServer(cfg, runtime, ids, ready, app)
	if err != nil {
		t.Fatal(err)
	}
	return server
}

func TestOperationalEndpointsAndRequestCorrelation(t *testing.T) {
	server := testServer(t, func(context.Context) error { return nil }, nil)
	for _, path := range []string{"/livez", "/readyz", "/metrics"} {
		recorder := httptest.NewRecorder()
		server.Handler().ServeHTTP(recorder, httptest.NewRequest("GET", path, nil))
		if recorder.Code != 200 {
			t.Fatalf("%s returned %d: %s", path, recorder.Code, recorder.Body.String())
		}
		if _, err := domain.ParseUUID(recorder.Header().Get(RequestIDHeader)); err != nil {
			t.Fatalf("%s request ID: %v", path, err)
		}
	}
}

func TestReadinessFailsClosedAsProblem(t *testing.T) {
	server := testServer(t, func(context.Context) error { return errors.New("database password=secret") }, nil)
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, httptest.NewRequest("GET", "/readyz", nil))
	if recorder.Code != 503 || recorder.Header().Get("Content-Type") != "application/problem+json" || strings.Contains(recorder.Body.String(), "password") {
		t.Fatalf("unsafe readiness response: %d %s", recorder.Code, recorder.Body.String())
	}
	var problem Problem
	if err := json.Unmarshal(recorder.Body.Bytes(), &problem); err != nil {
		t.Fatal(err)
	}
	if problem.RequestID == "" || problem.Type != "https://thinkpixel.dev/problems/not-ready" {
		t.Fatalf("unexpected problem: %+v", problem)
	}
}

func TestBodyLimitAndTypedProblemMapping(t *testing.T) {
	app := stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		_, err := io.ReadAll(r.Body)
		if err != nil {
			WriteProblem(w, r, err)
			return
		}
		WriteProblem(w, r, domain.NewError(domain.CodeForbidden, "secret detail", errors.New("token=secret")))
	})
	server := testServer(t, func(context.Context) error { return nil }, app)
	for _, body := range []io.Reader{strings.NewReader("123456789"), io.NopCloser(strings.NewReader("123456789"))} {
		req := httptest.NewRequest("POST", "/v1/test", body)
		if _, ok := body.(io.ReadCloser); ok {
			req.ContentLength = -1
		}
		recorder := httptest.NewRecorder()
		server.Handler().ServeHTTP(recorder, req)
		if recorder.Code != 413 || strings.Contains(recorder.Body.String(), "secret") {
			t.Fatalf("unexpected limited response: %d %s", recorder.Code, recorder.Body.String())
		}
	}
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, httptest.NewRequest("POST", "/v1/test", strings.NewReader("ok")))
	if recorder.Code != 403 || strings.Contains(recorder.Body.String(), "secret") {
		t.Fatalf("unsafe mapped problem: %d %s", recorder.Code, recorder.Body.String())
	}
}

func TestTracePropagationMetricsAndPanicContainment(t *testing.T) {
	app := stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.URL.Path == "/panic" {
			panic("sensitive")
		}
		w.WriteHeader(204)
	})
	server := testServer(t, func(context.Context) error { return nil }, app)
	req := httptest.NewRequest("GET", "/v1/test", nil)
	req.Header.Set("traceparent", "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, req)
	if recorder.Code != 204 || recorder.Header().Get("traceparent") != req.Header.Get("traceparent") {
		t.Fatalf("trace response: %d %q", recorder.Code, recorder.Header().Get("traceparent"))
	}
	panicRecorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(panicRecorder, httptest.NewRequest("GET", "/panic", nil))
	if panicRecorder.Code != 500 || strings.Contains(panicRecorder.Body.String(), "sensitive") {
		t.Fatalf("panic response: %d %s", panicRecorder.Code, panicRecorder.Body.String())
	}
	metrics := httptest.NewRecorder()
	server.Handler().ServeHTTP(metrics, httptest.NewRequest("GET", "/metrics", nil))
	if !strings.Contains(metrics.Body.String(), `thinkpixelmem_http_requests_total{method="GET",status_class="2xx"}`) {
		t.Fatalf("metrics missing: %s", metrics.Body.String())
	}
}

func TestRunStopsAfterContextCancellation(t *testing.T) {
	runtime, err := telemetry.New(context.Background(), config.TelemetryConfig{Mode: "noop", ServiceName: "run-test"})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = runtime.Shutdown(context.Background()) })
	ids, err := domain.NewIDGenerator(fixedClock{time.UnixMilli(1720000000000)}, bytes.NewReader(make([]byte, 100)))
	if err != nil {
		t.Fatal(err)
	}
	cfg := config.Defaults().HTTP
	cfg.Address = "127.0.0.1:0"
	server, err := NewServer(cfg, runtime, ids, func(context.Context) error { return nil }, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan error, 1)
	go func() { done <- server.run(ctx, &memoryListener{closed: make(chan struct{})}) }()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("server did not shut down")
	}
}
