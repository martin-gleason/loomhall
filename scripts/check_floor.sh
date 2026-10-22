#!/bin/sh
# Created: 2026-10-22 12:20 -05:00 · Edited: 2026-10-22 12:20 -05:00
# floor — the count may not fall below tests/TEST_FLOOR. Raise the floor when tests
# are added; never lower it without a decision row saying why.
set -eu
cd "$(dirname "$0")/.."
floor=$(tr -d '[:space:]' < tests/TEST_FLOOR)
count=$(scripts/count_tests.sh)
if [ "$count" -lt "$floor" ]; then
  echo "floor FAIL — $count tests passed, floor is $floor. A shrinking suite is blind, not clean."
  exit 1
fi
echo "floor pass — $count tests passed, floor is $floor"
