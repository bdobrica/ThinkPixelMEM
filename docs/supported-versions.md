# Supported versions

Baseline reviewed 2026-09-01. This page records the Phase 1 build and integration baseline; it does not widen any versioned wire contract. "Pinned" means the repository selects an exact tool or image version. "Qualified" means that version has passed the repository gate or the named integration suite. Versions described only as targets are not yet release-qualified.

| Component | Phase 1 baseline | Qualification and policy |
| --- | --- | --- |
| Go | 1.26.7 pinned | `.go-version`, `go.mod`, CI, and the image builder agree on the exact toolchain. A second Go minor becomes supported only after CI qualification. |
| PostgreSQL | 18.6 pinned for development; 17.x compatibility target | Pin the current supported minor within each qualified major; never ship on an end-of-life major; qualify upgrade from the previous supported major. Runtime qualification begins with the PostgreSQL adapter in Phase 2. |
| Qdrant | 1.19.0 image pinned by digest for development/test | The adapter requires named dense/sparse vectors, payload filters, Query API prefetch/fusion, snapshots, and scroll. Runtime feature qualification begins with the retrieval adapter in Phase 4. |
| OpenAPI / JSON Schema | OpenAPI 3.1.0 / JSON Schema 2020-12 | `api/openapi/openapi.yaml` is canonical and is validated by `make openapi-check`. |
| OIDC / JWT / JWS | OIDC Core 1.0 and issuer metadata | Library and algorithm support is not qualified until the Phase 7 authentication work. Implementations reject `none` and algorithms not explicitly configured. |
| W3C Trace Context | Recommendation `trace-context-1` | The HTTP baseline propagates `traceparent` and `tracestate`; telemetry dependencies are pinned in `go.mod`. |
| Problem Details | RFC 7807 | Responses use `application/problem+json`; later RFC 9457 alignment may be additive. |
| ThinkPixel wire contracts | `v1alpha1` | Every integration sends an exact schema version. Breaking changes require a new version; additive compatibility remains governed by each contract. |
| LLMGW / GR service deployments | `v1alpha1` adapter-contract target | No service release is qualified yet. Deployment versions will be negotiated and pinned by Phase 3 integration tests. |

## Repository pins and evidence

| Version concern | Repository source of truth | Current evidence |
| --- | --- | --- |
| Go build toolchain | `.go-version`, `go.mod`, `Dockerfile` | `make verify` and CI |
| PostgreSQL development image | `compose.yaml` | `docker compose config --quiet`; adapter tests when implemented |
| Qdrant development/test image | `compose.yaml` | `docker compose config --quiet`; adapter feature tests when implemented |
| Go modules and build tools | `go.mod`, `go.sum` | `go mod verify`, vulnerability and license gates in `make verify` |
| Public API schema | `api/openapi/openapi.yaml` | `make openapi-check` |
| ThinkPixel integration schemas | `api/schemas/` and `docs/contracts/` | schema validation and compatibility tests added with each adapter |

An image tag is descriptive; its digest is the immutable pin. A documentation baseline never overrides these machine-readable sources. If they disagree, the mismatch must be corrected and requalified before release.

## Review and support lifecycle

PostgreSQL 18 is supported through November 2030 and 17 through November 2029 under PostgreSQL's five-year policy. Qdrant 1.19 is the selected RC line; its hybrid/multi-stage Query API has been available since 1.10. Upstream references are the [Go release history](https://go.dev/doc/devel/release), [PostgreSQL versioning policy](https://www.postgresql.org/support/versioning/), [Qdrant 1.19 release](https://github.com/qdrant/qdrant/releases/tag/v1.19.0), and [Qdrant hybrid query documentation](https://qdrant.tech/documentation/search/hybrid-queries/).

Maintainers review this page, repository pins, upstream support status, and security advisories monthly and before every release. Security fixes take precedence over the normal cadence. A dependency becomes unsupported when upstream support ends, required security fixes are unavailable, or its qualification evidence no longer passes. Releases cannot ship on an unsupported or merely targeted baseline.

An upgrade changes the machine-readable pin first, updates this page in the same change, and reruns the relevant qualification evidence. Major datastore or wire-contract upgrades also require migration/compatibility evidence and, when architectural meaning changes, a superseding ADR.
