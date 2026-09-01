# ThinkPixelMEM Release-Candidate TODO

This is the chronological implementation checklist for ThinkPixelMEM.

Execute the first unchecked item whose dependencies are complete.

An item is checked only after its acceptance evidence passes.

Follow the coding-agent and commit protocol in `PLAN.md`.

Status notation:

- `[ ]` pending
- `[x]` implemented and verified

Phase 0 artifacts were implemented and structurally validated in commit `7f45aee`; see `docs/phase-0-validation.md` and `docs/phase-0-traceability.md`.

Completion metadata format:

    — completed YYYY-MM-DD, commit <sha>, evidence: <commands/artifacts>

---

## Phase 0 — Decisions, threats, and contracts

- [x] ARC-001 Create `docs/`, `docs/adr/`, and `docs/contracts/` plus ADR template. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-002 Write system-context diagram for MEM, AG, AR, WS, TG, LLMGW, GR, MP, PostgreSQL, Qdrant, and clients. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-003 Write trust-boundary diagram distinguishing canonical memory, indexes, source systems, model extraction, runtime agents, and governance. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-004 Write threat model including persistent prompt injection, memory poisoning, source spoofing, cross-space leakage, sensitive inference, stale memory, and malicious imported memory. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-005 Record invariant: AR owns raw Session/execution truth; MEM owns learned long-term memory. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-006 Record invariant: WS owns work/source content; MEM stores learned claims plus references. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-007 Record invariant: memory recall does not grant capability. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-008 Record invariant: memory content does not become governance policy. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-009 Record invariant: ProcedureCandidate != approved Skill. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-010 Record invariant: observation != inference. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-011 Record invariant: confidence != source trust. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-012 Record invariant: authoritative metadata cannot be modified by extraction models. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-013 Record invariant: retrieval indexes are rebuildable projections, not canonical state. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-014 Define MemorySpace model. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-015 Define initial scope types: user, agent, workspace, team, organization, custom. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-016 Define MemorySpace ownership, classification, residency, retention, read/write policy. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-017 Define Episode schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-018 Define Claim schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-019 Define Relationship schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-020 Define Outcome schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-021 Define Lesson schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-022 Define ProcedureCandidate schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-023 Define ProfileSchema and Profile contracts. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-024 Define Entity/subject identity representation. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-025 Define source-kind vocabulary. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-026 Define source-trust vocabulary. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-027 Define confidence semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-028 Define classification vocabulary and crossing rules. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-029 Define temporal model: valid_from, valid_until, observed_at, recorded_at, superseded_at. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-030 Define Claim status transitions: active, disputed, superseded, withdrawn, quarantined, expired, deleted. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-031 Define immutable MemoryRevision format. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-032 Define correction semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-033 Define contradiction/dispute semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-034 Define temporal supersession semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-035 Define EvidenceReference schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-036 Define AR evidence/reference format. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-037 Define WS generation/component evidence/reference format. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-038 Define TG invocation/result evidence/reference format. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-039 Define user-message/explicit assertion provenance. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-040 Define imported-memory provenance. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-041 Define authoritative vs derived metadata classes. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-042 Define IngestionEvent schema and idempotency. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-043 Define MemoryCandidate schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-044 Define candidate validation pipeline. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-045 Define async extraction semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-046 Define MemoryExtractor interface. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-047 Define extraction strategy configuration. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-048 Define Consolidator interface/semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-049 Define deduplication rules. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-050 Define contradiction matching rules. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-051 Define ContextPack schema. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-052 Define retrieval query contract. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-053 Define retrieval score signals. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-054 Define retrieval warning model for dispute, staleness, low trust, poison risk. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-055 Define ContextPack token/item budgeting. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-056 Define RetrievalIndex interface. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-057 Select Qdrant version/features for dense+sparse reference implementation. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-058 Define embedding version/migration semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-059 Define lexical/sparse strategy. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-060 Define graph relationships while explicitly deferring graph DB requirement. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-061 Define MemoryGrant contract with ThinkPixelAG. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-062 Define read/write/type/classification limits in MemoryGrant. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-063 Define MemoryGrant expiry/revocation behavior. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-064 Define standalone MemoryAuthorizer. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-065 Define LLMGW extraction/embedding contract. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-066 Define GR write-inspection contract. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-067 Define GR retrieval-inspection contract. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-068 Define procedure-candidate promotion boundary with MP. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-069 Define retention policies and TTL. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-070 Define forget-by-memory, subject, space, source, and data-subject semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-071 Define legal-hold seam. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-072 Define deletion behavior for revisions/audit according to privacy vs compliance policy. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-073 Define derived-index deletion/rebuild guarantees. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-074 Define Profile field explainability requirements. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-075 Define sensitive-inference restrictions. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-076 Define poisoning/quarantine rules. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-077 Define source corroboration policy seam. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-078 Define PostgreSQL schema/invariants. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-079 Define worker job/lease/fencing semantics. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-080 Define transactional outbox. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-081 Define authentication/OIDC/tenant mapping. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-082 Define API authorization vs runtime MemoryGrant distinction. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-083 Draft OpenAPI. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-084 Define RFC 7807, UUIDv7, pagination, idempotency, W3C tracing, limits, SSE conventions. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-085 Define audit/event vocabulary. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-086 Define telemetry/redaction policy. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-087 Define target SLOs/capacity assumptions. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-088 Define supported-version policy for Go, PostgreSQL, Qdrant, LLMGW, GR, and relevant schemas. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-089 Compare SecondContext concepts against new domain and document which ideas migrate, change, or remain experimental. — completed 2026-08-30, commit 7f45aee, evidence: docs/phase-0-traceability.md
- [x] ARC-090 Validate Phase 0 schemas/docs/OpenAPI and commit evidence. — completed 2026-08-30, commit 7f45aee, evidence: ./scripts/validate-phase0.sh; npx --yes @redocly/cli lint api/openapi/openapi.yaml; git diff --check

