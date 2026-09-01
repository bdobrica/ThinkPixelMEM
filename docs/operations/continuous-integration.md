# Continuous integration

GitHub Actions runs the repository gates for pushes to `main`, pull requests, and manual dispatches. The workflow has read-only repository permissions, cancels superseded runs for the same ref, and sets explicit job timeouts.

The repository-verification job installs the exact Go version from `.go-version`, disables automatic toolchain downloads, keeps the module graph read-only, and invokes the same `make verify` aggregate gate used locally. The hardened-image job separately invokes `make image-check` because it requires Docker and is not part of the host-only aggregate gate.

Third-party actions are pinned to immutable commit SHAs. Version comments make intentional upgrades reviewable; update both the commit and its comment together after reviewing the upstream release. The workflow does not receive or require repository secrets, publish images, deploy artifacts, or start PostgreSQL or Qdrant. Those responsibilities are introduced only with the corresponding integration and release phases.

Before opening a pull request, run:

```sh
make verify
make image-check
```

Branch protection should require both `Repository verification` and `Hardened image`. Repository administrators configure that policy in GitHub; the workflow itself does not modify repository settings.
