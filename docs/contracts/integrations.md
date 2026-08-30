# Integration contracts

All integrations use TLS, workload identity, explicit tenant/audience, W3C trace context, bounded payloads, UTC timestamps, schema versions, and stable idempotency IDs. References are opaque; possessing one grants no access.

## ThinkPixelAG MemoryGrant

A signed grant contains issuer, audience `thinkpixelmem`, subject/principal, tenant, Run ID, unique grant ID, issued/not-before/expiry times, readable and writable MemorySpace IDs, allowed memory types, classification ceiling, operation set, per-request/item/token limits, and optional policy version. MEM validates signature, issuer, audience, tenant/Run binding, time window, revocation seam, and requested-operation intersection on every runtime call.

Expiry is fail closed. Revocation is checked through a bounded cache/introspection strategy selected by deployment policy. A grant never derives from memory content. Administrative API authorization uses OIDC scopes/roles and is distinct from a Run-scoped MemoryGrant.

Standalone `MemoryAuthorizer` exposes `Authorize(ctx, principal, operation, resource) -> Decision`; the default policy denies absent explicit tenant/space permission.

## ThinkPixelAR

AR submits versioned Run/Session event envelopes with immutable event IDs and `ar://` evidence references. MEM accepts only user-visible/normalized evidence, never provider-private chain-of-thought. AR retrieves ContextPacks using an AG grant. AR remains canonical for execution history.

## ThinkPixelWS

WS evidence uses Workspace ID, immutable generation, optional component/path, digest, classification, and `ws://` URI. MEM stores claims plus references, not repository/document bodies. Missing source content does not invalidate historical provenance but is reported during inspection.

## ThinkPixelTG

TG evidence identifies invocation and result, tool identity, result state (`succeeded`, `failed`, `partial`, `unknown`), observed time, digest, and verification class. Only the authenticated TG adapter can assign `verified-tool-output`. `unknown` remains unknown.

## ThinkPixelLLMGW

Extraction requests contain service/tenant identity, logical operation ID, bounded content, JSON output schema, model-policy class, timeout, and trace context. Responses return request reference, model/policy identifiers, structured output, usage reference, and finish/error status. Embedding requests additionally pin model, dimensions, and input digest. MEM stores request/model references but no provider secrets. Retries reuse the logical operation ID.

## ThinkPixelGR

Write inspection receives candidate ID, safe content, source/trust, classification, and risk context; retrieval inspection receives selected item metadata and bounded content. Decisions are `allow`, `deny`, `quarantine`, or `allow-with-warnings`, with reason codes and policy version. Protected operations fail closed on GR unavailability; policy may allow low-risk operations with an explicit warning and audit event.

## ThinkPixelMP

MEM exports versioned ProcedureCandidates, MemoryStrategies, ProfileSchemas, or extractor/reranker configurations for evaluation. MP returns a separately identified qualified artifact and evidence. Qualification never mutates the candidate into a Skill and never authorizes use; ThinkPixelAG remains the authorization point.

