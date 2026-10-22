#!/bin/sh
# Created: 2026-10-22 12:20 -05:00 · Edited: 2026-10-22 12:20 -05:00
# count — prints ONE integer: Go tests and subtests that ran and passed, plus vitest
# tests that passed. A skipped or failed test is not counted. Compare it with
# tests/TEST_FLOOR (check_floor.sh does). Separate from run on purpose: a suite that
# silently stopped running still exits 0, and only the count shows it shrank.
set -eu
cd "$(dirname "$0")/.."
[ -d /opt/homebrew/bin ] && PATH="/opt/homebrew/bin:$PATH"
go_pass=$(go test -count=1 -json ./... 2>/dev/null | python3 -I -c '
import sys, json
n = 0
for line in sys.stdin:
    try:
        e = json.loads(line)
    except ValueError:
        continue
    if e.get("Action") == "pass" and e.get("Test"):
        n += 1
print(n)')
out=$(mktemp)
(cd web && npx vitest run --reporter=json --outputFile="$out" >/dev/null 2>&1) || true
vt_pass=$(python3 -I -c '
import sys, json
try:
    print(json.load(open(sys.argv[1]))["numPassedTests"])
except Exception:
    print(0)' "$out")
rm -f "$out"
echo $((go_pass + vt_pass))
