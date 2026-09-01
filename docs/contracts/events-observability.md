# API conventions, events, audit, telemetry, and SLOs

## API conventions

The canonical API is REST/JSON described by OpenAPI 3.1. IDs are lowercase canonical UUIDv7 strings. Times are RFC 3339 UTC. Errors use `application/problem+json` (RFC 7807) with stable `type`, `title`, `status`, `detail`, `instance`, `request_id`, and optional field errors. Every response carries `X-Request-ID`. A caller-supplied value is propagated only when it is a canonical UUIDv7; otherwise MEM generates one. Request IDs are correlation, not identity or authority.

Collection pagination uses an authenticated opaque cursor and bounded `limit` (default 50, maximum 200). Mutation retries use `Idempotency-Key`; reuse with a different digest returns `409`. Requests accept/propagate W3C `traceparent` and `tracestate`. Default JSON body limit is 1 MiB; explicit ingestion deployments may configure at most 4 MiB. String/list limits are schema-defined. SSE uses `text/event-stream`, `id`, `event`, and JSON `data`, resumes with `Last-Event-ID`, sends heartbeats, and never embeds unbounded memory content.

## Event vocabulary

Event names are past tense and versioned in the envelope: `memory-space.created`, `memory.created`, `memory.revised`, `memory.disputed`, `memory.superseded`, `memory.quarantined`, `memory.expired`, `memory.deleted`, `ingestion.accepted`, `candidate.accepted`, `candidate.rejected`, `candidate.quarantined`, `context-pack.created`, `profile.rebuilt`, `forget.started`, `forget.completed`, `projection.rebuild-started`, `projection.rebuild-completed`.

The envelope has event ID/version/type, tenant, aggregate type/ID/revision, occurred/recorded times, actor/service, correlation IDs, classification, trace context, and a bounded payload/reference. Events are facts, not commands.

Audit actions additionally cover authentication failure, authorization allow/deny, grant invalid/revoked/expired, source spoof attempt, classification denial, legal-hold denial, administrative inspection, and policy changes. Audit records contain identifiers and reason codes; raw content, bearer tokens, grants, secrets, embeddings, and full prompts are forbidden.

## Telemetry and redaction

Logs use structured identifiers: tenant (hashed where required), MemorySpace/memory/revision/event/job/ContextPack IDs, Run/Session/Workspace IDs, request and trace IDs, operation, outcome, latency, reason code, and pinned policy/strategy versions. Content fields are denylisted from default logging. Any approved safe excerpt is length-bounded, classified, redacted, and disabled by default.

Metrics cover request rates/errors/latency, ingestion and candidate outcomes, worker queue/lease/retry/dead-letter state, memory lifecycle counts, contradictions/quarantine, retrieval latency and budgets, index lag/repair, profile builds, forget backlog, and dependency health. Labels must be bounded; never label by raw IDs or content. Traces record operation and dependency timing without payloads.

## Initial SLO and capacity assumptions

These are RC targets and must be revised using Phase 4 load evidence:

- API availability: 99.9% monthly, excluding declared maintenance.
- Authorized memory reads/writes: p95 250 ms, p99 750 ms, excluding asynchronous extraction/index completion.
- Context retrieval: p95 750 ms, p99 2 s with healthy dependencies and up to 50 returned items.
- Explicit write durability: PostgreSQL commit before success; 99% indexed within 60 s.
- Ingestion: 99% of accepted events reach terminal candidate state within 5 min when dependencies are healthy.
- Forget: canonical invisibility before acknowledgement; 99% projection cleanup within 15 min.
- Recovery point: zero acknowledged canonical transactions; projection RPO is zero after rebuild.

Reference capacity for initial tests is 100 tenants, 10 million canonical memories, 1,000 MemorySpaces per tenant, 100 writes/s sustained, 250 retrievals/s sustained, and 5x burst for one minute. Budgets assume 1 MiB public bodies, 4 MiB trusted ingestion, 200 list items, 50 ContextPack items, and configurable 16k estimated ContextPack tokens.
