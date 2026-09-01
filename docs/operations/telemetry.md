# Metrics and tracing

`internal/telemetry` initializes a process-local Prometheus registry, an OpenTelemetry meter provider, W3C Trace Context and Baggage propagation, and tracing selected by `telemetry.mode`. It does not mutate OpenTelemetry or Prometheus global providers, so each API or worker process explicitly owns and shuts down its telemetry runtime.

Prometheus metrics are available from `Runtime.MetricsHandler`; the HTTP service will mount that handler at `/metrics` in ENG-008. Local Prometheus collection remains available in both modes. `noop` disables trace recording and export. `otlp` uses OTLP over HTTP, appends `/v1/traces` to the configured collector origin, batches spans, and applies parent-based trace-ID ratio sampling.

Metric labels and trace attributes must be bounded operational dimensions such as operation, outcome, memory type, state, dependency, and stable reason code. Raw IDs, tenant names, memory/source content, prompts, model output, tokens, grants, credentials, and unbounded error text must not be used as labels or attributes. Correlation identifiers belong in redaction-safe logs and trace context, not metric labels.

Call `Shutdown` with the process graceful-shutdown context to flush traces and metrics. Shutdown is idempotent.
