.PHONY: verify

verify:
	./scripts/validate-phase0.sh
	git diff --check
