# Phase 0 validation evidence

Validation date: 2026-08-30.

Observed results:

- `./scripts/validate-phase0.sh`: `Phase 0 structural validation passed`;
- Redocly CLI: API description valid, with advisory warnings that generic `default` problem responses are not explicit `4XX` responses;
- `git diff --check`: passed;
- implementation commit: `7f45aee` (`docs: define Phase 0 architecture and contracts`).

Run:

```sh
make verify
npx --yes @redocly/cli lint api/openapi/openapi.yaml
```

`make verify` is the stable local and CI entry point. In addition to the Phase 0 structural checks, OpenAPI validation and generated-artifact drift, and changed-file whitespace, it runs the repository's Go format, vet, lint, unit, race, vulnerability, license, and build gates. The Phase 0 script also parses the OpenAPI YAML when Ruby is available. The Redocly lint remains a separate optional network-dependent check. See the [developer command reference](operations/development.md) for focused targets.

Acceptance requires:

- every ARC-001–ARC-090 item maps to a non-empty artifact;
- Markdown code fences are balanced;
- OpenAPI parses as YAML and includes the required API conventions;
- the Phase 0 checklist identifiers remain complete and ordered;
- whitespace validation passes;
- a reviewer confirms the Phase 0 exit gate: ownership, authority, trust, temporal semantics, and deletion are unambiguous.

The validation script provides structural evidence, not implementation proof for later phases. `TODO.md` records the implementation commit SHA for every completed Phase 0 item.
