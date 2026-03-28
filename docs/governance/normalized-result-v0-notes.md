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

## Adapter Extension Field Strategy

Normalized Result v0.1 supports exactly one optional top-level extension field:

- `extensions`

### Rules

1. Core schema fields remain top-level and stable.
2. Adapter-specific data must appear only under:
   - `extensions.<adapter>`
3. Adapter namespace format:
   - lowercase
   - alphanumeric plus hyphen only
   - examples: `jest`, `playwright`, `k6`, `trivy`
4. Consumers may ignore `extensions` entirely without losing core semantics.
5. Alternate top-level extension-like fields such as:
   - `details`
   - `adapterSpecific`
   - `toolData`
   are not permitted.

### Rationale

This preserves:
- policy-safe core semantics
- strict top-level schema validation
- adapter evolution without top-level breaking changes

### Example

```json
{
  "extensions": {
    "jest": {
      "suite_names": ["auth.service.spec.ts"]
    }
  }
}
```

## Structured Failure Normalization Model

Normalized Result v0.1 supports canonical failure representation.

### Execution status mapping

| Execution outcome | Normalized execution.status |
|---|---|
| non-zero exit | failed |
| timeout | timed_out |
| missing binary | not_found |
| normalization/validation failure | normalization_failed |

### Rules

1. A normalized artifact must still exist on failure.
2. Failure results must validate against the same schema as successful results.
3. Missing metrics must be zeroed or omitted according to schema rules; they must never be fabricated.
4. Failure-specific detail lives in:
   - `error`
   - `evidence.issues`

### Goal

Downstream systems such as policy evaluation, inspect, and audit must be able to process failures without special missing-file logic.