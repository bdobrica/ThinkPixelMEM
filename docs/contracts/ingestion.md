# Ingestion, candidate, extraction, and consolidation contract

## IngestionEvent

An event envelope contains `id`, `tenant_id`, `memory_space_id`, `external_event_id`, `source_service`, `source_kind`, `source_trust`, `classification`, `residency`, actor/principal, optional Run/Session/Workspace IDs, evidence refs, bounded content/reference, `occurred_at`, server `received_at`, and schema version.

`(tenant_id, source_service, external_event_id)` is unique. Replays return the original accepted result and never schedule duplicate logical work. A payload digest mismatch for the same key is a conflict and security audit event. Ingestion authenticates the source before assigning kind/trust.

## MemoryCandidate

A candidate contains an ID, ingestion event ID, proposed memory type and normalized content, derived entities/topics/relationships, confidence, importance, extractor strategy/version and LLMGW request reference, evidence IDs inherited from the event, risk flags, and state.

States are `proposed`, `validating`, `accepted`, `rejected`, `quarantined`, `consolidated`. State changes are audited. Extractor output cannot contain authoritative fields; if present, the candidate is rejected rather than silently trusted.

## Validation pipeline

Validation is ordered and deterministic where possible:

1. Validate schema, sizes, enums, and typed values.
2. Attach authoritative event metadata and reject conflicts.
3. Enforce MemorySpace write authority, classification, residency, retention, and allowed types.
4. Apply sensitive-inference and secret policy.
5. Inspect instruction/poison risk locally and, when required, through GR.
6. Apply strategy-specific confidence threshold without changing source trust.
7. Detect exact/idempotent duplicates.
8. Match compatible claims and relationships.
9. Detect contradiction or temporal supersession.
10. Accept, reject with reason, or quarantine; accepted candidates enter consolidation.

## MemoryExtractor port

```go
type MemoryExtractor interface {
    Extract(context.Context, ExtractionRequest) ([]MemoryCandidate, error)
}
```

`ExtractionRequest` contains bounded source content, inherited evidence, allowed output types, schema/strategy version, and a stable logical operation ID. The adapter records its own ID/version and the LLMGW request reference. Provider credentials never enter MEM.

Extraction is asynchronous by default. Explicit `remember` may synchronously normalize a direct assertion but indexing remains asynchronous. Failed work is retryable only according to the job contract.

## MemoryStrategy

A versioned strategy defines enabled memory types, field constraints, confidence thresholds, sensitive-inference mode (`deny`, `review`, `allow`), retention by type, duplicate/matching thresholds, corroboration requirements, poison handling, embedding/reranking versions, and GR failure behavior. Existing records pin the exact strategy version.

## Consolidator port and rules

```go
type Consolidator interface {
    Consolidate(context.Context, ConsolidationRequest) (ConsolidationResult, error)
}
```

Deterministic matching first uses tenant, MemorySpace, normalized subject/entity, predicate, value type, and overlapping validity. Compatible repeated evidence appends a revision/evidence link; it does not create an unbounded duplicate. Conflicting values with overlapping validity create a dispute unless a trusted temporal rule proves supersession. Non-overlapping successor facts close/link intervals. LLM assistance may propose a match but trusted application code validates every transition.