---

## Phase 1 — Engineering foundation

- [x] ENG-001 Initialize Go module using supported pinned Go. — completed 2026-09-01, evidence: `go.mod`; `.go-version`; `go version`; `go env GOMOD GOVERSION GOTOOLCHAIN`; `go mod edit -json`
- [x] ENG-002 Create domain/app/ports/adapters repository layout. — completed 2026-09-01, evidence: `go list ./...`; layout comparison with `PLAN.md` §42; `git diff --check`
- [x] ENG-003 Add dependency/source/license policy. — completed 2026-09-01, evidence: `docs/dependency-policy.md`; `go mod verify`; `go list -m all`; policy link validation; scoped `git diff --check`
- [x] ENG-004 Implement typed configuration with validation and secret references. — completed 2026-09-01, evidence: `internal/config`; `docs/operations/configuration.md`; `go test ./internal/config`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-005 Implement structured logging with memory/run/workspace correlation and secret-content redaction. — completed 2026-09-01, evidence: `internal/telemetry/logging`; `docs/operations/logging.md`; `go test ./internal/telemetry/logging`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-006 Implement Prometheus and OpenTelemetry initialization. — completed 2026-09-01, evidence: `internal/telemetry`; `docs/operations/telemetry.md`; `go test ./internal/telemetry`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-007 Add UUIDv7, typed memory IDs, injectable clock, typed errors, bounded strings, digests, authenticated cursors. — completed 2026-09-01, evidence: `internal/domain`; `internal/ports/clock`; `internal/security`; `docs/operations/foundation-primitives.md`; `go test ./internal/domain ./internal/security ./internal/ports/clock`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-008 Implement HTTP baseline with request IDs, tracing, RFC 7807, limits, graceful shutdown, `/livez`, `/readyz`, `/metrics`. — completed 2026-09-01, evidence: `internal/adapters/http`; `docs/operations/http.md`; `go test ./internal/adapters/http`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-009 Add OpenAPI generation/validation/drift checks. — completed 2026-09-01, evidence: `api/openapi/oapi-codegen.yaml`; `internal/adapters/http/openapi/types.gen.go`; `internal/tools/openapicheck`; `scripts/openapi.sh`; `docs/operations/openapi.md`; `make openapi-validate`; `make openapi-check`; `go test ./...`; `go vet ./...`; scoped `git diff --check`
- [x] ENG-010 Create root Makefile. — completed 2026-09-01, evidence: `Makefile`; `docs/operations/development.md`; `make help`; `make phase0-validate`; `make openapi-check`; scoped `git diff --check`
- [x] ENG-011 Add format/vet/lint/unit/race/vulnerability/license/build checks. — completed 2026-09-01, evidence: `Makefile`; pinned Go tools in `go.mod`; `make format-check vet-check lint-check unit-check race-check vulnerability-check license-check build-check`; `go mod verify`; scoped `git diff --check`
- [x] ENG-012 Add PostgreSQL development dependency/migration command. — completed 2026-09-01, evidence: `compose.yaml`; pinned Tern tool in `go.mod`; `docs/operations/postgresql.md`; `docker compose config --quiet`; `make migrate-status`; scoped `git diff --check`
- [x] ENG-013 Add disposable Qdrant development/test dependency. — completed 2026-09-01, evidence: `compose.yaml`; `docs/operations/qdrant.md`; `docker compose config --quiet`; healthy `make qdrant-up`; scoped `git diff --check`
- [x] ENG-014 Create `thinkpixelmemctl` CLI skeleton. — completed 2026-09-01, evidence: `cmd/thinkpixelmemctl`; `internal/cli`; `docs/operations/cli.md`; `go test ./internal/cli`; `make cli-build`; scoped `git diff --check`
- [x] ENG-015 Create hardened non-root image. — completed 2026-09-01, evidence: `Dockerfile`; `.dockerignore`; `cmd/thinkpixelmem`; `docs/operations/container-image.md`; `make image-check`; `go test ./...`; scoped `git diff --check`
- [x] ENG-016 Add CI. — completed 2026-09-01, evidence: `.github/workflows/ci.yml`; `docs/operations/continuous-integration.md`; workflow syntax validation; `make verify`; `make image-check`; scoped `git diff --check`
- [x] ENG-017 Add hygiene checks preventing test memories, secrets, model keys, Qdrant dumps, and local DB artifacts from Git. — completed 2026-09-01, evidence: `.gitignore`; `internal/tools/repositoryhygiene`; `docs/operations/repository-hygiene.md`; `go test ./internal/tools/repositoryhygiene`; `make hygiene-check`; focused whitespace inspection
- [x] ENG-018 Start `docs/supported-versions.md`. — completed 2026-09-01, evidence: `docs/supported-versions.md`; pin consistency inspection across `.go-version`, `go.mod`, `Dockerfile`, and `compose.yaml`; `make phase0-validate`; scoped `git diff --check`
- [x] ENG-019 Run clean-checkout baseline gate. — completed 2026-09-01, evidence: clean committed 128-file repository snapshot; `make verify`; `make image-check`; clean post-gate `git status --porcelain`
- [ ] ENG-020 Publish Phase 1 evidence and commit.

