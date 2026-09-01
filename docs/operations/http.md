# HTTP service baseline

`internal/adapters/http` owns the standard-library HTTP server and transport middleware. Application routes remain injected so transport concerns do not enter the domain.

Every request receives a canonical UUIDv7 `X-Request-ID`, extracts W3C `traceparent` and `tracestate`, creates a server span, and records bounded request method/status-class metrics. Only canonical caller request IDs are propagated. Request IDs and trace context are correlation data and grant no authority.

JSON-capable request bodies are bounded by `http.max_body_bytes` (1 MiB by default and never more than 4 MiB). A known oversized body is rejected before dispatch; streamed bodies are wrapped with the same limit. Handlers should pass `*http.MaxBytesError` to `WriteProblem`. Typed domain errors map to stable RFC 7807 problem types. Problem details never contain underlying error strings, credentials, tokens, or request content. Unexpected panics become generic `500` problems.

Operational endpoints are:

- `GET /livez`: process responsiveness; it does not check dependencies.
- `GET /readyz`: injected dependency readiness; absence or failure returns `503` and fails closed.
- `GET /metrics`: the process-local Prometheus registry.

`Server.Run` listens on `http.address`. Canceling its context stops acceptance and drains in-flight requests within `http.shutdown_timeout`; after that deadline it closes remaining connections. The process owner should then shut down telemetry using the same bounded termination phase.
