#!/bin/sh
# Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
# floor — runs the suite once through count_tests.sh; fails if any test failed or if
# the count is below tests/TEST_FLOOR. Raise the floor when tests are added; never
# lower it without a decision row saying why.
set -eu
cd "$(dirname "$0")/.."
floor=$(tr -d '[:space:]' < tests/TEST_FLOOR)
if ! count=$(scripts/count_tests.sh); then
  echo "floor FAIL — a test failed (count so far: $count, floor $floor)"
  exit 1
fi
if [ "$count" -lt "$floor" ]; then
  echo "floor FAIL — $count tests passed, floor is $floor. A shrinking suite is blind, not clean."
  exit 1
fi
echo "floor pass — $count tests passed, floor is $floor"