---

## Phase 2 — Canonical MemorySpace and memory persistence

- [ ] DB-001 Add migration framework and tenant schema.
- [ ] DB-002 Add MemorySpace table/domain/repository.
- [ ] DB-003 Add MemorySpace lifecycle/retention/classification fields.
- [ ] DB-004 Add Episode table/domain/repository.
- [ ] DB-005 Add Claim logical-identity table.
- [ ] DB-006 Add MemoryRevision table.
- [ ] DB-007 Enforce immutable completed revisions.
- [ ] DB-008 Add active revision pointer.
- [ ] DB-009 Add temporal validity columns/indexes.
- [ ] DB-010 Add source kind/source trust/confidence.
- [ ] DB-011 Add EvidenceReference table.
- [ ] DB-012 Add Relationship table/domain.
- [ ] DB-013 Add Outcome table/domain.
- [ ] DB-014 Add Lesson table/domain.
- [ ] DB-015 Add ProcedureCandidate table/domain.
- [ ] DB-016 Add ProfileSchema table/domain.
- [ ] DB-017 Add Profile/ProfileField projection tables.
- [ ] DB-018 Add authoritative metadata fields.
- [ ] DB-019 Add classification/residency metadata.
- [ ] DB-020 Add quarantine/dispute/supersession metadata.
- [ ] DB-021 Add AuditEvent.
- [ ] DB-022 Add IdempotencyRecord.
- [ ] DB-023 Add transactional OutboxMessage.
- [ ] DB-024 Add transaction manager.
- [ ] IAM-001 Implement OIDC/JWT validation.
- [ ] IAM-002 Implement claim-to-tenant/principal mapping.
- [ ] IAM-003 Implement administrative MemoryAuthorizer.
- [ ] IAM-004 Add safe explicit development auth mode.
- [ ] API-001 Implement MemorySpace create/list/get.
- [ ] API-002 Implement explicit memory create.
- [ ] API-003 Implement memory get/list.
- [ ] API-004 Implement Episode create/get.
- [ ] API-005 Implement revision history read API.
- [ ] DB-025 Add real PostgreSQL migration tests.
- [ ] DB-026 Add tenant-isolation tests.
- [ ] DB-027 Add revision immutability tests.
- [ ] DB-028 Add active-revision concurrency tests.
- [ ] DB-029 Add temporal interval property tests.
- [ ] DB-030 Add outbox/idempotency race tests.
- [ ] DB-031 Commit Phase 2 evidence.

---

## Phase 3 — Ingestion, extraction, and consolidation

