#!/usr/bin/env python3
"""
validate-examples.py

Purpose:
    Validate canonical v0 example payloads against their colocated JSON Schemas
    and perform a small number of coherence checks across the fixture chain.

Why this exists:
    - keeps canonical example validation reproducible
    - prevents docs/examples drift from schema contracts
    - gives maintainers one command to prove the fixture set is healthy

Dependencies:
    - Python 3
    - jsonschema package

Install dependency:
    python3 -m pip install jsonschema
"""

from __future__ import annotations

import json
import sys
from pathlib import Path

try:
    from jsonschema import Draft202012Validator
except ImportError as exc:  # pragma: no cover
    print(
        "error: missing dependency 'jsonschema'. "
        "Install it with: python3 -m pip install jsonschema",
        file=sys.stderr,
    )
    raise SystemExit(2) from exc


REPO_ROOT = Path(__file__).resolve().parent.parent

SCHEMA_FIXTURES = [
    (
        "raw-result",
        REPO_ROOT / "schemas/raw-result/v0/schema.json",
        REPO_ROOT / "schemas/raw-result/v0/examples/valid.minimal.json",
    ),
    (
        "normalized-result",
        REPO_ROOT / "schemas/normalized-result/v0/schema.json",
        REPO_ROOT / "schemas/normalized-result/v0/examples/valid.minimal.json",
    ),
    (
        "policy-evaluation-result",
        REPO_ROOT / "schemas/policy-evaluation-result/v0/schema.json",
        REPO_ROOT / "schemas/policy-evaluation-result/v0/examples/valid.minimal.json",
    ),
    (
        "verdict",
        REPO_ROOT / "schemas/verdict/v0/schema.json",
        REPO_ROOT / "schemas/verdict/v0/examples/valid.minimal.json",
    ),
    (
        "audit-manifest",
        REPO_ROOT / "schemas/audit-manifest/v0/schema.json",
        REPO_ROOT / "schemas/audit-manifest/v0/examples/valid.minimal.json",
    ),
]


def load_json(path: Path) -> dict:
    """Load a JSON file into a Python dictionary."""
    with path.open("r", encoding="utf-8") as handle:
        return json.load(handle)

def artifact_family(artifact: dict) -> str | None:
    """Extract the canonical artifact family from an audit-manifest entry.

    The script prefers the explicit repository contract field if present, but
    also tolerates a few alternatives so the coherence check remains resilient
    during Phase 0 evolution.
    """
    for key in ("schema_family", "artifact_type", "type", "kind"):
        value = artifact.get(key)
        if isinstance(value, str) and value.strip():
            return value.strip()

    return None

def extract_run_id(payload: dict) -> str | None:
    """Extract run_id from either nested execution data or a legacy top-level field."""
    execution = payload.get("execution")
    if isinstance(execution, dict):
        run_id = execution.get("run_id")
        if isinstance(run_id, str) and run_id.strip():
            return run_id.strip()

    run_id = payload.get("runId")
    if isinstance(run_id, str) and run_id.strip():
        return run_id.strip()

    return None


def validate_fixture(schema_path: Path, fixture_path: Path) -> None:
    """Validate one fixture against one schema."""
    schema = load_json(schema_path)
    fixture = load_json(fixture_path)

    validator = Draft202012Validator(schema)
    errors = sorted(validator.iter_errors(fixture), key=lambda e: list(e.path))

    if errors:
        formatted = "\n".join(
            f"- {'/'.join(map(str, err.path)) or '<root>'}: {err.message}"
            for err in errors
        )
        raise ValueError(
            f"fixture validation failed\n"
            f"schema:  {schema_path}\n"
            f"fixture: {fixture_path}\n"
            f"errors:\n{formatted}"
        )


def validate_cross_fixture_coherence() -> None:
    """Validate lightweight semantic coherence across the canonical fixture set.

    These checks are intentionally small and should verify contract consistency
    without imposing assumptions that belong to runtime storage layout rather
    than schema ownership.

    Rules enforced:
      - all fixtures share the same run identifier
      - audit manifest contains the expected artifact types
      - audit manifest artifact paths are non-empty relative-style paths
    """
    raw = load_json(REPO_ROOT / "schemas/raw-result/v0/examples/valid.minimal.json")
    normalized = load_json(REPO_ROOT / "schemas/normalized-result/v0/examples/valid.minimal.json")
    policy = load_json(REPO_ROOT / "schemas/policy-evaluation-result/v0/examples/valid.minimal.json")
    verdict = load_json(REPO_ROOT / "schemas/verdict/v0/examples/valid.minimal.json")
    audit = load_json(REPO_ROOT / "schemas/audit-manifest/v0/examples/valid.minimal.json")

    expected_run_id = extract_run_id(raw)

    for name, payload in [
        ("normalized-result", normalized),
        ("policy-evaluation-result", policy),
        ("verdict", verdict),
        ("audit-manifest", audit),
    ]:
        run_id = extract_run_id(payload)
        if run_id != expected_run_id:
            raise ValueError(
                f"fixture coherence failed: {name} run_id {run_id!r} "
                f"does not match raw-result run_id {expected_run_id!r}"
            )

    artifacts = audit.get("artifacts", [])
    artifact_types = {artifact_family(artifact) for artifact in artifacts}
    expected_types = {
        "raw-result",
        "normalized-result",
        "policy-evaluation-result",
        "verdict",
    }

    if None in artifact_types:
        raise ValueError(
            "fixture coherence failed: one or more audit manifest artifacts are missing "
            "a recognizable family field (expected one of: schema_family, artifact_type, type, kind)"
        )

    if artifact_types != expected_types:
        raise ValueError(
            f"fixture coherence failed: audit manifest artifact types {sorted(artifact_types)!r} "
            f"do not match expected types {sorted(expected_types)!r}"
        )

    for artifact in artifacts:
        path = artifact.get("path")
        if not isinstance(path, str) or not path.strip():
            raise ValueError(
                f"fixture coherence failed: audit artifact path is missing or empty "
                f"for artifact {artifact!r}"
            )

        normalized_path = path.strip()

        if normalized_path.startswith("/"):
            raise ValueError(
                f"fixture coherence failed: audit artifact path {normalized_path!r} "
                f"must not be absolute"
            )

        if normalized_path.startswith("../") or normalized_path == "..":
            raise ValueError(
                f"fixture coherence failed: audit artifact path {normalized_path!r} "
                f"must not escape its artifact root"
            )

def main() -> int:
    """Run full fixture validation and coherence checks."""
    for name, schema_path, fixture_path in SCHEMA_FIXTURES:
        validate_fixture(schema_path, fixture_path)
        print(f"ok  {name}: {fixture_path.relative_to(REPO_ROOT)}")

    validate_cross_fixture_coherence()
    print("ok  cross-fixture coherence")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())