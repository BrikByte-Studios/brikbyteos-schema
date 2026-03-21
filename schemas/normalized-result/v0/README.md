# Normalized Result v0

Normalized Result v0 is the canonical BrikByteOS schema for **post-normalization,
pre-policy evidence**.

It exists to transform heterogeneous tool or adapter outputs into one stable,
policy-consumable structure.

## Purpose

This schema provides a common model for:

- normalized issues/findings
- normalized metrics
- severity rollups
- evidence completeness
- references to supporting raw evidence

## Required top-level fields

- `schema_family`
- `schema_version`
- `adapter_family`
- `normalized_at`
- `issues`
- `metrics`
- `severity_rollup`
- `evidence_completeness`

## Boundary rules

Normalized Result v0 does **not** contain:

- raw execution command details
- raw stdout/stderr payloads
- policy outcomes
- final verdicts
- AI summaries
- recommendations or action plans

If a field describes execution-only facts, it belongs in **Raw Result**.

If a field describes policy pass/fail or enforcement, it belongs in a
**Policy Evaluation Result**.

If a field describes final approval/rejection, it belongs in a **Verdict**.

## Issue model

Each normalized issue must express only canonical evidence semantics:

- stable id/reference
- category
- severity
- summary
- optional location
- optional evidence refs

## Metrics model

Metrics are represented canonically as named values with explicit kinds:

- integer
- number
- boolean

## Severity rollup

Severity rollup provides aggregate counts by severity bucket plus total.

## Evidence completeness

Evidence completeness is explicit and separate from:
- validation success
- severity counts
- issue presence

`policy_ready=true` means the normalized evidence is complete enough for
downstream policy evaluation.

## Example

See:

```text
examples/valid.minimal.json
```