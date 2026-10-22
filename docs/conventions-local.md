# Conventions — local to Loomhall

Created: 2026-10-22 11:15 -05:00 · Edited: 2026-10-22 11:15 -05:00

Rules true of **this project only**. The vendored `docs/conventions.md` is never edited; anything specific to Loomhall belongs here.

---

## Upstream

The Turing process is cloned at `~/Local_Dev_Projects/turing-review`. Its scripts run from there and take this project's path:

    TR=~/Local_Dev_Projects/turing-review
    python3 $TR/scripts/gen_status.py .            # regenerate docs/plans/00-status.md
    python3 $TR/scripts/gen_status.py . --check    # H1
    python3 $TR/scripts/check_citations.py .       # every cited register ID has a row
    python3 $TR/scripts/check_register.py .        # rows well formed
    python3 $TR/scripts/stamp.py <file>            # Created/Edited stamp
    python3 $TR/scripts/check_harness.py --run .   # the five harness capabilities

A standalone clone without the upstream cannot regenerate the status page. That is `RR2` in the register, open.

## Languages

- Go is the application. Python 3, standard library only, is allowed for tooling, installers' helpers and theme scripts, never in the shipped binary. POSIX sh for the installer and hooks. TypeScript only inside `web/`.
- Every source file carries the Created/Edited stamp as a comment on line 1, after any shebang or `package` clause where the language requires it.

## Learning tasks

Each feature plan names about one `owner: human` task. It is small by design, a constant, a variable or a function under ten lines, and the agent writes its test before handing it over, so Marty sees red turn green.

## Pull requests

Every PR body has four short sections: **What changed · Why the Go looks like this · What to test · What could break.** Written for someone who reads code and does not write Go daily.

## Words not used

`Phase`. Milestones are `v1`, `v1.5`, `v2`; lifecycle is a status field; gates are events.

## Harness

Not yet declared. `C1` builds the five capabilities and declares them here as `run`, `count`, `floor`, `mutate`, `precommit`. Until then `check_harness.py` reports this project as undeclared, by name.
