# Canonical domain contract

This document is normative for Phase 0. Fields named `*_id` use UUIDv7. Timestamps are RFC 3339 UTC. Unknown JSON properties are rejected on trusted/internal contracts and ignored only where a public schema explicitly permits forward compatibility.

## Common value objects

- `TenantID`, `PrincipalID`, `MemorySpaceID`, `MemoryID`, `RevisionID`, `EvidenceID`, `RunID`, `SessionID`, `WorkspaceID`: non-empty typed identifiers.
- `Classification`: ordered `public < internal < confidential < restricted`; deployments may add labels only with an explicit partial-order/crossing policy.
- `Residency`: policy-defined region set plus optional storage restrictions.
- `Confidence`: decimal `[0,1]`, model/rule support for content, never source authority.
- `SourceTrust`: `untrusted`, `low`, `medium`, `high`, `verified`; assigned by authenticated infrastructure policy.
- `SourceKind`: `explicit-user-assertion`, `verified-tool-output`, `workspace-observation`, `agent-observation`, `model-inference`, `imported-external-memory`, `system-derived`.
- `MemoryStatus`: `active`, `disputed`, `superseded`, `withdrawn`, `quarantined`, `expired`, `deleted`.
- `MemoryType`: `episode`, `claim`, `relationship`, `outcome`, `lesson`, `procedure-candidate`.

## MemorySpace

The primary isolation, ownership, classification, residency, retention, and strategy boundary.

Required fields: `id`, `tenant_id`, `scope_type`, `scope_subject`, `name`, `classification`, `residency`, `retention_policy_id`, `strategy_id`, `read_policy`, `write_policy`, `state`, `created_at`, `updated_at`.

`scope_type` is one of `user`, `agent`, `workspace`, `team`, `organization`, `custom`. The tuple `(tenant_id, scope_type, scope_subject)` is unique among live spaces. Ownership does not imply visibility. Classification may be raised; lowering requires an authorized reclassification workflow. Residency changes require a migration workflow. States are `active`, `suspended`, `deleting`, `deleted`.

## Entity identity

An `EntityRef` is `{entity_id?, type, namespace, external_id?, display_name?}`. Within a tenant, `(namespace, external_id)` is unique when `external_id` is present. Display names are non-authoritative. Candidate entities without stable identity receive a server UUID and may later be merged through an auditable alias operation.

## Memory objects

All canonical objects include `id`, `tenant_id`, `memory_space_id`, `classification`, `status`, `created_at`, `created_by_principal`, evidence links, and lifecycle metadata.

- `Episode`: immutable occurrence with `title`, `summary`, `occurred_from`, `occurred_until`, entity refs, source refs, importance, and optional outcome refs. Later interpretation uses annotations or revisions; the occurrence is not rewritten.
- `Claim`: stable logical identity for `(subject, predicate, value_type)` and an `active_revision_id`.
- `Relationship`: typed edge with source entity, predicate, target entity, temporal interval, and Claim-equivalent provenance/revision semantics.
- `Outcome`: result linked to an Episode/action with `state` (`succeeded`, `failed`, `partial`, `unknown`), observed effects, and evidence. `unknown` cannot be normalized to success/failure.
- `Lesson`: derived generalization with supporting Episode/Outcome IDs, derivation strategy/version, applicability interval, confidence, and revisions.
- `ProcedureCandidate`: untrusted learned procedure with steps, applicability, evidence, risk flags, and review state. It cannot be executed or treated as a Skill.

## Claim and MemoryRevision

A Claim revision contains `id`, `claim_id`, positive monotonic `revision`, normalized `subject`, `predicate`, typed `value`, `valid_from?`, `valid_until?`, `observed_at`, server `recorded_at`, `status`, `confidence`, `source_kind`, `source_trust`, `classification`, `residency`, evidence IDs, actor/source, reason, `previous_revision_id?`, extraction metadata, and integrity digest.

Intervals are half-open. A correction appends a revision under optimistic concurrency (`expected_revision`), never updates a completed revision, and atomically changes the Claim's active pointer. Supersession closes the previous validity interval and records links. Contradictory evidence creates a dispute when temporal succession cannot resolve it. Withdrawal means the asserting authority retracts the claim; quarantine prevents ordinary retrieval; expiry follows policy; deletion follows the forget contract.

## EvidenceReference

Fields: `id`, `tenant_id`, `kind`, `source_service`, `source_id`, `immutable_version?`, `uri?`, `digest?`, `observed_at`, `source_trust`, `classification`, `safe_excerpt?`, `created_at`. Kinds are `ar-event`, `ar-run`, `ar-session`, `ws-generation`, `ws-component`, `tg-invocation`, `tg-result`, `user-message`, `external-document`, `imported-memory`, `model-extraction-request`.

Canonical URI shapes:

- `ar://{tenant}/runs/{run_id}/events/{event_id}`;
- `ws://{tenant}/workspaces/{workspace_id}/generations/{generation}[/{component}]`;
- `tg://{tenant}/invocations/{invocation_id}/results/{result_id}`;
- `user-message://{tenant}/sessions/{session_id}/messages/{message_id}`;
- `import://{tenant}/{source_system}/{original_id}`.

Evidence identity and trust are immutable. Large source content remains in its owning system. Excerpts are bounded, classified, and digest-linked.

## ProfileSchema and Profile

`ProfileSchema` is immutable and versioned by `(name, version)`. It defines subject type, typed fields, derivation rules, allowed memory/source kinds, sensitivity policy, freshness, retention, and minimum trust/confidence.

`Profile` is a rebuildable projection keyed by schema version and subject. Every `ProfileField` includes value, confidence, source trust, freshness timestamp, and contributing memory/revision IDs. A field without traceable support is invalid. Sensitive fields are denied unless the schema and tenant policy explicitly allow their purpose and source categories.

## Authoritative versus derived metadata

Authoritative: tenant, authenticated principal, MemorySpace, Run/Session/Workspace IDs, evidence identities, source kind/trust, classification, residency, server timestamps, grant constraints, and retention/legal-hold decisions.

Derived: summaries, normalized entities, topics, candidate relationships, confidence, importance, extraction risk flags, reranking score. Derived values never overwrite authoritative values.

