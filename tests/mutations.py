# Created: 2026-10-22 11:25 -05:00 · Edited: 2026-10-22 11:38 -05:00
"""mutate — break the code on purpose and watch the named check go red (C1-T4).

Each mutation names a file, the text to replace, the replacement, and the check that
must fail while the mutation is in place. Before mutating, the check is run on the
clean tree and must pass: a check that is already red, or cannot run, proves nothing,
and is reported as BASELINE RED rather than as a kill. The file is always restored.

Outcomes, per the register's vocabulary: `killed` (baseline green, mutant red),
`survived` (baseline green, mutant still green), `baseline red`, `skipped` (the file
does not exist yet). Anything but killed makes this script exit 1.

    python3 tests/mutations.py            # run every mutation
    python3 tests/mutations.py M2 M3      # run some; an unknown id is an error

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
     "return Version", 'return ""',
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
KNOWN = {m[0] for m in MUTATIONS} | {"M3"}


def run(cmd: list[str]) -> int:
    env = dict(os.environ, PATH="/opt/homebrew/bin:" + os.environ.get("PATH", ""))
    try:
        return subprocess.run(cmd, cwd=ROOT, env=env, capture_output=True, text=True).returncode
    except FileNotFoundError:
        return 127


def mutate(mid: str, rel: str, find: str, repl: str, check: list[str], protects: str) -> bool:
    path = ROOT / rel
    if not path.is_file():
        print(f"{mid} skipped — {rel} does not exist yet ({protects})")
        return False
    original = path.read_text(encoding="utf-8")
    if find not in original:
        print(f"{mid} survived — the text to mutate is not in {rel}; the mutation is stale")
        return False
    baseline = run(check)
    if baseline != 0:
        print(f"{mid} baseline red — `{' '.join(check)}` exited {baseline} on the clean tree; nothing was proven")
        return False
    path.write_text(original.replace(find, repl, 1), encoding="utf-8")
    try:
        code = run(check)
    finally:
        path.write_text(original, encoding="utf-8")
    if code != 0:
        print(f"{mid} killed — `{' '.join(check)}` exited {code} with the mutation in place ({protects})")
        return True
    print(f"{mid} survived — `{' '.join(check)}` stayed green with the mutation in place")
    return False


def commit_msg_mutation() -> bool:
    """M3: a register ID in a commit subject is refused by the commit-msg hook, and a
    clean subject is accepted — both halves, or the hook that rejects everything passes."""
    hook = ROOT / ".githooks/commit-msg"
    def verdict(subject: str) -> int:
        with tempfile.NamedTemporaryFile("w", suffix=".txt", delete=False) as f:
            f.write(subject + "\n")
            name = f.name
        try:
            return run(["sh", str(hook), name])
        finally:
            os.unlink(name)
    clean = verdict("fix(C1-T4): a clean subject")
    if clean != 0:
        print(f"M3 baseline red — the commit-msg hook refused a clean subject (exit {clean})")
        return False
    dirty = verdict("fix(D1): a subject that cites a register row")
    if dirty != 0:
        print(f"M3 killed — commit-msg hook exited {dirty} on `fix(D1): …` and 0 on a clean subject")
        return True
    print("M3 survived — the commit-msg hook accepted a register ID in the subject")
    return False


def main(argv: list[str]) -> int:
    wanted = set(argv[1:])
    unknown = wanted - KNOWN
    if unknown:
        print(f"unknown mutation id(s): {', '.join(sorted(unknown))}; known: {', '.join(sorted(KNOWN))}")
        return 2
    results = []
    for mid, rel, find, repl, check, protects in MUTATIONS:
        if wanted and mid not in wanted:
            continue
        results.append(mutate(mid, rel, find, repl, check, protects))
    if not wanted or "M3" in wanted:
        results.append(commit_msg_mutation())
    if not results:
        print("nothing ran")
        return 2
    killed = sum(results)
    print(f"mutations: {killed} of {len(results)} killed" + ("" if killed == len(results) else " — NOT all killed"))
    return 0 if killed == len(results) else 1


if __name__ == "__main__":
    sys.exit(main(sys.argv))
