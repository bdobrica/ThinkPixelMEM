# ADR-0001: Service boundaries and authority invariants

- Status: accepted
- Date: 2026-08-30
- Deciders: ThinkPixelMEM maintainers
- Supersedes: none

## Context

Durable memory can accidentally duplicate source systems or become an authority channel. The boundary must remain valid in standalone and integrated deployments.

## Decision

ThinkPixelMEM owns durable learned context. ThinkPixelAR owns raw Session, Run, Attempt, and execution history. ThinkPixelWS owns Workspace files, documents, snapshots, and generations. MEM stores claims derived from those sources plus immutable evidence references; it does not copy their canonical content.

ThinkPixelAG owns Run-scoped read/write authority. Recall never creates or broadens a capability. Memory content never becomes governance policy. A `ProcedureCandidate` is untrusted data until separately evaluated, qualified by ThinkPixelMP, and authorized by ThinkPixelAG.

The invariant set is:

- `AR history != MEM learned context`;
- `WS source content != MEM learned claim`;
- `Recall(memory) != CapabilityGrant`;
- `MemoryContent != GovernancePolicy`;
- `ProcedureCandidate != ApprovedSkill`.

Standalone mode uses a local `MemoryAuthorizer` but makes no claim of full ThinkPixel governance equivalence.

## Consequences

Integration payloads carry stable references rather than bulk source content. Callers must obtain authority independently. Answer generation and tool execution remain outside this service.

## Alternatives considered

A combined assistant, history, and memory database was rejected because it couples retention, trust, and authority. Treating learned procedures as executable instructions was rejected because it creates durable prompt-injection authority.

## Verification

Architecture and security tests must reject cross-scope access, forged source types, and attempts to convert recalled content into grants or policy.

