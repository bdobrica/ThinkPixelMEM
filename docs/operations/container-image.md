# Service container image

The root `Dockerfile` builds the `thinkpixelmem` service as a statically linked Linux binary and copies only that binary and the system CA bundle into a `scratch` runtime image. The runtime has no shell, package manager, compiler, source tree, or writable application data. It runs as numeric UID and GID `65532:65532`; deployments must not override that identity or grant privilege escalation.

The Dockerfile frontend and the supported Go 1.26.7 Debian builder are pinned by immutable multi-platform digests. The runtime has no base-image packages. The CA bundle is retained so configured HTTPS telemetry can validate server certificates. `.dockerignore` admits only the module metadata and service source required by the build, keeping repository metadata, caches, documentation, local databases, and unrelated artifacts out of the build context.

Build and inspect the image with:

```sh
make image-build
make image-check
```

The development image tag is `thinkpixelmem:dev`. It listens on all container interfaces at port 8080. Configuration still follows [configuration.md](configuration.md); production mode requires a database secret reference, normally mounted read-only below `/run/secrets`. The image does not contain credentials or a default database URL.

```sh
docker run --rm --read-only --cap-drop=ALL \
  --security-opt=no-new-privileges \
  --publish 127.0.0.1:8080:8080 \
  thinkpixelmem:dev
```

`/livez` reports process liveness. During the Phase 1 foundation, `/readyz` deliberately returns `503` because the canonical PostgreSQL readiness adapter is not implemented until Phase 2. Container orchestration should use HTTP probes rather than an in-image command: the hardened runtime intentionally contains no probe utility.

Release automation must assign an immutable image tag/digest and add vulnerability, SBOM, and provenance evidence. The local `:dev` tag is not a release identifier.
