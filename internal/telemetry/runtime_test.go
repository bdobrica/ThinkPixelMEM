package telemetry

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bdobrica/ThinkPixelMEM/internal/config"
	"go.opentelemetry.io/otel/propagation"
)

func TestNewInitializesIsolatedPrometheusAndW3CPropagation(t *testing.T) {
	t.Parallel()
	runtime, err := New(context.Background(), config.TelemetryConfig{Mode: "noop", ServiceName: "test-service"})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = runtime.Shutdown(context.Background()) })
	counter, err := runtime.MeterProvider.Meter("test").Int64Counter("jobs.completed")
	if err != nil {
		t.Fatal(err)
	}
	counter.Add(context.Background(), 2)
	recorder := httptest.NewRecorder()
	runtime.MetricsHandler().ServeHTTP(recorder, httptest.NewRequest("GET", "/metrics", nil))
	body := recorder.Body.String()
	if recorder.Code != 200 || !strings.Contains(body, "jobs_completed_total{") || !strings.Contains(body, "} 2") {
		t.Fatalf("unexpected metrics response %d: %s", recorder.Code, recorder.Body.String())
	}
	carrier := propagation.MapCarrier{"traceparent": "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01"}
	ctx := runtime.Propagator.Extract(context.Background(), carrier)
	out := propagation.MapCarrier{}
	runtime.Propagator.Inject(ctx, out)
	if out["traceparent"] != carrier["traceparent"] {
		t.Fatalf("trace context was not propagated: %q", out["traceparent"])
	}
}

func TestNewRejectsInvalidConfiguration(t *testing.T) {
	t.Parallel()
	if _, err := New(context.Background(), config.TelemetryConfig{Mode: "noop"}); err == nil {
		t.Fatal("expected invalid configuration to fail")
	}
	//lint:ignore SA1012 Deliberately exercise the public nil-context guard.
	if _, err := New(nil, config.TelemetryConfig{Mode: "noop", ServiceName: "test"}); err == nil {
		t.Fatal("expected nil context to fail")
	}
	if _, err := New(context.Background(), config.TelemetryConfig{Mode: "otlp", ServiceName: "test", Endpoint: "https://collector.test/custom"}); err == nil {
		t.Fatal("expected non-origin OTLP endpoint to fail")
	}
}

func TestShutdownIsIdempotent(t *testing.T) {
	t.Parallel()
	runtime, err := New(context.Background(), config.TelemetryConfig{Mode: "noop", ServiceName: "test"})
	if err != nil {
		t.Fatal(err)
	}
	if err := runtime.Shutdown(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := runtime.Shutdown(context.Background()); err != nil {
		t.Fatal(err)
	}
}