- [ ] ING-001 Add IngestionEvent table/domain.
- [ ] ING-002 Add stable external event ID/idempotency.
- [ ] ING-003 Add extraction Job schema.
- [ ] ING-004 Add worker lease/fence.
- [ ] ING-005 Add job retry/dead-letter metadata.
- [ ] CAND-001 Add MemoryCandidate persistence/domain.
- [ ] CAND-002 Ensure authoritative metadata comes from IngestionEvent, never extractor output.
- [ ] CAND-003 Add candidate acceptance/rejection/quarantine states.
- [ ] EXT-001 Implement MemoryExtractor port.
- [ ] LLM-001 Implement ThinkPixelLLMGW client adapter.
- [ ] LLM-002 Implement extraction request with bounded source content.
- [ ] LLM-003 Record extractor strategy/version and LLMGW request reference.
- [ ] LLM-004 Validate structured extractor output.
- [ ] LLM-005 Reject model attempts to alter tenant/MemorySpace/source trust/classification.
- [ ] EXT-002 Extract Episode candidate.
- [ ] EXT-003 Extract Claim candidates.
- [ ] EXT-004 Extract Relationship candidates.
- [ ] EXT-005 Extract Outcome candidates where supported.
- [ ] EXT-006 Extract preference/profile candidates.
- [ ] EXT-007 Extract ProcedureCandidates only when strategy enables them.
- [ ] VAL-001 Implement candidate schema validation.
- [ ] VAL-002 Implement confidence threshold.
- [ ] VAL-003 Implement source trust attachment from infrastructure metadata.
- [ ] VAL-004 Implement sensitive-inference policy hook.
- [ ] VAL-005 Implement duplicate candidate detection.
- [ ] VAL-006 Implement instruction/poison-risk flagging hook.
- [ ] CON-001 Implement Consolidator port.
- [ ] CON-002 Implement deterministic Claim matching baseline.
- [ ] CON-003 Implement revision creation for compatible new evidence.
- [ ] CON-004 Implement dispute creation for conflicting evidence.
- [ ] CON-005 Implement temporal supersession.
- [ ] CON-006 Preserve evidence on consolidated revisions.
- [ ] ING-006 Implement internal ingestion API.
- [ ] ING-007 Implement extraction worker.
- [ ] ING-008 Implement consolidation worker.
- [ ] ING-009 Add duplicate-event replay tests.
- [ ] ING-010 Kill/restart extraction worker tests.
- [ ] ING-011 Kill/restart consolidation worker tests.
- [ ] ING-012 Verify duplicate LLM response cannot duplicate canonical memory.
- [ ] ING-013 Verify LLM cannot spoof verified-tool source type.
- [ ] ING-014 Commit Phase 3 evidence.

---

## Phase 4 — Retrieval and ContextPack

- [ ] RET-001 Implement RetrievalIndex port.
- [ ] QDR-001 Implement Qdrant adapter.
- [ ] QDR-002 Create tenant/MemorySpace metadata schema.
- [ ] QDR-003 Implement dense-vector upsert.
- [ ] QDR-004 Implement sparse/lexical representation.
- [ ] QDR-005 Implement metadata filtering.
- [ ] EMB-001 Implement embedding adapter through LLMGW.
- [ ] EMB-002 Store embedding model/version metadata.
- [ ] EMB-003 Implement re-embedding support.
- [ ] IDX-001 Add index queue/job.
- [ ] IDX-002 Add index worker.
- [ ] IDX-003 Add index deletion.
- [ ] IDX-004 Add index rebuild from PostgreSQL.
- [ ] IDX-005 Add index consistency/repair scanner.
- [ ] RET-002 Implement retrieval query parser.
- [ ] RET-003 Enforce MemorySpace before query.
- [ ] RET-004 Enforce memory-type filters.
- [ ] RET-005 Enforce classification ceiling.
- [ ] RET-006 Enforce temporal validity.
- [ ] RET-007 Apply source-trust signal.
- [ ] RET-008 Apply confidence signal.
- [ ] RET-009 Apply recency/freshness signal.
- [ ] RET-010 Apply dispute/supersession penalty.
- [ ] RET-011 Apply redundancy dedup.
- [ ] RET-012 Add goal relevance.
- [ ] RET-013 Add poison-risk penalty/flag.
- [ ] CP-001 Implement ContextPack domain/schema.
- [ ] CP-002 Implement token/item budget.
- [ ] CP-003 Include provenance/source trust/confidence/status with items.
- [ ] CP-004 Include warnings for disputed/stale/low-trust memory.
- [ ] CP-005 Persist bounded ContextPack metadata/evidence.
- [ ] CP-006 Implement `POST /v1/context/retrieve`.
- [ ] CP-007 Implement ContextPack get.
- [ ] RET-014 Add retrieval-debug explain output for authorized callers.
- [ ] RET-015 Add dense+sparse retrieval tests.
- [ ] RET-016 Add temporal query tests.
- [ ] RET-017 Add cross-space leakage tests.
- [ ] RET-018 Add Qdrant loss/rebuild test.
- [ ] RET-019 Verify Qdrant deletion/recreation loses no canonical memory.
- [ ] RET-020 Commit Phase 4 evidence.

