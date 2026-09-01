# Runtime configuration

ThinkPixelMEM loads typed process configuration from four layers, in increasing precedence:

1. compiled safe defaults;
2. an optional strict JSON file selected by `--config`;
3. `TPMEM_*` environment variables;
4. command-line flags.

Unknown JSON fields, `TPMEM_*` variables, flags, positional arguments, malformed values, repeated environment variables, trailing JSON, non-regular files, and configuration files larger than 1 MiB are rejected. Invalid configuration prevents startup.

## Initial fields

| JSON path | Environment | Flag |
| --- | --- | --- |
| `mode` | `TPMEM_MODE` | `--mode` |
| `http.address` | `TPMEM_HTTP_ADDRESS` | `--http-address` |
| `http.read_header_timeout` | `TPMEM_HTTP_READ_HEADER_TIMEOUT` | `--http-read-header-timeout` |
| `http.read_timeout` | `TPMEM_HTTP_READ_TIMEOUT` | `--http-read-timeout` |
| `http.write_timeout` | `TPMEM_HTTP_WRITE_TIMEOUT` | `--http-write-timeout` |
| `http.idle_timeout` | `TPMEM_HTTP_IDLE_TIMEOUT` | `--http-idle-timeout` |
| `http.shutdown_timeout` | `TPMEM_HTTP_SHUTDOWN_TIMEOUT` | `--http-shutdown-timeout` |
| `http.max_header_bytes` | `TPMEM_HTTP_MAX_HEADER_BYTES` | `--http-max-header-bytes` |
| `http.max_body_bytes` | `TPMEM_HTTP_MAX_BODY_BYTES` | `--http-max-body-bytes` |
| `database.url` | `TPMEM_DATABASE_URL` | `--database-url` |
| `database.connect_timeout` | `TPMEM_DATABASE_CONNECT_TIMEOUT` | `--database-connect-timeout` |
| `database.health_timeout` | `TPMEM_DATABASE_HEALTH_TIMEOUT` | `--database-health-timeout` |
| `database.statement_timeout` | `TPMEM_DATABASE_STATEMENT_TIMEOUT` | `--database-statement-timeout` |
| `database.lock_timeout` | `TPMEM_DATABASE_LOCK_TIMEOUT` | `--database-lock-timeout` |
| `database.max_connection_lifetime` | `TPMEM_DATABASE_MAX_CONNECTION_LIFETIME` | `--database-max-connection-lifetime` |
| `database.max_connection_idle_time` | `TPMEM_DATABASE_MAX_CONNECTION_IDLE_TIME` | `--database-max-connection-idle-time` |
| `database.min_connections` | `TPMEM_DATABASE_MIN_CONNECTIONS` | `--database-min-connections` |
| `database.max_connections` | `TPMEM_DATABASE_MAX_CONNECTIONS` | `--database-max-connections` |
| `log.level` | `TPMEM_LOG_LEVEL` | `--log-level` |
| `telemetry.mode` | `TPMEM_TELEMETRY_MODE` | `--telemetry-mode` |
| `telemetry.endpoint` | `TPMEM_TELEMETRY_ENDPOINT` | `--telemetry-endpoint` |
| `telemetry.service_name` | `TPMEM_TELEMETRY_SERVICE_NAME` | `--telemetry-service-name` |
| `telemetry.sample_ratio` | `TPMEM_TELEMETRY_SAMPLE_RATIO` | `--telemetry-sample-ratio` |

Durations use Go duration syntax. Modes are `development`, `test`, and `production`. Defaults bind HTTP to loopback, limit bodies to the contract default of 1 MiB, disable external telemetry, and contain no database credential. Production requires `database.url`.

## Secret references

Secret-bearing fields accept only `env:VARIABLE_NAME` or `file:/clean/absolute/path`. Inline URLs containing credentials, relative paths, and unknown schemes are rejected. Loading retains an opaque reference; the owning adapter resolves it when needed. File secrets must be non-empty regular files no larger than 1 MiB.

References, resolved secrets, and complete configuration values redact ordinary string, debug, and JSON output. Safe output exposes only whether a reference is configured and whether its source is an environment variable or file—not its target or value.

```json
{
  "mode": "production",
  "database": {
    "url": "file:/run/secrets/thinkpixelmem-database-url"
  },
  "log": {
    "level": "info"
  }
}
```

Qdrant and ThinkPixel integration settings will be introduced with their owning adapter contracts rather than accepted as untyped extension data. Provider credentials remain outside MEM configuration.
