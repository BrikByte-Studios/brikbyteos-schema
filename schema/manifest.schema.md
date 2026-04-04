# Phase 0 Manifest Schema

This schema defines the canonical BrikByteOS Phase 0 persisted run summary artifact.

## Purpose

The manifest is the authoritative summary artifact for a run. It exists to provide:

- run identity
- lifecycle visibility
- selection summary
- bounded per-adapter summary
- artifact discovery

## Design constraints

- deterministic structure
- bounded vocabularies
- relative artifact paths only
- no additional properties
- no score/policy/verdict embedding in Phase 0

## Lifecycle posture

One schema supports:
- in-progress manifests
- completed manifests
- failed/partial manifests

## Important rule

The Go model in `brikbyteos-cli/internal/manifest` and this schema must remain semantically aligned.