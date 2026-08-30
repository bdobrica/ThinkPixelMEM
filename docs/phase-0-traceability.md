# Phase 0 traceability and decisions

This matrix maps every Phase 0 checklist item to its normative artifact. A grouped row applies to every inclusive item in the range.

| TODO item(s) | Artifact and decision |
| --- | --- |
| ARC-001 | `docs/adr/template.md`, `docs/adr/README.md`, and `docs/contracts/` |
| ARC-002–003 | `docs/architecture.md` system-context and trust-boundary Mermaid diagrams |
| ARC-004 | `docs/threat-model.md` |
| ARC-005–013 | ADR-0001 through ADR-0004; authority, observation/inference, trust/confidence, authoritative metadata, and projection invariants |
| ARC-014–016 | `contracts/domain.md` MemorySpace identity, six scope types, ownership, classification, residency, retention, and read/write policy |
| ARC-017–024 | `contracts/domain.md` Episode, Claim, Relationship, Outcome, Lesson, ProcedureCandidate, Profile, and Entity contracts |
| ARC-025–028 | `contracts/domain.md` source-kind, source-trust, confidence, and ordered classification vocabularies |
| ARC-029–034 | ADR-0004 and `contracts/lifecycle.md` temporal fields, seven statuses, immutable revisions, correction, dispute, and supersession |
| ARC-035–040 | `contracts/domain.md` EvidenceReference schema and canonical AR/WS/TG/user/import URI formats |
| ARC-041 | ADR-0003 and `contracts/domain.md` authoritative/derived field classes |
| ARC-042–050 | `contracts/ingestion.md` event idempotency, candidate schema/states, ordered validation, async extraction, interfaces, strategies, consolidation, deduplication, and contradiction rules |
| ARC-051–055 | `contracts/retrieval.md` ContextPack, retrieval request/scoring/warnings, and deterministic dual budgeting |
| ARC-056–060 | `contracts/retrieval.md` RetrievalIndex, Qdrant 1.19 dense+sparse RRF, embedding generations, lexical strategy, and deferred graph DB |
| ARC-061–064 | `contracts/integrations.md` AG MemoryGrant permissions/limits, expiry/revocation, and deny-by-default standalone authorizer |
| ARC-065–068 | `contracts/integrations.md` LLMGW, GR write/retrieval inspection, and MP promotion boundary |
| ARC-069–073 | `contracts/lifecycle.md` retention/TTL, five forget selectors, legal hold, privacy/compliance revision deletion, and projection rebuild guarantees |
| ARC-074–077 | `contracts/domain.md`, `contracts/retrieval.md`, `contracts/ingestion.md`, and `threat-model.md` profile explainability, sensitive inference, poisoning/quarantine, and corroboration seam |
| ARC-078–080 | `contracts/persistence.md` schema/invariants, leased fenced jobs, and transactional outbox |
| ARC-081–082 | `contracts/integrations.md` OIDC identity/tenant mapping requirements and admin authorization versus runtime MemoryGrant |
| ARC-083–084 | `api/openapi/openapi.yaml` and `contracts/events-observability.md` OpenAPI, RFC 7807, UUIDv7, pagination, idempotency, tracing, limits, and SSE |
| ARC-085–087 | `contracts/events-observability.md` audit/event vocabulary, redaction-safe telemetry, initial SLOs and capacity |
| ARC-088 | `docs/supported-versions.md` |
| ARC-089 | `docs/secondcontext-comparison.md` |
| ARC-090 | `docs/phase-0-validation.md` and `scripts/validate-phase0.sh`; commit evidence intentionally remains pending until a human/agent creates a commit |

## Review outcomes

- Memory ownership, authority, trust, temporal semantics, and deletion semantics are explicit.
- Qdrant 1.19 is selected as the RC projection line; PostgreSQL 18 is canonical and Go 1.26 is the implementation baseline.
- Open questions requiring implementation evidence—not contract ambiguity—are tracked as Phase 1+ work: exact image digests, ranking weights, load-derived SLO revision, library choices, and physical privacy-erasure mechanism.

