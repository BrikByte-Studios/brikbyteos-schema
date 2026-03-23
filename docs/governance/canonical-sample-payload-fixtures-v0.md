# Canonical Sample Payload Fixtures for Schema v0

## Purpose

This document defines the authoritative Phase 0 valid sample payload fixtures for
the current BrikByteOS schema families.

These fixtures are canonical:
- for validation
- for onboarding
- for documentation
- for future smoke-test reference inputs

They are not incidental examples.

## Canonical ownership

The authoritative valid fixture for each schema family lives next to the schema
it validates:

- `schemas/raw-result/v0/examples/valid.minimal.json`
- `schemas/normalized-result/v0/examples/valid.minimal.json`
- `schemas/policy-evaluation-result/v0/examples/valid.minimal.json`
- `schemas/verdict/v0/examples/valid.minimal.json`
- `schemas/audit-manifest/v0/examples/valid.minimal.json`

## Why this location is canonical

This repository already organizes schemas by family and version. Keeping each
canonical valid example adjacent to its schema:

- reduces duplication
- improves discoverability
- reduces drift risk
- keeps version ownership explicit

## Phase 0 scope rules

Phase 0 includes:
- one valid canonical fixture per schema family
- realistic but synthetic values
- internally coherent references where applicable

Phase 0 excludes:
- invalid fixtures
- large scenario libraries
- multiple variants per schema family
- cross-repo duplicated fixture copies

## Validation

Run:

```bash
make test
```
This validates:
- each canonical fixture against its corresponding schema
- basic coherence across the contract chain