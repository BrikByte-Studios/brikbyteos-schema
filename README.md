# BrikByteOS Schema Repository

This repository is the canonical source-of-truth for published BrikByteOS schema
contracts.

Authoritative governance document:

```text
docs/governance/schema-versioning-strategy.md
```

### Core rules
- schema ownership lives here
- schema versioning is independent from product/repository release versions
- schema families use explicit version namespaces such as `v0`, `v1`
- runtime repositories consume canonical contracts rather than redefine them