---

## Phase 5 — Correction, contradiction, retention, and forget

- [ ] COR-001 Implement memory correction API.
- [ ] COR-002 Create new immutable revision on correction.
- [ ] COR-003 Record correction reason/source.
- [ ] COR-004 Implement Claim dispute.
- [ ] COR-005 Implement supersession.
- [ ] COR-006 Implement withdrawal.
- [ ] COR-007 Implement quarantine.
- [ ] COR-008 Verify correction does not rewrite prior revision.
- [ ] LIFE-001 Implement TTL/expiry worker.
- [ ] LIFE-002 Implement MemorySpace retention policy evaluation.
- [ ] LIFE-003 Add expiry events.
- [ ] FGT-001 Implement forget single memory.
- [ ] FGT-002 Implement forget by subject.
- [ ] FGT-003 Implement forget MemorySpace.
- [ ] FGT-004 Implement forget by source.
- [ ] FGT-005 Implement data-subject deletion hook.
- [ ] FGT-006 Implement legal-hold denial seam.
- [ ] FGT-007 Delete/rebuild Qdrant projection.
- [ ] FGT-008 Rebuild affected Profiles.
- [ ] FGT-009 Remove affected relationship projection.
- [ ] FGT-010 Preserve permitted non-sensitive deletion audit marker.
- [ ] FGT-011 Add crash-during-forget recovery.
- [ ] FGT-012 Add repeated idempotent forget.
- [ ] FGT-013 Verify forgotten memory cannot reappear after index rebuild.
- [ ] LIFE-004 Add retention/forget metrics.
- [ ] LIFE-005 Commit Phase 5 evidence.

---

## Phase 6 — Profiles, outcomes, lessons, and procedure candidates

- [ ] PROF-001 Implement ProfileSchema registration/read API.
- [ ] PROF-002 Version ProfileSchemas immutably.
- [ ] PROF-003 Add allowed source/memory type fields.
- [ ] PROF-004 Add sensitive-field policy metadata.
- [ ] PROF-005 Implement Profile builder worker.
- [ ] PROF-006 Record underlying memory references for every field.
- [ ] PROF-007 Record confidence/source trust/freshness per field.
- [ ] PROF-008 Implement profile rebuild after memory correction/forget.
- [ ] PROF-009 Add profile inspect/explain API.
- [ ] OUT-001 Implement Outcome creation/linking.
- [ ] OUT-002 Link TG outcome evidence.
- [ ] OUT-003 Link user feedback outcome.
- [ ] LES-001 Implement Lesson derivation strategy.
- [ ] LES-002 Store supporting Episode/Outcome references.
- [ ] LES-003 Make Lesson revisable/disputable.
- [ ] PROC-001 Implement ProcedureCandidate derivation.
- [ ] PROC-002 Mark procedure candidates as untrusted memory context.
- [ ] PROC-003 Prevent procedure candidate from entering trusted ContextPack instruction section.
- [ ] PROC-004 Add export/read API for future MP promotion.
- [ ] PROF-010 Add tests proving profile value traceability.
- [ ] PROF-011 Add sensitive-inference denied-field tests.
- [ ] PROC-005 Add malicious persistent instruction tests.
- [ ] PROC-006 Commit Phase 6 evidence.

---

## Phase 7 — ThinkPixel integrated MVP

