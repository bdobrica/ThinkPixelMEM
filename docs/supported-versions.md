# Supported versions

Baseline reviewed 2026-08-30. Production images pin an exact patch/digest; the contract below states compatibility ranges.

| Component | RC baseline | Policy |
| --- | --- | --- |
| Go | 1.26.x | Build with latest security patch in 1.26; support the current and previous Go minor after CI qualification. |
| PostgreSQL | 18.x; compatibility target 17.x | Pin latest supported minor; never run an end-of-life major; test upgrade from the previous supported major. |
| Qdrant | 1.19.x | Pin the exact current 1.19 patch image digest in Phase 1; requires named dense/sparse vectors, payload filters, Query API prefetch/fusion, snapshots, and scroll. |
| OpenAPI | 3.1.0 / JSON Schema 2020-12 | `api/openapi/openapi.yaml` is canonical. |
| OIDC/JWT | OIDC Core 1.0, JWT/JWS implementations with issuer metadata | Pin library versions in Phase 1; reject `none` and unconfigured algorithms. |
| W3C Trace Context | Recommendation (trace-context-1) | Propagate `traceparent`/`tracestate`; pin telemetry libraries in Phase 1. |
| RFC 7807 | Problem Details | Use `application/problem+json`; a later RFC 9457 alignment may be additive. |
| ThinkPixel contracts | `v1alpha1` | Exact schema version sent on every integration; breaking change requires a new version. |
| LLMGW / GR | `v1alpha1` adapter contracts | Service deployment versions are negotiated and pinned by integration tests. |

PostgreSQL 18 is supported through November 2030 and 17 through November 2029 under the PostgreSQL five-year policy. Qdrant 1.19 is the selected RC line; hybrid/multi-stage Query API has been available since 1.10. References: [Go 1.26 release notes](https://go.dev/doc/go1.26), [PostgreSQL versioning policy](https://www.postgresql.org/support/versioning/), [Qdrant 1.19 release](https://github.com/qdrant/qdrant/releases/tag/v1.19.0), and [Qdrant hybrid query documentation](https://qdrant.tech/documentation/search/hybrid-queries/).

Version review occurs monthly and before every release. Security fixes take precedence over the normal cadence. A dependency reaches unsupported status when its upstream support ends or required security fixes are unavailable; releases cannot ship on an unsupported baseline.
