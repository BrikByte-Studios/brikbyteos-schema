# Raw Result v0

Raw Result v0 is the canonical BrikByteOS schema for **direct execution evidence**
captured immediately after adapter or tool invocation and before any downstream
normalization, interpretation, policy evaluation, or AI augmentation.

## Purpose

This schema exists to preserve:

- what was executed
- when it ran
- how it completed
- where raw output artifacts were captured

It does **not** represent:

- normalized findings
- policy outcomes
- verdict decisions
- AI summaries
- recommendations
- adapter-specific semantic interpretation

## Required top-level fields

- `schema_family`
- `schema_version`
- `adapter`
- `execution`

## Required execution fields

- `status`
- `started_at`
- `finished_at`
- `command`

## Status semantics

- `success`  
  The execution completed successfully at the raw process/invocation layer.

- `failure`  
  The execution completed but returned a failing raw outcome.

- `timeout`  
  The execution exceeded its allowed time and did not complete normally.

- `cancelled`  
  The execution was deliberately cancelled before normal completion.

- `unknown`  
  The raw execution outcome could not be reliably classified.

## Output reference rules

- prefer references to captured raw output artifacts
- do not inline large raw output by default
- keep references traceable and stable
- `stdout`, `stderr`, and optional `extras` are allowed

## Boundary rule

If a field describes:
- a finding
- severity interpretation
- policy pass/fail
- recommendation
- AI reasoning

it does **not** belong in Raw Result v0.

## Example

See:

```text
examples/valid.minimal.json
```