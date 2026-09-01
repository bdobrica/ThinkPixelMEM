// Package telemetry initializes process-local metrics, tracing, and context
// propagation without using global registries or providers.
package telemetry

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"sync"

	"github.com/bdobrica/ThinkPixelMEM/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	tracenoop "go.opentelemetry.io/otel/trace/noop"
)

// Runtime owns the providers and Prometheus registry for one process. Call
// Shutdown during graceful process termination.
type Runtime struct {
	MeterProvider  *metric.MeterProvider
	TracerProvider trace.TracerProvider
	Propagator     propagation.TextMapPropagator
	Registry       *prometheus.Registry

	shutdownOnce sync.Once
	shutdownErr  error
	shutdown     func(context.Context) error
}

// New constructs local Prometheus metrics in every mode. OTLP mode additionally
// exports sampled traces over OTLP/HTTP; noop mode records no spans.
func New(ctx context.Context, cfg config.TelemetryConfig) (*Runtime, error) {
	if ctx == nil {
		return nil, errors.New("telemetry: nil context")
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(attribute.String("service.name", cfg.ServiceName)))
	if err != nil {
		return nil, err
	}
	registry := prometheus.NewRegistry()
	reader, err := otelprom.New(otelprom.WithRegisterer(registry))
	if err != nil {
		return nil, err
	}
	meterProvider := metric.NewMeterProvider(metric.WithReader(reader), metric.WithResource(res))

	var tracerProvider trace.TracerProvider = tracenoop.NewTracerProvider()
	shutdown := meterProvider.Shutdown
	if cfg.Mode == "otlp" {
		exporter, exportErr := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(strings.TrimRight(cfg.Endpoint, "/")+"/v1/traces"))
		if exportErr != nil {
			_ = meterProvider.Shutdown(ctx)
			return nil, exportErr
		}
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter), sdktrace.WithResource(res),
			sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(cfg.SampleRatio))),
		)
		tracerProvider = tp
		shutdown = func(ctx context.Context) error {
			return errors.Join(tp.Shutdown(ctx), meterProvider.Shutdown(ctx))
		}
	}
	return &Runtime{
		MeterProvider: meterProvider, TracerProvider: tracerProvider,
		Propagator: propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
		Registry:   registry, shutdown: shutdown,
	}, nil
}

// MetricsHandler returns a handler bound only to this runtime's registry.
func (r *Runtime) MetricsHandler() http.Handler {
	return promhttp.HandlerFor(r.Registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError})
}

// Shutdown flushes and releases exporters. It is safe to call more than once.
func (r *Runtime) Shutdown(ctx context.Context) error {
	if r == nil {
		return nil
	}
	r.shutdownOnce.Do(func() { r.shutdownErr = r.shutdown(ctx) })
	return r.shutdownErr
}