- [ ] TAG-001 Implement AG MemoryGrant verifier adapter.
- [ ] TAG-002 Implement secure MEM↔AG authentication.
- [ ] TAG-003 Verify readable MemorySpaces.
- [ ] TAG-004 Verify writable MemorySpaces.
- [ ] TAG-005 Verify memory-type restrictions.
- [ ] TAG-006 Verify classification ceiling.
- [ ] TAG-007 Enforce retrieval item/token limits.
- [ ] TAG-008 Enforce grant expiry.
- [ ] TAG-009 Enforce grant revocation.
- [ ] TAG-010 Ensure remembered claims cannot expand AG grant.
- [ ] TAR-001 Implement AR normalized event ingestion adapter.
- [ ] TAR-002 Accept only user-visible/operational AR evidence, not hidden chain-of-thought.
- [ ] TAR-003 Preserve Run/Session IDs.
- [ ] TWS-001 Implement WS evidence-reference adapter.
- [ ] TWS-002 Bind Claims to exact Workspace generation/component where available.
- [ ] TWS-003 Avoid copying entire Workspace content into MEM.
- [ ] TTG-001 Implement TG outcome/evidence adapter.
- [ ] TTG-002 Preserve UNKNOWN/ambiguous tool outcomes.
- [ ] TTG-003 Mark verified TG results with trusted source class.
- [ ] LLM-006 Complete LLMGW extraction integration.
- [ ] LLM-007 Complete LLMGW embedding integration.
- [ ] GR-001 Implement GR candidate-write inspection adapter.
- [ ] GR-002 Implement GR ContextPack inspection hook.
- [ ] GR-003 Define fail-closed behavior for protected memory operations.
- [ ] E2E-001 Run agent on Workspace and ingest review Episode.
- [ ] E2E-002 Extract initial Claim with model-inference source.
- [ ] E2E-003 Ingest contradictory verified TG/WS evidence.
- [ ] E2E-004 Supersede/dispute prior Claim without deletion.
- [ ] E2E-005 Start new AR Session/new agent version.
- [ ] E2E-006 Retrieve authorized ContextPack containing current Claim + historical warning.
- [ ] E2E-007 Verify new agent receives no extra TG/model/runtime capability from memory.
- [ ] E2E-008 Correct memory explicitly and verify revision history.
- [ ] E2E-009 Forget selected memory and verify derived-index removal.
- [ ] MVP-001 Run complete ThinkPixelMEM reference scenario.
- [ ] MVP-002 Publish `docs/mvp-thinkpixel-evidence.md`.
- [ ] MVP-003 Commit Phase 7 integrated milestone.

---

## Phase 8 — Security, poisoning, resilience, and retrieval hardening

- [ ] SEC-001 Create persistent prompt-injection memory fixture.
- [ ] SEC-002 Verify instruction-like memory remains untrusted data.
- [ ] SEC-003 Attempt memory claiming additional AG capability.
- [ ] SEC-004 Verify capability remains unchanged.
- [ ] SEC-005 Attempt memory claiming platform policy override.
- [ ] SEC-006 Verify no governance change.
- [ ] SEC-007 Attempt malicious ProcedureCandidate activation.
- [ ] SEC-008 Verify not promoted to trusted Skill/instruction.
- [ ] SEC-009 Attempt forged `verified-tool-output`.
- [ ] SEC-010 Verify authenticated-source enforcement.
- [ ] SEC-011 Add cross-tenant retrieval attacks.
- [ ] SEC-012 Add cross-MemorySpace retrieval attacks.
- [ ] SEC-013 Add classification-ceiling attacks.
- [ ] SEC-014 Add imported poisoned-memory tests.
- [ ] SEC-015 Add hostile Workspace/repository source poisoning tests.
- [ ] SEC-016 Add hostile web/tool-output poisoning tests.
- [ ] SEC-017 Add source corroboration tests.
- [ ] SEC-018 Add quarantine workflow.
- [ ] SEC-019 Add secret-like memory write tests.
- [ ] SEC-020 Add sensitive-personal-attribute extraction tests.
- [ ] SEC-021 Verify authoritative tenant/source/classification metadata cannot be model-overwritten.
- [ ] CHAOS-001 Restart API during write.
- [ ] CHAOS-002 Kill ingestion worker.
- [ ] CHAOS-003 Kill extraction worker after LLM response/before commit.
- [ ] CHAOS-004 Kill consolidation worker.
- [ ] CHAOS-005 Lose Qdrant.
- [ ] CHAOS-006 Rebuild Qdrant from PostgreSQL.
- [ ] CHAOS-007 Interrupt PostgreSQL.
- [ ] CHAOS-008 Interrupt LLMGW.
- [ ] CHAOS-009 Interrupt GR.
- [ ] CHAOS-010 Crash during profile rebuild.
- [ ] CHAOS-011 Crash during forget.
- [ ] CHAOS-012 Replay duplicated outbox/ingestion events.
- [ ] CAP-001 Add per-tenant MemorySpace/memory/index quotas.
- [ ] CAP-002 Add worker queue/backpressure.
- [ ] CAP-003 Load test ingestion.
- [ ] CAP-004 Load test extraction/consolidation.
- [ ] CAP-005 Load test ContextPack retrieval.
- [ ] CAP-006 Load test Qdrant index/rebuild.
- [ ] CAP-007 Load test forget/profile rebuild.
- [ ] CAP-008 Document capacity envelope.
- [ ] HARD-001 Publish security/poisoning evidence.
- [ ] HARD-002 Commit Phase 8.

