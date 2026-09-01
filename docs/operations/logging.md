# Structured logging

ThinkPixelMEM emits JSON logs through `internal/telemetry/logging`. The configured minimum level is one of `debug`, `info`, `warn`, or `error`.

Operational messages are stable descriptions rather than runtime data. Correlation is carried in structured fields: `tenant`, `memory_space_id`, `memory_id`, `revision_id`, `ingestion_event_id`, `extraction_job_id`, `context_pack_id`, `run_id`, `session_id`, `workspace_id`, `request_id`, and `trace_id`. Callers hash tenant identifiers before logging when deployment policy requires it. Empty identifiers are omitted.

Operations may add bounded fields such as `operation`, `outcome`, `latency_ms`, `reason_code`, and pinned policy or strategy versions. Raw identifiers must not become metric labels when metrics are added.

The logging handler drops fields named for content, prompts, completions, model output, payloads, embeddings, excerpts, request/response bodies, memory text, grants, raw errors, headers, or cookies. Operations emit a stable `reason_code` instead of an arbitrary error string. Fields named for credentials, authorization, tokens, keys, passwords, database URLs, DSNs, or secrets retain their key but replace their value with `[REDACTED]`. Matching is case-insensitive after normalizing dots and hyphens. The rules also apply within structured groups.

Redaction is a last line of defense, not permission to submit sensitive values. Log messages must never contain runtime memory, prompts, model output, tokens, credentials, or secret material. Approved safe excerpts remain disabled until a later policy-backed implementation supplies classification, bounding, and redaction.
