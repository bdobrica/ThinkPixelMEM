# OpenAPI development workflow

[`api/openapi/openapi.yaml`](../../api/openapi/openapi.yaml) is the canonical public API contract. Generated Go types are transport models owned by the HTTP adapter; they do not replace domain types or establish identity, trust, or authority.

The repository pins `oapi-codegen` as a Go tool. Its Apache-2.0 license is compatible with the repository, and it is used because it supports reproducible Go transport models from the canonical OpenAPI document. The generated file identifies the tool and version in its header.

Run:

```sh
make openapi-validate
make openapi-generate
make openapi-check
```

Validation resolves local references, rejects external references, and applies OpenAPI schema validation. Generation writes `internal/adapters/http/openapi/types.gen.go`. The drift check regenerates into a temporary file and fails when committed output differs. Contract changes must update the OpenAPI document first, regenerate types, and include compatibility, documentation, and test updates required by the affected behavior.
