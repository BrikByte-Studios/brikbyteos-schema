# Policy Evaluation Result v0 Notes

## Contract identity

- schema family: `policy-evaluation-result`
- schema version: `v0`

## Boundary

Policy Evaluation Result v0 is the canonical contract between:
- normalized evidence
- later verdict aggregation

It is not:
- a raw execution contract
- a normalized evidence contract
- a verdict contract

## Key design rules

### Stable policy identity
A policy result must always identify:
- stable policy id
- policy version

A human-readable name may exist, but it cannot be the sole identity.

### Fixed status vocabulary
The status vocabulary is intentionally fixed in v0:
- passed
- failed
- not_applicable
- error

### Controlled explanation
The explanation model is deliberately minimal in v0:
- one concise reason string
- useful for CLI and audit presentation
- no narrative tree
- no AI content

### Evidence traceability
Policy results cite normalized evidence through references rather than embedding
all upstream payloads inline.

## Change discipline

Breaking changes:
- changing status vocabulary incompatibly
- removing required identity or evidence fields
- weakening policy identity to name-only
- changing evidence reference semantics incompatibly

Non-breaking changes:
- adding optional metadata fields
- clarifying descriptions while keeping semantics intact
- adding optional policy category extensions that preserve compatibility

Editorial-only changes:
- comments
- examples
- wording
- formatting