#!/bin/sh
set -eu

###############################################################################
# validate-examples.sh
#
# Purpose:
#   Run canonical v0 example validation in one stable entrypoint.
#
# Why this exists:
#   - Makefile should call one simple script
#   - local and CI usage should be identical
###############################################################################

REPO_ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

cd "$REPO_ROOT"
python3 ./scripts/validate-examples.py