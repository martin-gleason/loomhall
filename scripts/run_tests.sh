#!/bin/sh
# Created: 2026-10-22 12:20 -05:00 · Edited: 2026-10-22 12:20 -05:00
# run — every check, non-zero on the first failure (C1-T4). Prints no count on
# purpose: an exit code cannot show a suite that stopped running. count_tests.sh does.
set -eu
cd "$(dirname "$0")/.."
# Homebrew's bin is not on PATH inside some git clients; add it if it exists.
[ -d /opt/homebrew/bin ] && PATH="/opt/homebrew/bin:$PATH"
echo "== go vet"
go vet ./...
echo "== go test"
go test -count=1 ./...
echo "== vitest"
(cd web && npx vitest run)
echo "== run: all green"
