package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const RequestIDHeader = "X-Request-ID"

type requestIDKey struct{}

func RequestID(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey{}).(string)
	return value
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	if w.status != 0 {
		return
	}
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
func (w *statusWriter) Write(p []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(200)
	}
	return w.ResponseWriter.Write(p)
}
func (w *statusWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

type middleware struct {
	ids        *domain.IDGenerator
	propagator propagation.TextMapPropagator
	tracer     trace.Tracer
	maxBody    int64
	requests   *prometheus.CounterVec
	duration   *prometheus.HistogramVec
}

func (m *middleware) wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := m.requestID(r)
		if err != nil {
			WriteProblem(w, r, err)
			return
		}
		ctx := context.WithValue(r.Context(), requestIDKey{}, id)
		r = r.WithContext(ctx)
		w.Header().Set(RequestIDHeader, id)
		ctx = m.propagator.Extract(ctx, propagation.HeaderCarrier(r.Header))
		ctx, span := m.tracer.Start(ctx, "HTTP "+r.Method, trace.WithSpanKind(trace.SpanKindServer), trace.WithAttributes(attribute.String("http.request.method", r.Method)))
		defer span.End()
		r = r.WithContext(ctx)
		r.Body = http.MaxBytesReader(w, r.Body, m.maxBody)
		m.propagator.Inject(ctx, propagation.HeaderCarrier(w.Header()))
		sw := &statusWriter{ResponseWriter: w}
		started := time.Now()
		defer func() {
			if recover() != nil {
				span.SetStatus(codes.Error, "panic")
				WriteProblem(sw, r, domain.NewError(domain.CodeInternal, "handler panic", nil))
			}
			status := sw.status
			if status == 0 {
				status = 200
			}
			class := strconv.Itoa(status/100) + "xx"
			method := methodLabel(r.Method)
			m.requests.WithLabelValues(method, class).Inc()
			m.duration.WithLabelValues(method, class).Observe(time.Since(started).Seconds())
			span.SetAttributes(attribute.Int("http.response.status_code", status))
			if status >= 500 {
				span.SetStatus(codes.Error, http.StatusText(status))
			}
		}()
		if r.ContentLength > m.maxBody {
			WriteProblem(sw, r, &http.MaxBytesError{Limit: m.maxBody})
			return
		}
		next.ServeHTTP(sw, r)
	})
}

func methodLabel(method string) string {
	switch method {
	case "GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS":
		return method
	default:
		return "OTHER"
	}
}
func (m *middleware) requestID(r *http.Request) (string, error) {
	if supplied := r.Header.Get(RequestIDHeader); supplied != "" {
		if _, err := domain.ParseUUID(supplied); err == nil {
			return supplied, nil
		}
	}
	id, err := m.ids.New()
	if err != nil {
		return "", fmt.Errorf("generate request ID: %w", err)
	}
	return id.String(), nil
}
