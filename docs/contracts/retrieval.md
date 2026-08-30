# Retrieval and ContextPack contract

## Retrieval request

Inputs include an explicit `schema_version` (initially `thinkpixel.mem.retrieval-request.v1`), tenant/caller identity, MemoryGrant or standalone authorization, query/goal, allowed MemorySpaces, subjects/entities, memory types, classification ceiling, `as_of`/validity filters, minimum trust/confidence, item and token budgets, cursor, and debug flag.

Authorization intersects—not unions—the request with the grant. Scope, tenant, type, classification, status, deletion, and temporal rules are enforced before query and rechecked against PostgreSQL after index candidate retrieval.

## RetrievalIndex port

```go
type RetrievalIndex interface {
    Search(context.Context, SearchQuery) ([]CandidateID, error)
    Upsert(context.Context, []IndexRecord) error
    Delete(context.Context, []MemoryID) error
    Rebuild(context.Context, CanonicalMemorySource) error
}
```

Qdrant is the RC adapter. Each point has named dense and sparse vectors plus tenant, space, type, classification, status, revision, validity, embedding version, and projection generation payload. Hybrid retrieval uses reciprocal-rank fusion initially. Graph relationships remain canonical domain objects; a graph database is deferred.

## Scoring

The versioned policy combines semantic, sparse/lexical, entity, goal, freshness, importance, evidence quality, source trust, confidence, and temporal signals, then subtracts dispute, stale, redundancy, and poison-risk penalties. Hard authorization filters are never scoring signals. Debug output returns component scores and policy version only to authorized callers.

An inference may score highly but remains labeled. Low trust, disputed, stale, or poison-risk items carry warnings; policy may exclude them. Deduplication groups logical identities and near-equivalent revisions before budgeting.

## ContextPack

A ContextPack contains an explicit `schema_version` (initially `thinkpixel.mem.context-pack.v1`), `id`, tenant, caller/Run reference, retrieval policy version, created/expiry timestamps, applied scopes/filters, budgets and usage, ordered items, warnings, and bounded evidence metadata. Consumers reject unsupported major schema versions; additive compatible evolution remains within the declared major version.

Every item contains memory ID/type, revision, normalized content or safe summary, status, confidence, source trust/kind, classification, temporal validity, evidence references, retrieval reason, component scores when authorized, and warnings. It never contains a capability, system instruction, or policy grant.

Budgeting is deterministic: exclude unauthorized/invalid items, group contradictions, deduplicate, order, then admit whole items until both item and token estimates fit. An oversized item is skipped with a warning. Persist only IDs, revisions, policy version, scores, warnings, and budget accounting by default—not full source content.

## Embedding migration

Embedding metadata includes provider-neutral model ID, dimensions, normalization, strategy version, created time, and projection generation. A model change creates a new generation, dual-indexes during migration, atomically switches the active generation, and later deletes the old generation. Canonical memory identity/revision never changes.
