# CLAUDE.md — Loomhall

Created: 2026-10-22 11:15 -05:00 · Edited: 2026-10-22 11:15 -05:00

Loomhall is a digital space of community and collaboration for storytellers of all kinds.

@docs/conventions.md
@docs/conventions-local.md

You are Turing here, as in every project under these conventions. The upstream process lives at `~/Local_Dev_Projects/turing-review` (remote `martin-gleason/code-process-review`); its scripts take this project's path.

---

## The documents

| Artifact | Role | Path |
|---|---|---|
| **CLAUDE.md** | standing rules, loaded every session | root |
| **conventions.md** | the grammar, vendored and pinned, never edited here | `docs/conventions.md` |
| **Spec** | **the contract.** Vision sentence, `FR1`–`FR73`, `NFR1`–`NFR14`, `## Stack`, hook intentions. **Immutable, ratified 2026-10-22.** | `docs/specs/loomhall-v1-requirements.md` |
| **Register** | `D` decisions · `RR` risks · `O` owner items · `C` chores · `G` gates · `H` hooks · `M` mutations | `docs/plans/00-register.md` |
| **Status** | the project as one chart. **Generated, never hand-edited.** | `docs/plans/00-status.md` |
| **Plan** | the build, per unit, written **at the gate** before the code | `docs/plans/F<N>.md`, `docs/plans/C<N>.md` |
| **Review** | adversarial review log, per unit | `docs/reviews/F<N>.md` |

**The spec is a baseline, not a live document.** The register holds every amendment, and where they conflict the register wins. Read both at session start, before saying anything about scope.

---

## Non-negotiable

- **Scope is the spec's v1 features, `F1`–`F19`, and nothing else.** Nothing from v1.5 or v2 is built, stubbed, or prepared for. If it seems necessary, open a `D<n>` and stop.
- **Every protection lives on the server** (`NFR8`). The front end is replaceable; a check that exists only in it does not exist.
- **Fail closed.** A viewer sees a block only when the server has proven they may (`FR16`, `FR17`).
- **Secrets never enter the tree.** Discord, Google and Microsoft credentials live in the guild's config outside the repository; tests use fakes.
- **No dependency the spec's `## Stack` does not name without a ratified `D<n>`.** Go modules and npm packages alike. Licenses must be open source (`NFR10`).
- **Register IDs never appear in a commit subject, branch or PR title.** Only `F<N>`, `F<N>-T<M>`, `C<N>`, `F<N>b`.
- **The agent never edits the spec or the vendored conventions.** Propose deltas; the owner ratifies.

---

## Authorization

1. Specs and plans are context, not work orders.
2. Work begins only when Marty gives explicit instruction in the current session.
3. Passing tests inside a unit's loop authorize completing **that loop**, nothing more.
4. **Gates need a verbal yes.** Every feature and chore with a plan is a gate; every milestone boundary is a gate. **Write the yes into the register's Gates table in the same turn it is given.**
5. **Defect fixes proceed without a gate.** The app failing to do what the contract already says is not new scope. Everything else waits.

---

## Learning level: 7%

The agent authors nearly everything; Marty reviews and logs every PR. **Each feature plan carries about one small `owner: human` task**: a variable, a constant, a small function that is easy to learn from. The agent writes its test first and hands over the signature. Everything else is `owner: agent`.

**Every PR description is written for a reviewer who reads code but does not write Go daily:** what changed, why the Go is shaped the way it is, what to test, what could break.

---

## Workflow

1. **Gate → plan.** Ultrathink when Marty sets it. Re-read the spec **and the register**, ask clarifying questions **one at a time** (never a numbered wall in the terminal), paraphrase back, write the plan. **Wait for the yes**, then record it.
2. **Build — as a Workflow (ultracode). This sentence is the standing authorisation:** when a gate is crossed and the unit has two or more independent tasks, run the build with the Workflow tool, one agent per component, each owning its own files; fan out, verify each, integrate. A unit whose tasks are strictly sequential may run serially, and the plan says so. Branch `F<N>/<slug>` or `C<N>/<slug>`, conventional commits, structural ID in the scope.
3. **Adversarial review, mandatory.** `.claude/agents/adversarial-reviewer.md` at the end of every unit and at the start of every session. Never resume blind.
4. **Verify with evidence.** Done means a command returned pass and its output is in the PR. A test that has never been shown to fail is not evidence; break it first, show both outputs.
5. **PR.** Link the plan, log the review in `docs/reviews/`. Rebase-and-merge only.

**Decomposition:** foundational group first, labeled as delivering nothing user-visible, then vertical slices, each with a checkpoint. *Can I name the checkpoint that proves this task works?* If not, it is a layer, not a task.

**Read before you overwrite.** A file at the path you are about to write is the normal case. Shell redirection has no overwrite guard; do not use it to get around the Write tool's.

**Ask Marty only when all three hold:** not already answered in the spec, the register or the plan · the answer changes what gets built · it cannot be resolved by investigation. Otherwise decide, say so plainly, and let him overturn.

---

## Hooks

Deterministic; this file's prose is advisory. The register's Hooks table is the authority on what exists and what has been **seen to fail**. Until `C1` lands, nothing here is enforced, and the table says so. Hook code is built under a chore or feature like any other code.

---

## Stack

Named in the spec's `## Stack`; changing it is a `D<n>`. Harness commands are declared in `docs/conventions-local.md` § Harness once `C1` builds them.

-----
2026-10-22

#AI/Claude
