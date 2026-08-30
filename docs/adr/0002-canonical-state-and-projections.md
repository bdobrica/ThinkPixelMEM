# ADR-0002: PostgreSQL canonical state and rebuildable projections

- Status: accepted
- Date: 2026-08-30
- Deciders: ThinkPixelMEM maintainers
- Supersedes: none

## Context

Vector, lexical, graph, and profile indexes improve retrieval but cannot provide the transaction, revision, provenance, and deletion guarantees required of canonical memory.

## Decision

PostgreSQL is the only canonical memory store. Qdrant is the RC dense-and-sparse retrieval projection. Graph indexes, profiles, summaries, embeddings, and caches are derived projections. Every projection record identifies tenant, MemorySpace, canonical memory ID, revision ID, and projection generation.

Projection writes originate from a transactional outbox. A projection can be deleted and rebuilt from non-deleted canonical rows. Search results are identifiers only until canonical authorization and state revalidation succeeds.

## Consequences

Index loss reduces availability or quality but never loses memory. Forget completes only after canonical tombstoning/deletion and durable cleanup work is recorded; rebuilding must not resurrect forgotten state.

## Alternatives considered

Qdrant as canonical state and a required graph database were rejected. Both weaken transactional history and increase the number of authorities.

## Verification

Phase 4 and 5 tests destroy and recreate projections, check canonical equivalence, and prove forgotten IDs cannot reappear.

