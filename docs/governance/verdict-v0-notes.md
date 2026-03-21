# Verdict v0 Notes

## Contract identity

- schema family: `verdict`
- schema version: `v0`

## Boundary

Verdict v0 is the final release-level decision artifact.

It sits above:
- Raw Result
- Normalized Result
- Policy Evaluation Result

It must consume their outcomes through summary and reference semantics, not by
embedding their full payloads.

## Key design rules

### Stable verdict vocabulary
The vocabulary is intentionally small and fixed in v0:
- approved
- warning
- rejected

### Policy summary is release-level only
The summary exists to support final decision consumption and should remain small.

### Score is out of scope
No score field is included in v0 because score semantics are not yet stable
enough to justify a canonical contract field.

### Traceability by reference
Traceability must point to lower-layer artifacts rather than duplicating them.

## Change discipline

Breaking changes:
- changing verdict enum incompatibly
- removing required decision context
- replacing summary-only semantics with embedded lower-layer payloads
- changing traceability meaning incompatibly

Non-breaking changes:
- adding optional traceability fields
- adding optional context fields that preserve compatibility
- clarifying descriptions while preserving payload semantics

Editorial-only changes:
- examples
- comments
- formatting
- wording