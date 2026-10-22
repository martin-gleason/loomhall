#!/bin/sh
# Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
# count — prints ONE integer: Go tests and subtests that ran and passed, plus vitest
# tests that passed. Skipped and failed tests are not counted. Exits 1 if any test
# or package failed (a package that does not build is a failure, not a zero), so a caller that only looks at the number cannot mistake red for small.
# Separate from run on purpose: a suite that silently stopped running still exits 0,
# and only the count shows it shrank. check_floor.sh compares it with tests/TEST_FLOOR.
set -eu
cd "$(dirname "$0")/.."
[ -d /opt/homebrew/bin ] && PATH="/opt/homebrew/bin:$PATH"
gostat=$(mktemp)
go_pass=$(go test -count=1 -json ./... 2>/dev/null | python3 -I -c '
import sys, json
passed = failed = 0
for line in sys.stdin:
    try:
        e = json.loads(line)
    except ValueError:
        continue
    if e.get("Action") == "fail":
        failed += 1          # a test, or a whole package that failed to build
    elif e.get("Action") == "pass" and e.get("Test"):
        passed += 1
open(sys.argv[1], "w").write(str(failed))
print(passed)' "$gostat")
go_fail=$(cat "$gostat"); rm -f "$gostat"
out=$(mktemp)
(cd web && npx vitest run --reporter=json --outputFile="$out" >/dev/null 2>&1) || true
read -r vt_pass vt_fail <<EOT
$(python3 -I -c '
import sys, json
try:
    d = json.load(open(sys.argv[1]))
    print(d["numPassedTests"], d["numFailedTests"])
except Exception:
    print(0, 1)' "$out")
EOT
rm -f "$out"
echo $((go_pass + vt_pass))
if [ "$go_fail" -ne 0 ] || [ "$vt_fail" -ne 0 ]; then
  echo "count: $go_fail Go and $vt_fail vitest test(s) FAILED" >&2
  exit 1
fi