---

## Phase 9 — Production packaging and operations

- [ ] OPS-001 Finalize reproducible hardened `thinkpixelmem` image.
- [ ] OPS-002 Finalize `thinkpixelmemctl`.
- [ ] OPS-003 Create Helm chart for API/workers/migrations/config/secrets/Service.
- [ ] OPS-004 Add least-privilege Kubernetes RBAC.
- [ ] OPS-005 Add NetworkPolicies for PostgreSQL, Qdrant, AG, LLMGW, GR, and configured integrations.
- [ ] OPS-006 Add hardened pod security context.
- [ ] OPS-007 Add startup/readiness/liveness probes.
- [ ] OPS-008 Add PDB/topology guidance.
- [ ] OPS-009 Add optional HPA.
- [ ] OPS-010 Add ServiceMonitor/PodMonitor.
- [ ] OPS-011 Build dashboards for ingestion, memory state, extraction, retrieval, Qdrant, profile rebuild, forget, workers, PostgreSQL.
- [ ] OPS-012 Define SLO alerts/runbooks.
- [ ] OPS-013 Write installation/configuration runbook.
- [ ] OPS-014 Write MemorySpace/retention policy runbook.
- [ ] OPS-015 Write AG MemoryGrant integration runbook.
- [ ] OPS-016 Write LLMGW model/embedding configuration runbook.
- [ ] OPS-017 Write Qdrant rebuild/recovery runbook.
- [ ] OPS-018 Write memory-poisoning/quarantine incident runbook.
- [ ] OPS-019 Write correction/dispute runbook.
- [ ] OPS-020 Write forget/data-subject deletion runbook.
- [ ] OPS-021 Write backup/restore runbook.
- [ ] OPS-022 Test PostgreSQL backup/restore.
- [ ] OPS-023 Prove Qdrant can be completely rebuilt after restore.
- [ ] OPS-024 Test fresh install.
- [ ] OPS-025 Test schema/chart upgrade.
- [ ] OPS-026 Test failed upgrade/roll-forward/rollback path.
- [ ] OPS-027 Test rolling restart during ingestion/retrieval.
- [ ] OPS-028 Run production-like load test.
- [ ] OPS-029 Generate SBOM/vulnerability reports.
- [ ] OPS-030 Add build provenance/signing/checksums.
- [ ] OPS-031 Add release automation.
- [ ] OPS-032 Commit Phase 9 evidence.

---

## Phase 10 — Release-candidate closure

- [ ] RC-001 Freeze OpenAPI and error contracts.
- [ ] RC-002 Freeze MemorySpace, Claim, Episode, Revision, EvidenceReference, ContextPack, ProfileSchema, and MemoryGrant integration schemas.
- [ ] RC-003 Freeze memory/source/status vocabulary for RC.
- [ ] RC-004 Run generated-artifact/backward-compatibility checks.
- [ ] RC-005 Run `make verify` from clean checkout.
- [ ] RC-006 Archive unit/race/fuzz/PostgreSQL/Qdrant/security/integration/e2e evidence.
- [ ] RC-007 Confirm AR raw Session truth is not duplicated as MEM canonical memory.
- [ ] RC-008 Confirm WS files/documents remain references rather than duplicated corpus by default.
- [ ] RC-009 Confirm observation and inference remain distinguishable.
- [ ] RC-010 Confirm confidence and source trust remain independent.
- [ ] RC-011 Confirm authoritative metadata cannot be model-modified.
- [ ] RC-012 Confirm revisions preserve corrected history.
- [ ] RC-013 Confirm temporal supersession answers current and historical queries correctly.
- [ ] RC-014 Confirm disputed Claims are surfaced as disputed.
- [ ] RC-015 Confirm ContextPack is limited to AG-authorized MemorySpaces/types/classification.
- [ ] RC-016 Confirm memory cannot expand AG capabilities.
- [ ] RC-017 Confirm memory cannot override governance policy.
- [ ] RC-018 Confirm ProcedureCandidate cannot become trusted Skill automatically.
- [ ] RC-019 Confirm persistent prompt injection remains untrusted context.
- [ ] RC-020 Confirm forged trusted-source metadata is rejected.
- [ ] RC-021 Confirm forget removes canonical/derived retrievable state according to policy.
- [ ] RC-022 Confirm Qdrant can be destroyed/rebuilt without losing canonical memory.
- [ ] RC-023 Confirm profile fields are explainable from source memories.
- [ ] RC-024 Confirm LLMGW/provider credentials are absent from MEM persistent state.
- [ ] RC-025 Confirm no unresolved critical/high security findings.
- [ ] RC-026 Confirm no undocumented fail-open memory authorization path.
- [ ] RC-027 Confirm no required flaky/skipped tests without disposition.
- [ ] RC-028 Confirm SLO/capacity envelope.
- [ ] RC-029 Exercise install/upgrade/rollback/backup/restore/Qdrant-loss/PostgreSQL-loss/LLMGW-loss/GR-loss game days.
- [ ] RC-030 Reconcile every TODO with implementation/tests/docs/commits.
- [ ] RC-031 Update README with architecture, memory model, temporal Claims, ContextPack, security, ThinkPixel integration, deployment, and limitations.
- [ ] RC-032 Create numbered ADRs for durable decisions.
- [ ] RC-033 Ensure ADRs preserve rejected alternatives and lessons.
- [ ] RC-034 Prepare RC release notes, support matrix, operator checklist, and artifact inventory.
- [ ] RC-035 Document post-RC scope.
- [ ] RC-036 Remove `PLAN.md` and `TODO.md` only after durable rationale is transferred.
- [ ] RC-037 Run docs/link validation and `make verify`.
- [ ] RC-038 Commit final documentation transition.
- [ ] RC-039 Build release artifacts from exact commit and verify digest/checksum/SBOM/provenance.
- [ ] RC-040 Tag RC only after all gates pass.

