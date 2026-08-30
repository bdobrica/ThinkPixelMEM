# Threat model

## Scope and assets

Protected assets are tenant isolation, canonical memory and revision integrity, provenance, deletion guarantees, MemoryGrants, classifications, audit history, and availability. Trust boundaries are diagrammed in `architecture.md`. Raw source content, agent/model output, imported memories, and retrieved memory are hostile by default.

## Threats and required controls

| Threat | Abuse case | Preventive controls | Detection/recovery |
| --- | --- | --- | --- |
| Persistent prompt injection | Source text stores instructions that later steer an agent | Treat memory as quoted data; never place it in policy channels; type ProcedureCandidates; GR hook | `instruction_like` risk flag, audit, quarantine |
| Memory poisoning | Attacker creates or reinforces false claims | authenticated source identity, independent trust/confidence, corroboration seam, candidate validation | dispute warnings, provenance inspection, correction |
| Source spoofing | Caller claims tool-verified evidence | integration-specific credentials and audience; infrastructure assigns source kind/trust | reject and audit metadata conflicts |
| Cross-tenant/space leakage | Query or index filter exposes another scope | tenant key on every row; canonical authorization; grant scope intersection; opaque IDs | security tests, access audit, anomaly metrics |
| Classification crossing | ContextPack exceeds caller ceiling | compare canonical classification to grant before and after retrieval | fail closed, audit denial |
| Sensitive inference | Model derives protected traits | schema allowlist, deny-by-default sensitive inference, purpose and consent hooks | candidate rejection/quarantine and metrics |
| Stale memory | Expired or superseded claim is presented as current | valid-time filters, status filters, freshness scoring | explicit warnings and repair scan |
| Malicious import | External memory imports authority or forged provenance | importer trust class, original identity, quarantine option, no authority conversion | import audit and deletion by source |
| Forged/replayed grant | Caller reuses another Run's authority | signed issuer/audience/tenant/run binding, expiry, nonce/JTI seam | denial audit and revocation cache |
| Index resurrection | Deleted memory returns after rebuild | canonical deletion filter, durable deletion generation, projection cleanup | rebuild tests and consistency scanner |
| Revision tampering | History is silently rewritten | immutable completed revisions, monotonic numbers, digests, DB constraints | audit verification |
| Secret persistence | Credentials enter memory or logs | bounded inputs, secret inspection/redaction policy, no provider credentials | quarantine, delete-by-source, redacted telemetry |
| Worker replay/race | Duplicate jobs create duplicate state | idempotency keys, leases with fencing tokens, unique constraints, outbox | retry/dead-letter audit and repair |
| Dependency compromise/outage | LLMGW, GR, Qdrant, or source service fails | timeouts, bounded requests, least-privilege identity; fail-closed protected operations | replay queues; rebuild Qdrant; alerting |

## Security invariants

- Retrieved memory cannot authorize tools, change policy, or become a trusted system instruction.
- Extractors cannot set authoritative metadata.
- Every disclosure is authorized against canonical state.
- Every accepted inference remains labeled as inference.
- Confidence and source trust are independent.
- A forgotten memory cannot reappear from any derived projection.
- GR augments but does not replace deterministic authorization and classification enforcement.

## Abuse-case tests

Phase gates include forged source kinds, malicious ProcedureCandidates, cross-tenant and cross-space search, classification crossing, expired grants, poisoned imported memory, duplicate events, worker crashes, Qdrant loss, and crash-during-forget.

