#!/usr/bin/env bash
set -euo pipefail

repo_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_dir"

required=(
  docs/adr/template.md
  docs/adr/0001-service-boundaries.md
  docs/adr/0002-canonical-state-and-projections.md
  docs/adr/0003-trust-and-metadata.md
  docs/adr/0004-temporal-revisions.md
  docs/architecture.md
  docs/threat-model.md
  docs/contracts/domain.md
  docs/contracts/ingestion.md
  docs/contracts/retrieval.md
  docs/contracts/lifecycle.md
  docs/contracts/integrations.md
  docs/contracts/persistence.md
  docs/contracts/events-observability.md
  docs/supported-versions.md
  docs/secondcontext-comparison.md
  docs/phase-0-traceability.md
  api/openapi/openapi.yaml
)

for path in "${required[@]}"; do
  test -s "$path" || { echo "missing or empty: $path" >&2; exit 1; }
done

python3 - <<'PY'
from pathlib import Path
import re

for path in [Path("README.md"), Path("PLAN.md"), *Path("docs").rglob("*.md")]:
    text = path.read_text(encoding="utf-8")
    if text.count("```") % 2:
        raise SystemExit(f"unbalanced Markdown fences: {path}")

todo = Path("TODO.md").read_text(encoding="utf-8")
ids = re.findall(r"ARC-(\d{3})", todo)
expected = [f"{n:03d}" for n in range(1, 91)]
if ids[:90] != expected:
    raise SystemExit("TODO Phase 0 ARC identifiers are missing or out of order")

trace = Path("docs/phase-0-traceability.md").read_text(encoding="utf-8")
for n in expected:
    if f"ARC-{n}" not in trace and not any(
        int(a) <= int(n) <= int(b)
        for a, b in re.findall(r"ARC-(\d{3})[–-](\d{3})", trace)
    ):
        raise SystemExit(f"missing traceability for ARC-{n}")

spec = Path("api/openapi/openapi.yaml").read_text(encoding="utf-8")
for token in ("openapi: 3.1.0", "application/problem+json", "Idempotency-Key", "text/event-stream", "openIdConnect"):
    if token not in spec:
        raise SystemExit(f"OpenAPI convention missing: {token}")
print("Phase 0 structural validation passed")
PY

if command -v ruby >/dev/null 2>&1; then
  ruby -e 'require "yaml"; YAML.safe_load_file("api/openapi/openapi.yaml", aliases: true); puts "OpenAPI YAML parse passed"'
else
  echo "ruby unavailable; skipped YAML parser check"
fi
