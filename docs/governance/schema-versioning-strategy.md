# BrikByteOS Schema Versioning and Ownership Strategy

## Purpose

This document defines the authoritative Phase 0 schema versioning and ownership
strategy for BrikByteOS.

The goal is to ensure contract evolution remains:

- stable
- predictable
- governable
- implementation-language agnostic
- compatible with a multi-repo architecture

This strategy applies to canonical BrikByteOS schema contracts and exists so
runtime and tooling repositories consume contract truth rather than invent it.

---

## 1. Canonical ownership decision

### Source of truth
`brikbyteos-schema` is the canonical source-of-truth repository for published
BrikByteOS schema contracts.

### Consumer rule
Other repositories, including `brikbyteos-cli`, may:
- consume schema contracts
- validate against schema contracts
- reference schema contracts
- include examples or fixtures derived from schema contracts

They must not:
- redefine canonical contract truth independently
- silently fork schema definitions
- invent parallel versioning rules outside this repository

---

## 2. Architecture alignment

This strategy is written for the actual BrikByteOS architecture:

- multi-repo
- Go-first runtime
- dedicated schema/contract concern
- runtime consumption of canonical contracts

This strategy must not assume monorepo-only paths or monorepo-only governance.

---

## 3. What unit is versioned

The primary versioned unit is the **schema family**.

Examples of schema families include:
- raw-result
- normalized-result
- policy-evaluation-result
- verdict
- audit-manifest

Each schema family owns its own explicit version namespace.

### Why this unit is chosen
This keeps versioning:
- small
- understandable
- contract-focused
- independent from repository release cycles

---

## 4. Schema family naming convention

Schema family names must be:

- lowercase
- hyphen-separated
- stable
- descriptive of the contract boundary

Examples:
- `raw-result`
- `normalized-result`
- `policy-evaluation-result`
- `verdict`
- `audit-manifest`

Avoid:
- language-specific names
- implementation-specific names
- ambiguous internal abbreviations unless formally defined

---

## 5. Schema version naming convention

Schema versions use explicit major namespaces such as:

- `v0`
- `v1`
- `v2`

### Phase 0 rule
Phase 0 introduces the namespace strategy, not a large multi-version catalog.

### Version format rule
Use short explicit major-style contract versions rather than product release
numbers.

Approved examples:
- `raw-result/v0`
- `normalized-result/v0`
- `verdict/v0`

---

## 6. Approved repository layout rule

Canonical schemas live under a stable repository-owned schema root.

Recommended layout:

```text
schemas/
  raw-result/
    v0/
      schema.json
      README.md
  normalized-result/
    v0/
      schema.json
      README.md
  policy-evaluation-result/
    v0/
      schema.json
      README.md
  verdict/
    v0/
      schema.json
      README.md
  audit-manifest/
    v0/
      schema.json
      README.md
```

### Layout principles
- family first
- version second
- explicit coexistence by directory
- no monorepo-only path assumptions
- no language/runtime-specific path ownership

---

## 7. Schema version vs product version

Schema versioning is distinct from:
- repository versioning
- product versioning
- CLI/runtime release tagging
- adapter release versioning

### Rule

A product release does not automatically imply a schema version change.

### Reason

Schema contracts evolve on their own compatibility boundaries and should not be
forced to match application release numbers unless explicitly justified.

---

## 8. Change classification policy

Every schema change must be classified as one of:

### A. Breaking change

A change is breaking if an existing compatible producer or consumer would need
behavioral or structural changes to continue working correctly.

Examples:
- removing a required field
- renaming a field without compatibility support
- changing field type incompatibly
- changing enum semantics incompatibly
- changing meaning in a way that invalidates prior consumers

### Breaking-change rule

Breaking changes require a new schema version namespace.

Example:
- `raw-result/v0` -> `raw-result/v1`

---

### B. Non-breaking change

A change is non-breaking if existing compatible producers and consumers can
continue to operate without required structural change.

Examples:
- adding an optional field
- clarifying constraints without invalidating current valid payloads
- tightening documentation around already-required semantics when payload shape
remains compatible

### Non-breaking-change rule

Non-breaking changes may remain within the same schema version namespace if
compatibility is preserved.

---

## C. Editorial-only change

A change is editorial-only if it changes documentation or presentation but does
not change the contract semantics.

Examples:
- wording improvements
- comments
- examples
- formatting
- typo fixes

### Editorial-change rule

Editorial-only changes do not change the schema version.

---

## 9. Backward compatibility expectations

### Producer rule

A producer must write artifacts that clearly identify the schema family and
schema version it is producing.

### Consumer rule

A consumer must explicitly decide which schema versions it supports.

Consumers may:
- support one version only
- support multiple versions deliberately

Consumers must not:
- assume all versions are interchangeable
- infer compatibility ad hoc

### Compatibility rule

Compatibility must be documented, not assumed.

---

## 10. Coexistence strategy for multiple versions

Multiple schema versions may coexist in future.

Examples:
- `raw-result/v0`
- `raw-result/v1`

### Coexistence rule

Coexistence is allowed only through explicit, versioned namespace separation.

That means:
- no ambiguous duplicate filenames without version context
- no hidden “latest” behavior as canonical truth
- no silent replacement of older schema contracts

### Reader/writer rule

Readers and writers must identify which version they consume or produce.

---

## 11. Relationship between canonical schemas and runtime repositories

### `brikbyteos-schema`

Owns canonical contract truth.

### `brikbyteos-cli`

Consumes approved canonical contracts.  
May validate artifacts against those contracts.  
Must not redefine ownership or versioning rules independently.

### `brikbyteos-adapters`

Must target canonical published schema families/versions when normalizing or
emitting contract-governed payloads.

brikbyteos-examples

May provide sample payloads and usage examples based on canonical schemas.
Must not become the source of truth.

---

## 12. Language-agnostic governance rule

Schema governance must remain independent from any one implementation language.

This strategy must work even if future consumers are implemented in:
- Go
- TypeScript
- Python
- Java
- other languages

Schema truth belongs to the contract repository, not to runtime code.

---

## 13. Contributor guidance

When adding a new schema, a contributor must be able to answer:
1. Which schema family is this?
2. Is this a new family or a new version of an existing family?
3. Is the change breaking, non-breaking, or editorial-only?
4. Does the schema belong in `brikbyteos-schema`?
5. Is the schema version distinct from the application release version?
6. Can future coexistence remain unambiguous?

If those answers are unclear, the schema change is not ready.

---

## 14. Phase 0 scope boundary

This strategy does not define:
- every individual schema in detail
- code generation
- migration tooling
- future version implementations beyond the namespace strategy
- product-version coupling rules beyond the explicit separation defined above