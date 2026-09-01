# Developer commands

The root `Makefile` is the stable local and CI command interface. Run `make` or `make help` to list the commands available in the current implementation phase.

The hosted workflow and required checks are documented in [continuous-integration.md](continuous-integration.md).

The aggregate gate is:

```sh
make verify
```

It runs all non-mutating repository checks:

- `format-check` verifies every Go file with `gofmt`;
- `vet-check` runs `go vet`;
- `lint-check` runs the pinned `staticcheck` tool;
- `unit-check` runs all Go tests;
- `race-check` runs all Go tests with the race detector;
- `vulnerability-check` runs the pinned Go vulnerability scanner;
- `license-check` checks runtime and test dependencies against the licenses permitted without special review by the dependency policy;
- `build-check` builds every Go package;
- `hygiene-check` rejects tracked runtime memory, credential material, Qdrant dumps, and local database artifacts;
- `phase0-validate`, `openapi-check`, and `whitespace-check` validate the architecture/contract baseline, the OpenAPI schema and generated-code drift, and changed-file whitespace.

Each name is also a focused Make target. Repository hygiene policy and remediation are documented in [repository-hygiene.md](repository-hygiene.md). The analyzer versions are pinned as Go tools in `go.mod`, so local and CI checks use the same implementations. Vulnerability results depend on the vulnerability database available when the command runs and therefore require network access unless the Go vulnerability cache is already populated.

`make openapi-generate` is intentionally separate because it updates generated transport models. The administrative CLI skeleton is documented in [cli.md](cli.md). The hardened service image is documented in [container-image.md](container-image.md). PostgreSQL development and migration commands are documented in [postgresql.md](postgresql.md). Disposable Qdrant commands are documented in [qdrant.md](qdrant.md). Later infrastructure and integration commands are added with their corresponding Phase 1 work rather than being speculative no-op targets here.
