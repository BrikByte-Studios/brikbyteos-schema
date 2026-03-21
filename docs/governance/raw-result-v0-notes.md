# Raw Result v0 Notes

## Contract identity

- schema family: `raw-result`
- schema version: `v0`

## Core boundary

Raw Result v0 captures **execution evidence only**.

It is the contract boundary between:
- adapter/tool invocation
- downstream Normalized Result generation

## Change discipline

Examples of breaking changes:
- removing `execution.status`
- changing `started_at` from RFC3339 string to another type
- renaming `adapter.name`

Examples of non-breaking changes:
- adding a new optional metadata field
- adding a new optional output reference field if compatibility is preserved

Examples of editorial-only changes:
- clarifying descriptions
- improving examples
- formatting changes

## Consumption rule

Runtime consumers may validate payloads against this contract, but they must not
reinterpret Raw Result as a normalized or policy-layer schema.