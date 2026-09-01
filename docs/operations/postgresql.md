# Development PostgreSQL and migrations

ThinkPixelMEM uses PostgreSQL as its only canonical store. Local development uses the official PostgreSQL 18.6 image pinned by multi-platform digest in `compose.yaml`. The service publishes only to loopback, persists data in the Compose-managed `postgres-data` volume, and exposes a health check.

Start and stop it with:

```sh
make postgres-up
make postgres-down
```

`postgres-down` preserves the database volume. `make postgres-logs` follows server logs. Docker Compose 2.20 or newer is required for the health-checked startup command.

The defaults are deliberately local-only credentials:

```text
postgresql://thinkpixelmem:thinkpixelmem_dev@127.0.0.1:5432/thinkpixelmem?sslmode=disable
```

Override the database, user, password, or published port with `TPMEM_DEV_POSTGRES_DB`, `TPMEM_DEV_POSTGRES_USER`, `TPMEM_DEV_POSTGRES_PASSWORD`, and `TPMEM_DEV_POSTGRES_PORT`. These values are development conveniences, must not be reused in a shared or production deployment, and do not replace the secret-reference rules in [configuration.md](configuration.md).

## Migrations

The PostgreSQL-specific Tern CLI is pinned as a Go tool in `go.mod`. Apply all migrations with:

```sh
make migrate
```

Migrations live in `migrations/`. Phase 2 adds the first tenant/schema migration; until then, `migrate` and `migrate-status` report the empty foundation baseline without connecting to PostgreSQL. ENG-012 only establishes the reproducible runner. Released migration files are immutable as required by the [persistence contract](../contracts/persistence.md).

`MIGRATION_DATABASE_URL` overrides the development connection string. `MIGRATION_DESTINATION` accepts a Tern destination such as `last`, `+1`, or a numeric version; it defaults to `last`. Use `make migrate-status` to inspect applied and pending versions. Supplying a credential-bearing URL on a command line can expose it through process inspection, so shared and production automation should use a protected Tern configuration or environment-based secret injection rather than this developer target.

## Dependency review

Tern v2.4.3 is used only as a development/build tool. It is PostgreSQL-specific, uses the same pgx driver family planned for the PostgreSQL adapter, supports ordered up/down SQL migrations, is actively maintained, and is straightforward to replace because migrations remain ordinary SQL. Its upstream license is MIT. The official PostgreSQL image carries the PostgreSQL License and is pinned by digest; neither dependency crosses into the domain layer. `go mod verify`, the repository vulnerability scan, and the runtime/test license gate provide the automated evidence described in the dependency policy; tool and image licensing is reviewed here because the package-focused license gate does not cover them.
