# Verdict v0

Verdict v0 is the canonical BrikByteOS schema for the **final release-level
decision** after policy evaluation has completed.

It is the final decision contract boundary for:

- CLI verdict presentation
- audit linkage
- future control-plane/API surfaces

## Purpose

This schema exists to represent:

- the final verdict
- the environment/context that verdict applies to
- a compact summary of policy outcomes
- enough traceability for future audit composition

## Required top-level fields

- `schema_family`
- `schema_version`
- `verdict`
- `decision_context`
- `policy_summary`
- `decided_at`
- `traceability`

## Stable verdict vocabulary

- `approved`
- `warning`
- `rejected`

### Meaning

- `approved`  
  The release decision is positive and no blocking policy outcome prevents progression.

- `warning`  
  The release may proceed, but warning-level policy outcomes exist and must remain visible to downstream consumers.

- `rejected`  
  The release must not proceed because blocking policy outcomes exist.

## Decision context

Verdict v0 must identify the environment or context to which the decision applies.

Examples:
- `dev`
- `staging`
- `production`

Optional identifiers such as `target_id` may be included when they add stable
decision traceability.

## Policy summary boundary

Policy summary contains only release-level rollup information, such as:
- how many policies were evaluated
- how many failed
- how many warnings exist
- optionally how many passed

Policy summary must **not** contain:
- full policy result objects
- full normalized evidence objects
- raw execution details
- AI narratives

## Traceability

Verdict v0 must include enough references to support future audit composition.

These references may point to:
- policy result set
- normalized result set
- future audit bundle

Verdict v0 must not duplicate those lower-level payloads inline.

## Explicit v0 score decision

Score is **out of scope** for Verdict v0.

Reason:
- no stable score semantics are yet defined
- placeholder score fields create ambiguity
- policy-gated decisions should remain explicit and deterministic

## Boundary rules

Verdict v0 does **not** contain:
- raw evidence
- normalized evidence payloads
- embedded full policy results
- AI explanations
- analytics/reporting payloads unrelated to final decision

## Example

See:

```text
examples/valid.minimal.json
```