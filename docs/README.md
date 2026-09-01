# ThinkPixelMEM documentation

The accepted ADRs and versioned API/contracts are authoritative. `PLAN.md` and `TODO.md` describe implementation intent and sequencing; the root README is orientation only.

## Architecture and security

- [Platform alignment and ownership boundary](../ALIGNMENT.md)
- [Architecture and trust boundaries](architecture.md)
- [Threat model](threat-model.md)
- [Supported versions](supported-versions.md)
- [Dependency, source, and license policy](dependency-policy.md)
- [Developer commands](operations/development.md)
- [Continuous integration](operations/continuous-integration.md)
- [Repository hygiene](operations/repository-hygiene.md)
- [Runtime configuration](operations/configuration.md)
- [Structured logging](operations/logging.md)
- [Metrics and tracing](operations/telemetry.md)
- [HTTP service baseline](operations/http.md)
- [Command-line client](operations/cli.md)
- [Service container image](operations/container-image.md)
- [OpenAPI development workflow](operations/openapi.md)
- [Development PostgreSQL and migrations](operations/postgresql.md)
- [Disposable development and test Qdrant](operations/qdrant.md)
- [Foundation primitives](operations/foundation-primitives.md)
- [SecondContext comparison](secondcontext-comparison.md)

## Decisions

- [Architecture decision record index](adr/README.md)

Accepted ADRs are immutable in meaning. A changed decision is recorded in a superseding ADR.

## Contracts

- [Canonical domain](contracts/domain.md)
- [Ingestion and consolidation](contracts/ingestion.md)
- [Retrieval and ContextPack](contracts/retrieval.md)
- [Lifecycle and forget](contracts/lifecycle.md)
- [ThinkPixel integrations](contracts/integrations.md)
- [Persistence, jobs, and outbox](contracts/persistence.md)
- [API conventions, events, and observability](contracts/events-observability.md)
- [OpenAPI 3.1](../api/openapi/openapi.yaml)

## Evidence

- [Phase 0 traceability](phase-0-traceability.md)
- [Phase 0 validation](phase-0-validation.md)
- [Phase 1 engineering-foundation validation](evidence/phase-1-validation.md)

The flat architecture, security, and evidence paths predate the common ThinkPixel documentation shape and remain intentional until moving them provides more value than link churn. New evidence should use `docs/evidence/` once more than the current Phase 0 set exists.
