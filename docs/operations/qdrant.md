# Disposable development and test Qdrant

Qdrant is ThinkPixelMEM's reference dense-and-sparse retrieval projection. It is not canonical memory state: PostgreSQL remains authoritative, and Qdrant content must always be safe to delete and rebuild under [ADR-0002](../adr/0002-canonical-state-and-projections.md).

## Local lifecycle

Docker Compose 2.20 or newer is required. Start a healthy Qdrant instance with:

```sh
make qdrant-up
```

The service exposes HTTP on `http://127.0.0.1:6333` and gRPC on `127.0.0.1:6334`. Both ports bind only to loopback. Follow logs or remove the service with:

```sh
make qdrant-logs
make qdrant-down
```

The container stores `/qdrant/storage` in an in-memory `tmpfs`. Stopping and removing the container therefore discards every collection, point, snapshot, and projection artifact. Tests and development workflows must recreate their collections and must never depend on data surviving `make qdrant-down`.

Override host ports when the defaults are occupied:

```sh
TPMEM_DEV_QDRANT_HTTP_PORT=16333 \
TPMEM_DEV_QDRANT_GRPC_PORT=16334 \
make qdrant-up
```

These variables affect local host bindings only. They do not configure a shared or production deployment.

## Version and dependency review

Compose pins the official `qdrant/qdrant:v1.19.0` multi-platform image by digest. Qdrant is Apache-2.0 licensed and is admitted as development/test infrastructure for the retrieval adapter selected by the accepted architecture. It does not enter the domain model or make Qdrant canonical. Provider-specific integration remains behind the `RetrievalIndex` port, which is implemented in Phase 4.

The local health check verifies that Qdrant's HTTP listener accepts connections. Application and integration tests remain responsible for checking API behavior and feature compatibility. Registry digest inspection and an image startup check provide local admission evidence; release automation will add repeatable image vulnerability and provenance evidence in its assigned phase.
