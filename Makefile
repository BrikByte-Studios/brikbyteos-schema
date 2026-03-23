.PHONY: validate-examples test test-go

## Validate canonical schema v0 example payloads.
validate-examples:
	./scripts/validate-examples.sh

## Run Go schema validation tests.
test-go:
	go test ./...

## Run full repository validation baseline.
test: validate-examples test-go