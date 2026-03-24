# run-identity v0

This schema defines the canonical run identity contract for BrikByteOS.

## Notes

- `runId` is used for storage and traceability.
- Deterministic semantic equivalence must **not** depend on:
  - `runId`
  - `startedAt`
  - `finishedAt`

## Allowed environments

- `dev`
- `staging`
- `production`
- `unknown`

## Allowed execution modes

- `all`
- `single`
- `explicit_list`