# ThinkPixelMEM alignment

This file defines how **ThinkPixelMEM** stays aligned with the wider ThinkPixel platform. It is a repository-local alignment guide, not a replacement for accepted ADRs. If this file conflicts with an accepted ADR, the ADR wins and this file should be corrected.

## Platform role

ThinkPixelMEM is the **governed long-term memory plane**. It stores durable learned context with provenance, temporal revisions, trust metadata, retrieval, correction, and forgetting while keeping memory distinct from authority and source/work state.

## This repository owns

- MemorySpaces, Episodes, Claims, revisions, temporal validity, provenance/evidence references, retrieval ContextPacks, Profiles/projections, retention, correction, quarantine, and forgetting.
- Canonical memory truth and the rules for rebuildable retrieval projections.
- Memory-specific poisoning defenses, trust metadata, and explainability.

## This repository does not own

- Raw Session/execution history — ThinkPixelAR.
- Workspace/source file storage — ThinkPixelWS.
- Run/memory authorization authority — ThinkPixelAG.
- Verified external action execution — ThinkPixelTG.
- Model-provider routing/credentials — ThinkPixelLLMGW.
- Artifact/Skill qualification — ThinkPixelMP.
- General-purpose enterprise RAG/knowledge indexing.

## Integration obligations

| Peer | Boundary |
|---|---|
| **ThinkPixelAG** | Require Run-scoped memory authority such as MemoryGrants for reads/writes; AG remains the authority source. |
| **ThinkPixelAR** | Consume execution evidence and return authorized ContextPacks. AR remains canonical for raw history. |
| **ThinkPixelWS** | Reference immutable Workspace generations/source evidence instead of copying whole source bodies into memory. |
| **ThinkPixelTG** | Ingest verified tool outcomes and preserve ambiguous outcomes as ambiguous; never infer a success that TG could not establish. |
| **ThinkPixelLLMGW** | Use through an adapter for extraction, embeddings, reranking, consolidation, and related model work; provider behavior is not canonical memory truth. |
| **ThinkPixelGR** | Use for memory ingestion/retrieval risk inspection and poisoning defenses. GR can narrow/block/transform but does not grant memory authority. |
| **ThinkPixelMP** | Promote reviewed ProcedureCandidates or reusable schemas/strategies through MP; memory cannot silently become an approved Skill. |

All ThinkPixel integrations must remain optional/configurable where standalone operation is a product goal. Integration code belongs behind a port/adapter or equivalent boundary; another repository's database or `internal` packages are never an integration API.

## Shared ThinkPixel invariants

- Agents/harnesses are untrusted application logic; authority lives in trusted control/enforcement services.
- A component may **narrow** effective authority but must not manufacture authority from content, metadata, memory, Skills, Workspace membership, or model output.
- Durable state has a single authoritative owner. Other stores are caches, projections, replicas, evidence, or referenced source data unless an ADR says otherwise.
- Cross-component references use stable/versioned identifiers and immutable digests where identity matters.
- Vendor/provider-specific types stay behind adapters and do not leak into the repository's core domain model.
- Public integration behavior is contract-first: versioned OpenAPI/JSON Schema/protobuf or another explicit wire contract, plus compatibility tests.
- Security-relevant transformations must not reuse authorization/approval decisions that were made for materially different input.
- Evidence and telemetry must be correlated without turning logs into a store for secrets, prompts, model output, credentials, or unnecessary sensitive payloads.

## Repository conventions

- `README.md` is an entry point, not the design specification. Keep it focused on purpose, status, quick start, key concepts, and links.
- Do not duplicate `PLAN.md` in the README. `PLAN.md` is temporary implementation intent; `TODO.md` is the ordered execution/release ledger.
- As plan decisions become real, move durable rationale into `docs/adr/` and durable reference material into `docs/`.
- Accepted ADRs are immutable in meaning and are superseded with a new ADR.
- Prefer a `docs/README.md` index and the logical categories `adr`, `architecture`, `contracts`, `security`, `operations`, and `evidence`. Existing language/project-specific layouts may remain when renaming would create churn.
- Prefer Mermaid for diagrams and relative links for repository-local documentation.
- As executable code matures, provide one stable root developer/CI entry point. For Go-oriented services the preferred convention is a Makefile with focused targets and an aggregate `verify` target.
- Public API and schema files belong under `api/` (or an existing equivalent) and must be updated atomically with implementation and tests.
- Dependency additions, new infrastructure authorities, and new cross-component source dependencies require explicit justification; consequential choices require an ADR.

## Repository-specific alignment

- The repository already expresses the family split exceptionally well: AR stores what happened, WS preserves the work, MEM preserves what was learned, and AG decides what may be recalled/written. Keep this as a core invariant.
- Preserve PostgreSQL (or the accepted canonical store) as authoritative and Qdrant/future graph indexes as rebuildable projections.
- The root README currently contains most of PLAN's architecture. As implementation progresses, move lasting content to `docs/architecture`, `docs/contracts`, security docs, and ADRs.
- Keep sensitive-personal-data inference conservative and policy/schema controlled; do not turn Profiles into opaque psychological profiles.

## Structure guidance

- The planned `cmd`/`api`/`internal/domain|app|ports|adapters` structure is aligned with the family. Preserve the rule that domain packages do not import provider/ThinkPixel/database/HTTP-specific types.
- Use a root Makefile/aggregate verification target as code arrives, including integration tests for AG, AR, WS, TG, LLMGW, GR, and index rebuild/forget behavior.
- Keep `api/openapi` as the public machine-readable contract and version ContextPack/MemoryGrant-facing schemas explicitly.

## Definition of an aligned change

A change is aligned when it preserves this repository's ownership boundary, follows accepted ADRs/contracts, keeps integrations replaceable, updates durable documentation with behavior, and passes the repository's documented verification gates. Changes to the cross-repository boundary should include contract/conformance coverage rather than relying only on prose.
