# PostgreSQL, jobs, and outbox contract

## Relational model

Every tenant-owned table carries `tenant_id`; every space-owned table also carries `memory_space_id`. Foreign keys include tenant identity so cross-tenant references are structurally impossible. UUIDv7 primary keys are application-generated and non-enumerating.

Canonical tables are `tenants`, `memory_spaces`, `entities`, `episodes`, `claims`, `memory_revisions`, `relationships`, `outcomes`, `lessons`, `procedure_candidates`, `evidence_references`, link tables, `profile_schemas`, `profiles`, `profile_fields`, `ingestion_events`, `memory_candidates`, `jobs`, `idempotency_records`, `context_packs`, `retrieval_uses`, `audit_events`, `outbox_messages`, `forget_operations`, and `projection_generations`.

Required constraints include unique ingestion external IDs, unique `(logical_memory_id, revision)`, exactly one valid active revision pointer, valid half-open intervals, confidence range, immutable completed revisions/evidence, no cross-tenant relationships, and unique idempotency key per tenant/operation. Released migrations are immutable.

Application transactions set tenant context and always predicate by tenant. Database row-level security is defense in depth, not a replacement for application authorization. Sensitive JSON is minimized; fields used for isolation, lifecycle, and policy are typed columns.

## Jobs, leases, and fencing

A job stores ID, tenant, kind, logical source ID, strategy version, state, attempt, next-attempt time, lease owner, lease expiry, monotonically increasing fence token, timestamps, bounded error class/detail, and idempotency key. States are `pending`, `leased`, `succeeded`, `retryable`, `dead-letter`, `cancelled`.

Claiming atomically increments the fence. Workers may commit side effects only while their fence is current. Lease expiry permits replay. Operations are idempotent at the destination. Backoff is bounded exponential with jitter and a policy-specific maximum attempt count. Semantic/validation failures dead-letter without blind retry.

## Transactional outbox

Canonical mutation and an `outbox_messages` row commit in one PostgreSQL transaction. Each message contains aggregate ID/revision, tenant, event type/version, payload or reference, ordering key, created time, attempt state, and unique logical message ID. A dispatcher publishes at least once; consumers deduplicate. Publication order is guaranteed only per ordering key. Cleanup occurs after a retention window, never before durable acknowledgement.

## Representative invariants

- A revision insert and active-pointer change are atomic under optimistic concurrency.
- Accepted candidate consolidation and emitted index/audit events are atomic.
- Forget visibility change and all cleanup work records are atomic.
- Projection state never drives canonical transitions.
- Audit is append-only and excludes raw content by default.

