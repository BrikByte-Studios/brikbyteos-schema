# Normalized Result v0 Notes

## Contract identity

- schema family: `normalized-result`
- schema version: `v0`

## Boundary

Normalized Result v0 is the canonical contract between:
- Raw Result execution evidence
- downstream Policy Evaluation / Verdict layers

## Change discipline

Breaking changes:
- removing required normalized fields
- changing issue severity semantics incompatibly
- changing metrics model in a way that breaks consumers

Non-breaking changes:
- adding optional metadata-like references
- clarifying descriptions without changing payload meaning
- adding optional fields that preserve existing compatibility

Editorial-only changes:
- comments
- examples
- wording
- formatting

## Key rule

Normalized Result must remain canonical and adapter-agnostic at the core.
It must not become a place to dump:
- raw tool-native payloads
- policy outcomes
- verdicts
- AI interpretation