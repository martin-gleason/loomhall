# Created: 2026-10-22 12:30 -05:00 · Edited: 2026-10-22 12:30 -05:00
"""mutate — break the code on purpose and watch the named check go red (C1-T4).

Each mutation names a file, the text to replace, the replacement, and the check that
must fail while the mutation is in place. The file is always restored, even if the
check crashes. A mutation whose check stays green is reported as ESCAPED and makes
this script exit 1: that is a finding, not a pass.

    python3 tests/mutations.py            # run every mutation
    python3 tests/mutations.py M2 M3      # run some

Standard library only. The register's Mutations table (M<n>) mirrors this list.
"""
from __future__ import annotations

import os
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
TR = Path(os.environ.get("TURING_REVIEW", Path.home() / "Local_Dev_Projects/turing-review"))

# (id, file, find, replace, check command, what it protects)
MUTATIONS = [
    ("M1", "internal/version/version.go",
     'return Version', 'return ""',
     ["go", "test", "-count=1", "./internal/version/"],
     "the version package test catches a String() that returns nothing"),
    ("M2", "internal/id/id_test.go",
     "func TestNew_NeverRepeats(t *testing.T) {\n",
     "func TestNew_NeverRepeats(t *testing.T) {\n\tt.Skip(\"mutation M2\")\n",
     ["scripts/check_floor.sh"],
     "the floor catches a test that stops running while the suite stays green"),
    ("M4", "docs/plans/00-status.md",
     "# Status — loomhall", "# Status — loomhall (edited by hand)",
     ["python3", str(TR / "scripts/gen_status.py"), ".", "--check"],
     "the status check catches a hand edit of the generated page"),
]


def run(cmd: list[str]) -> int:
    env = dict(os.environ, PATH="/opt/homebrew/bin:" + os.environ.get("PATH", ""))
    return subprocess.run(cmd, cwd=ROOT, env=env, capture_output=True, text=True).returncode


def mutate(mid: str, rel: str, find: str, repl: str, check: list[str], protects: str) -> bool:
    path = ROOT / rel
    if not path.is_file():
        print(f"{mid} SKIPPED — {rel} does not exist yet ({protects})")
        return True
    original = path.read_text(encoding="utf-8")
    if find not in original:
        print(f"{mid} ESCAPED — the text to mutate is not in {rel}; the mutation is stale")
        return False
    path.write_text(original.replace(find, repl, 1), encoding="utf-8")
    try:
        code = run(check)
    finally:
        path.write_text(original, encoding="utf-8")
    if code != 0:
        print(f"{mid} run — caught: `{' '.join(check)}` exited {code} ({protects})")
        return True
    print(f"{mid} ESCAPED — `{' '.join(check)}` stayed green with the mutation in place")
    return False


def commit_msg_mutation() -> bool:
    """M3: a register ID in a commit subject is refused by the commit-msg hook."""
    hook = ROOT / ".githooks/commit-msg"
    with tempfile.NamedTemporaryFile("w", suffix=".txt", delete=False) as f:
        f.write("fix(D1): a subject that cites a register row\n")
        name = f.name
    try:
        code = run(["sh", str(hook), name])
    finally:
        os.unlink(name)
    if code != 0:
        print(f"M3 run — caught: commit-msg hook exited {code} on `fix(D1): …`")
        return True
    print("M3 ESCAPED — the commit-msg hook accepted a register ID in the subject")
    return False


def main(argv: list[str]) -> int:
    wanted = set(argv[1:])
    ok = True
    for mid, rel, find, repl, check, protects in MUTATIONS:
        if wanted and mid not in wanted:
            continue
        ok &= mutate(mid, rel, find, repl, check, protects)
    if not wanted or "M3" in wanted:
        ok &= commit_msg_mutation()
    print("mutations:", "all caught" if ok else "AT LEAST ONE ESCAPED")
    return 0 if ok else 1


if __name__ == "__main__":
    sys.exit(main(sys.argv))
