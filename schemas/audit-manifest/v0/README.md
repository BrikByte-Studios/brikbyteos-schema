# Audit Manifest v0

Audit Manifest v0 is the canonical BrikByteOS schema for the **top-level
inventory and linkage contract** of one audit bundle.

It is the root contract that tells a consumer:

- which bundle this is
- which artifacts are included
- which schema families/versions are represented
- which artifacts are the important ones
- where future integrity metadata can attach

## Purpose

This schema exists to provide a stable structure for:

- bundle identity
- typed artifact inventory
- schema usage inventory
- references to key artifacts
- future integrity extension

## Required top-level fields

- `schema_family`
- `schema_version`
- `bundle`
- `created_at`
- `artifacts`
- `schema_inventory`
- `key_refs`

## Bundle identity

One manifest represents one audit bundle or run.

`bundle.id` is the stable manifest-level identifier.

## Path rules

Artifact paths in v0 are:

- bundle-relative
- deterministic
- portable within the bundle

Artifact paths must not be:
- absolute filesystem paths
- host-specific paths
- ambiguous plain filenames without manifest context

## Artifact inventory

Each artifact entry is typed and must include:

- `id`
- `kind`
- `path`

Optional fields:
- `media_type`
- `role`
- `schema_ref`

The manifest inventories artifacts but does **not** embed their payloads.

## Schema inventory

The manifest must explicitly list schema families/versions used in the bundle.

This allows bundle inspection without re-parsing every artifact payload first.

## Key references

`key_refs` points to important artifacts, especially the canonical verdict.

This helps consumers find the core decision artifact quickly while still keeping
the full inventory explicit.

## Integrity extension

v0 reserves a structured `integrity` section for future use.

This section exists so later hash/signature metadata can be added without
redesigning the manifest.

v0 does **not** define full hashing or signing semantics.

## Boundary rules

Audit Manifest v0 does **not** contain:

- embedded raw result payloads
- embedded normalized result payloads
- embedded policy result payloads
- embedded verdict payloads
- generic archive metadata unrelated to bundle inventory
- full integrity algorithm specifications

## Example

See:

```text
examples/valid.minimal.json
```