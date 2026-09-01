# Dependency, source, and license policy

This policy applies to source code, generated code, build tools, Go modules, container images, and other third-party material distributed with or used to build ThinkPixelMEM. The repository is licensed under Apache-2.0; dependency choices must preserve the ability to build, test, operate, and redistribute the service under that license.

## Dependency admission

The standard library and existing dependencies are preferred when they are sufficient. A new dependency must have a concrete repository-local use, a maintained upstream, a disclosed license, and no narrower alternative that provides the required security or interoperability behavior. The change introducing it must document the purpose and evaluate maintenance activity, release provenance, known vulnerabilities, transitive dependencies, and replacement cost.

Go dependencies must use canonical module paths and be recorded in `go.mod` and `go.sum`. Versions are pinned by the module graph; production images and external tools must be pinned to an immutable version or digest. Floating branches, unversioned downloads, and runtime installation of build dependencies are not permitted in release builds. `replace` directives, forks, pseudo-versions, vendored source, and dependencies fetched directly from a VCS host require an explicit rationale in the introducing change. Local-path `replace` directives must not be committed.

Dependencies are reviewed monthly, before a release, and promptly when a relevant security advisory appears. A dependency with an unpatched applicable critical or high-severity vulnerability must be upgraded, replaced, removed, or covered by a documented time-bounded risk acceptance before release. Removal must also remove unused transitive modules and associated configuration.

## Source provenance

Third-party source must come from the canonical upstream project or an explicitly documented fork over an authenticated transport. Release signatures, checksums, or immutable digests must be verified when upstream publishes them. Download-and-execute installation pipelines are not release inputs unless the downloaded artifact is pinned and its integrity is verified before execution.

Copied or adapted source and generated artifacts must identify their origin, upstream version or commit, applicable license, and local modifications in a nearby notice or durable repository document. Generated files must identify the generator and reproducible command. Contributors must not add code, data, fixtures, model output, or other material whose provenance or redistribution rights are unknown.

Private modules and private source archives must not embed credentials in module paths, configuration, or committed files. Access uses developer or CI secret stores. Another ThinkPixel repository's `internal` packages or database are never a dependency; cross-component behavior uses versioned wire contracts and stable identifiers.

## License admission

License review covers direct and transitive dependencies, copied source, generated output that carries licensing terms, embedded assets, build tools included in distributed artifacts, and container base images.

- Apache-2.0, MIT, BSD-2-Clause, BSD-3-Clause, and ISC material is normally acceptable when its notice and attribution conditions are preserved.
- MPL-2.0 and other file-scoped or weak-copyleft licenses require maintainer review of distribution and source-availability obligations before admission.
- GPL, AGPL, LGPL, SSPL, Business Source License, non-commercial, field-of-use-restricted, source-available, custom, or unknown licenses require explicit maintainer and legal approval before use. Approval must record the exact component, version, use, distribution impact, obligations, and reviewer.
- Strong-copyleft or network-copyleft code must not be linked into or combined with distributed service binaries without that approval. Development-only use is still reviewed because CI images, generated artifacts, and redistribution may trigger obligations.

Required copyright, attribution, and license texts must accompany distributed artifacts. A dependency's license change is treated as a new admission review. Package metadata is evidence, not conclusive proof; ambiguous or conflicting licensing blocks release until resolved.

## Verification and exceptions

Every dependency change must leave `go mod verify` and the repository verification gate passing. The gate scans Go packages for known vulnerabilities and checks runtime and test dependencies against the licenses permitted without special review above. A passing automated license classification is evidence, not a substitute for reviewing ambiguous metadata, copied source, generated output, tools, images, and other release inputs. Release SBOM and provenance enforcement remain assigned to the release-engineering phase.

An exception must be narrow, time-bounded, owned by a named maintainer, and record the reason, affected version and artifacts, risk, compensating controls, expiry, and removal or re-review plan. Exceptions cannot waive credential handling, ThinkPixel component boundaries, or license obligations.
