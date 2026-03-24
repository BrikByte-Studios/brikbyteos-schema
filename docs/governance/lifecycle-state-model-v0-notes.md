# lifecycle-state-model v0 notes

## Purpose

This document defines the canonical execution lifecycle vocabulary for BrikByteOS Phase 1.

## Separation of concerns

The following concepts must remain distinct:

1. **Adapter execution lifecycle state**
   - `pending`
   - `running`
   - `completed`
   - `failed`
   - `timed_out`
   - `skipped`
   - `unavailable`

2. **Run aggregate lifecycle state**
   - `pending`
   - `running`
   - `completed`
   - `failed`
   - `timed_out`

3. **Domain result summary state**
   - `pass`
   - `fail`
   - `warn`
   - `unknown`

## Important rule

`completed` does **not** mean domain pass.

Example:
- adapter execution lifecycle state = `completed`
- domain result summary state = `fail`

This means the tool executed successfully and produced a valid failing result.

## Phase 1 decision

Use `unavailable` as the canonical lifecycle state.

Do **not** use `binary_missing` as a lifecycle enum.
If needed, represent that as an issue/error code such as:
- `TOOL_NOT_FOUND`
- `BINARY_MISSING`

## Aggregate run derivation

Run is `completed` when:
- orchestration finished
- manifest was persisted
- adapter results were persisted

even if some adapters are:
- `failed`
- `timed_out`
- `unavailable`

Run is `failed` when orchestration itself cannot complete.

Run is `timed_out` only for top-level run timeout.