.DEFAULT_GOAL := help

MIGRATION_DATABASE_URL ?= postgresql://thinkpixelmem:thinkpixelmem_dev@127.0.0.1:5432/thinkpixelmem?sslmode=disable
MIGRATION_DESTINATION ?= last

.PHONY: build-check cli-build format-check help hygiene-check image-build image-check license-check lint-check migrate migrate-status openapi-check openapi-generate openapi-validate phase0-validate postgres-down postgres-logs postgres-up qdrant-down qdrant-logs qdrant-up race-check unit-check vet-check verify vulnerability-check whitespace-check

help:
	@printf '%s\n' \
		'ThinkPixelMEM developer commands:' \
		'  make verify              Run the aggregate repository verification gate' \
		'  make format-check        Check Go formatting' \
		'  make vet-check           Run go vet' \
		'  make lint-check          Run staticcheck' \
		'  make unit-check          Run unit tests' \
		'  make race-check          Run tests with the race detector' \
		'  make vulnerability-check Scan Go dependencies for known vulnerabilities' \
		'  make license-check       Enforce the Go dependency license allowlist' \
		'  make build-check         Build every Go package' \
		'  make hygiene-check       Reject tracked secrets and local data artifacts' \
		'  make cli-build           Build thinkpixelmemctl' \
		'  make image-build         Build the hardened service image' \
		'  make image-check         Build and inspect image hardening' \
		'  make postgres-up         Start the development PostgreSQL service' \
		'  make postgres-down       Stop the development PostgreSQL service' \
		'  make postgres-logs       Follow development PostgreSQL logs' \
		'  make migrate             Apply PostgreSQL migrations' \
		'  make migrate-status      Show PostgreSQL migration status' \
		'  make qdrant-up           Start disposable development Qdrant' \
		'  make qdrant-down         Remove disposable development Qdrant' \
		'  make qdrant-logs         Follow development Qdrant logs' \
		'  make phase0-validate     Validate the architecture and contract baseline' \
		'  make openapi-check       Validate OpenAPI and detect generated-code drift' \
		'  make openapi-validate    Validate the canonical OpenAPI document' \
		'  make openapi-generate    Regenerate OpenAPI transport models' \
		'  make whitespace-check    Check tracked changes for whitespace errors'

format-check:
	@files="$$(gofmt -l .)"; if [ -n "$$files" ]; then printf '%s\n' 'Go files need formatting:' "$$files"; exit 1; fi

vet-check:
	go vet ./...

lint-check:
	go tool staticcheck ./...

unit-check:
	go test ./...

race-check:
	go test -race ./...

vulnerability-check:
	go tool govulncheck ./...

license-check:
	go tool go-licenses check --include_tests --allowed_licenses=Apache-2.0,MIT,BSD-2-Clause,BSD-3-Clause,ISC ./...

build-check:
	go build ./...

hygiene-check:
	go run ./internal/tools/repositoryhygiene

cli-build:
	@mkdir -p .cache/bin
	go build -o .cache/bin/thinkpixelmemctl ./cmd/thinkpixelmemctl

image-build:
	docker build --tag thinkpixelmem:dev .

image-check: image-build
	@test "$$(docker image inspect thinkpixelmem:dev --format '{{.Config.User}}')" = '65532:65532'
	@test "$$(docker image inspect thinkpixelmem:dev --format '{{json .Config.Entrypoint}}')" = '["/thinkpixelmem"]'
	@test "$$(docker image inspect thinkpixelmem:dev --format '{{json .Config.Cmd}}')" = 'null'

postgres-up:
	docker compose up --detach --wait postgres

postgres-down:
	docker compose down

postgres-logs:
	docker compose logs --follow postgres

migrate:
	@set -- migrations/*.sql; if [ ! -e "$$1" ]; then printf '%s\n' 'No migrations found; Phase 2 adds the initial schema migration.'; exit 0; fi; \
		go tool tern migrate --migrations migrations --conn-string '$(MIGRATION_DATABASE_URL)' --destination '$(MIGRATION_DESTINATION)'

migrate-status:
	@set -- migrations/*.sql; if [ ! -e "$$1" ]; then printf '%s\n' 'No migrations found; database is at the foundation baseline.'; exit 0; fi; \
		go tool tern status --migrations migrations --conn-string '$(MIGRATION_DATABASE_URL)'

qdrant-up:
	docker compose up --detach --wait qdrant

qdrant-down:
	docker compose rm --force --stop qdrant

qdrant-logs:
	docker compose logs --follow qdrant

openapi-generate:
	./scripts/openapi.sh generate

openapi-validate:
	./scripts/openapi.sh validate

openapi-check:
	./scripts/openapi.sh check

phase0-validate:
	./scripts/validate-phase0.sh

whitespace-check:
	git diff --check

verify: format-check vet-check lint-check unit-check race-check vulnerability-check license-check build-check hygiene-check phase0-validate openapi-check whitespace-check
