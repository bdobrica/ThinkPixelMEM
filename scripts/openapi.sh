#!/usr/bin/env bash
set -euo pipefail

repo_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_dir"

spec="api/openapi/openapi.yaml"
config="api/openapi/oapi-codegen.yaml"
generated="internal/adapters/http/openapi/types.gen.go"

validate() {
  go run ./internal/tools/openapicheck "$spec"
}

generate_to() {
  local output="$1"
  go tool oapi-codegen --config "$config" -o "$output" "$spec"
}

case "${1:-check}" in
  generate)
    validate
    mkdir -p "$(dirname "$generated")"
    generate_to "$generated"
    ;;
  validate)
    validate
    ;;
  check)
    validate
    temporary="$(mktemp "${TMPDIR:-/tmp}/thinkpixelmem-openapi.XXXXXX.go")"
    trap 'rm -f "$temporary"' EXIT
    generate_to "$temporary"
    if ! cmp -s "$temporary" "$generated"; then
      echo "generated OpenAPI types are stale; run: make openapi-generate" >&2
      diff -u "$generated" "$temporary" || true
      exit 1
    fi
    echo "OpenAPI generated-artifact drift check passed"
    ;;
  *)
    echo "usage: $0 [generate|validate|check]" >&2
    exit 2
    ;;
esac
