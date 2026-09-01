# Phase 1 engineering-foundation evidence

Validation date: 2026-09-01.

Phase 1 establishes the service engineering baseline without implementing canonical
memory persistence or widening ThinkPixelMEM's ownership boundary. The delivered
foundation includes:

- a pinned Go module and explicit domain, application, port, and adapter package
  boundaries;
- typed configuration, secret references, structured redacted logging, Prometheus
  metrics, and OpenTelemetry tracing;
- typed identifiers, bounded values, digests, clocks, errors, and authenticated
  cursors;
- an HTTP service baseline with RFC 7807 errors, limits, health endpoints, metrics,
  tracing, request IDs, and graceful shutdown;
- OpenAPI validation, generation, and generated-artifact drift detection;
- developer, CI, vulnerability, license, repository-hygiene, and build gates;
- pinned disposable PostgreSQL and Qdrant development dependencies;
- CLI and migration command skeletons; and
- a hardened, non-root service container image.

## Acceptance evidence

The final baseline gate ran from a clean, committed snapshot containing all 128
repository files that existed before this evidence page was added.

| Gate | Result |
| --- | --- |
| `make verify` | Passed: formatting, vet, static analysis, unit tests, race tests, vulnerability scan, license policy, builds, repository hygiene, Phase 0 validation, OpenAPI validation, and generated-artifact drift checks. |
| `make image-check` | Passed: image build plus non-root user (`65532:65532`) and entrypoint inspection. |
| `git diff --check` | Passed in the clean committed snapshot. |
| `git status --porcelain` | Empty after both aggregate gates. |

`govulncheck` reported no called or imported vulnerabilities. It found two known
module vulnerabilities in code paths the service does not call. The license gate
passed with its documented warning that assembly in `github.com/cespare/xxhash/v2`
and `golang.org/x/sys/unix` cannot be further inspected automatically.

The Phase 0 validator passed its required structural checks. Ruby was unavailable,
so its optional secondary YAML parser check was skipped; OpenAPI validation and
generated-artifact drift checks still passed through the required Go-based gates.

The container manifest-list digest produced by the image gate was
`sha256:c22fd076bb007d19f1f4f47f1798226de23117d7bec445b06eb802a861367220`.
This is local acceptance evidence for the development image, not a release image
identifier or provenance claim.

## Reproduction

From a clean checkout with the pinned Go toolchain, Docker, and the repository's
pinned verification tools available:

```sh
make verify
make image-check
git diff --check
git status --porcelain
```

The Phase 1 implementation commit is recorded in `TODO.md` after publication so
the evidence does not attempt to embed the hash of the commit that contains it.
