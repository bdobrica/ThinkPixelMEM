# Foundation primitives

Phase 1 application and adapter code shares a small set of security-sensitive primitives.

- `internal/domain` owns canonical UUIDv7 parsing/generation, distinct identifier types, stable coded errors, bounded UTF-8 strings, and SHA-256 digests.
- `internal/ports/clock` supplies trusted server time through an injectable interface. Production code uses `clock.System`; tests provide deterministic clocks.
- `internal/security.CursorCodec` signs opaque pagination state with HMAC-SHA-256. Cursors are bound to a collection purpose, expire within 24 hours, and are limited to the OpenAPI maximum of 2048 bytes.

UUID parsing accepts only lowercase canonical UUIDv7 strings. ID generation uses the injected clock for the 48-bit Unix-millisecond timestamp and cryptographic entropy for the remaining random bits. Domain-specific ID types must not be interchanged by conversion outside parsing and persistence boundaries.

Cursor keys must contain at least 32 bytes and must come from an operator-managed secret provider. Keys and decoded cursor state must not be logged. Services should use a distinct purpose string for each pagination query shape, and deployments must coordinate key rotation with the maximum cursor lifetime.

Bounded strings count Unicode code points and reject malformed UTF-8. Callers remain responsible for applying the limits defined by the relevant OpenAPI or JSON Schema field. Digests use lowercase hexadecimal SHA-256 and establish content integrity, not identity, trust, or authorization.
