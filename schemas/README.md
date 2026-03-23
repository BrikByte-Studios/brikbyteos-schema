# Schemas

This directory contains canonical BrikByteOS schema families and their explicit
version namespaces.

Recommended pattern:

```text
schemas/<schema-family>/<schema-version>/
```

Examples:
- `schemas/raw-result/v0/`
- `schemas/normalized-result/v0/`
- `schemas/verdict/v0/`

## Canonical examples

Each schema family/version directory contains an `examples/valid.minimal.json`
file. These are the authoritative Phase 0 valid sample payload fixtures for
that schema version.

They are used for:
- validation
- documentation
- onboarding
- future smoke-test reference inputs

Validate them with:

```bash
make test
```