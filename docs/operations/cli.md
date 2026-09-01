# Command-line client

`thinkpixelmemctl` is the administrative command-line client for ThinkPixelMEM. The Phase 1 skeleton establishes the planned command hierarchy:

```text
memory-space create|describe
memory get|inspect|correct|forget|quarantine
context retrieve
profile inspect
ingestion status
index rebuild
```

Build it into `.cache/bin/thinkpixelmemctl` with `make cli-build` or inspect it without installing it:

```sh
go run ./cmd/thinkpixelmemctl help
go run ./cmd/thinkpixelmemctl memory --help
go run ./cmd/thinkpixelmemctl version
```

The operation commands intentionally return a non-zero “API operation is not implemented” result during the foundation phase. Later phases add request arguments and generated API-client calls alongside the corresponding endpoints. The CLI will communicate only through the public, versioned HTTP API; it must not access PostgreSQL, Qdrant, another component's database, or service `internal` implementation state directly. Administrative API authentication remains distinct from Run-scoped MemoryGrants.
