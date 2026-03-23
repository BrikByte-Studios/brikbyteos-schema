.PHONY: validate-examples test

## Validate canonical schema v0 example payloads.
validate-examples:
	./scripts/validate-examples.sh

test: validate-examples