---

## Deferred / post-RC backlog

- [ ] FUT-001 Add graph retrieval backend.
- [ ] FUT-002 Add Graphiti-compatible temporal graph projection if useful.
- [ ] FUT-003 Add richer MemoryStrategy artifact support through ThinkPixelMP.
- [ ] FUT-004 Add ProfileSchema distribution through ThinkPixelMP.
- [ ] FUT-005 Add evaluated ProcedureCandidate → Skill promotion workflow.
- [ ] FUT-006 Add memory federation/import/export standard.
- [ ] FUT-007 Add external managed-memory importers.
- [ ] FUT-008 Add cross-enterprise shared-memory federation.
- [ ] FUT-009 Add multimodal memory evidence.
- [ ] FUT-010 Add richer entity resolution.
- [ ] FUT-011 Add contradiction reasoning across independent MemorySpaces.
- [ ] FUT-012 Add organization-wide memory analytics.
- [ ] FUT-013 Add richer memory quality/evaluation framework integration.
- [ ] FUT-014 Add automatic corroboration workflows.
- [ ] FUT-015 Add advanced temporal reasoning.
- [ ] FUT-016 Add hybrid retrieval integration with separate ThinkPixel knowledge/RAG service.
- [ ] FUT-017 Add privacy-preserving or confidential-memory execution modes.
- [ ] FUT-018 Add offline/air-gapped memory synchronization.
- [ ] FUT-019 Add policy-controlled memory sharing between teams/agents.
- [ ] FUT-020 Evaluate CRDT-style shared memory only if real multi-writer requirements emerge.

---

## Progress log

Append one row per completed atomic item or tightly coupled group.

Do not delete historical entries.

Supersede obsolete assumptions with a later entry.

Date | TODO IDs | Commit | Verification evidence | Notes/deviations
--- | --- | --- | --- | ---
YYYY-MM-DD | `ARC-...` | `<sha>` | `<commands/artifacts>` | `<notes>`
2026-09-01 | `ENG-001` | pending | `go version`; `go env GOMOD GOVERSION GOTOOLCHAIN`; `go mod edit -json` | Module `github.com/bdobrica/ThinkPixelMEM`; Go language baseline 1.26.0; toolchain pinned to 1.26.7 in `go.mod` and `.go-version`. Aggregate gate blocked by unrelated pre-existing CRLF working-tree changes.
2026-09-01 | `ENG-002` | pending | `go list ./...`; layout comparison with `PLAN.md` §42; `git diff --check` | Added tracked package boundaries and empty placeholders only; executable, migration, deployment, and test implementations remain assigned to later TODO items.
2026-09-01 | `ENG-003` | pending | `docs/dependency-policy.md`; `go mod verify`; `go list -m all`; policy link validation; scoped `git diff --check` | Defined dependency admission, source provenance, license review, vulnerability handling, and exception requirements; automated enforcement remains assigned to ENG-011 and release SBOM/provenance to Phase 9.
