# Audit Manifest v0 Notes

## Contract identity

- schema family: `audit-manifest`
- schema version: `v0`

## Boundary

Audit Manifest v0 is the canonical contract for:
- bundle identity
- artifact inventory
- schema-family/version inventory
- references to important artifacts

It is not the payload owner for:
- raw execution evidence
- normalized evidence
- policy evaluation content
- final verdict content

Those remain in their own schema-governed artifacts.

## Key design rules

### Typed artifact entries
Artifacts must be inventory items with explicit kinds and paths.
A plain untyped path list is not sufficient.

### Relative path semantics
Paths are bundle-relative in v0 to avoid host-specific ambiguity.

### Explicit schema inventory
Bundle consumers should be able to see which schema families/versions are present
without scanning every payload first.

### Integrity-ready without overbuilding
A structured integrity extension area exists, but v0 does not define full
hash/signature algorithms yet.

## Change discipline

Breaking changes:
- changing path semantics incompatibly
- removing required artifact typing
- embedding full referenced payloads into the manifest
- changing key reference meaning incompatibly

Non-breaking changes:
- adding optional artifact metadata
- clarifying descriptions
- adding optional future integrity subfields while preserving v0 compatibility

Editorial-only changes:
- examples
- comments
- wording
- formatting