# Policy Evaluation Result v0

Policy Evaluation Result v0 is the canonical BrikByteOS schema for the outcome
of evaluating **one policy rule** against normalized evidence.

It is the **post-policy, pre-verdict** contract boundary.

## Purpose

This schema exists to record:

- which policy was evaluated
- which version of that policy ran
- what evaluation status it produced
- why it produced that status
- which normalized evidence supports the result
- when the evaluation completed

## Required top-level fields

- `schema_family`
- `schema_version`
- `policy`
- `evaluation`

## Required policy fields

- `id`
- `version`

## Required evaluation fields

- `status`
- `reason`
- `evaluated_at`
- `evidence_refs`

## Fixed status vocabulary

- `passed`
- `failed`
- `not_applicable`
- `error`

### Meaning

- `passed`  
  The policy evaluated successfully and its condition was satisfied.

- `failed`  
  The policy evaluated successfully and its condition was violated.

- `not_applicable`  
  The policy was valid but did not apply to the provided normalized evidence.

- `error`  
  The policy could not be evaluated correctly due to an evaluation/runtime problem.

## Enforcement level semantics

`enforcement_level` is optional and refers to the severity or strictness of the
policy result itself.

Examples:
- `info`
- `warning`
- `error`
- `critical`

This field is **not** the final release verdict.

## Boundary rules

This schema does **not** contain:

- final release verdict
- release-wide policy aggregation
- AI-generated narratives
- action recommendations
- embedded normalized evidence payloads
- rule-engine internals or policy source code

## Evidence references

Evidence references point to normalized evidence only, such as:

- normalized issues
- normalized metrics
- severity rollups
- evidence completeness indicators
- whole normalized-result references when needed

## Composition rule

One Policy Evaluation Result represents one evaluated rule.

A later Verdict schema may aggregate many Policy Evaluation Results, but this
schema must not already become a verdict contract.

## Example

See:

```text
examples/valid.minimal.json
```