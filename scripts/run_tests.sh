#!/bin/sh
# Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
# run — every check, non-zero on the first failure (C1-T4): Go vet, the front end's
# type check, then the whole suite once through check_floor.sh, which also enforces
# the floor. Prints the count by way of the floor line.
set -eu
cd "$(dirname "$0")/.."
[ -d /opt/homebrew/bin ] && PATH="/opt/homebrew/bin:$PATH"
echo "== go vet"
go vet ./...
echo "== svelte-check"
(cd web && npm run --silent check)
echo "== tests and floor"
scripts/check_floor.sh
echo "== run: all green"
