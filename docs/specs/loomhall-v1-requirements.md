# Loomhall v1 Requirements Spec

Created: 2026-10-22 10:59 -05:00 · Edited: 2026-10-22 11:10 -05:00

**Ratified 2026-10-22 by Martin Gleason.** Immutable from this date. Every later change is a `D<n>` in `docs/plans/00-register.md`, and the register wins where the two conflict. The pre-ratification text is kept in `docs/archive/`.

## Vision

**Loomhall is a digital space of community and collaboration for storytellers of all kinds.**

**Origins.** Built first for people who played World of Warcraft and used that game as a platform to tell all sorts of stories in Azeroth. Its principles became a community engine for telling, retelling, scheduling, and playing in the worlds we love.

**Core use case:** a member shares a story they wrote with their guild (and, from v1.5, with allied guilds).

## Principles

1. **Complexity lives in the design, not in the guild's lap.** Simple to set up, run and moderate, so people spend their energy on creativity and fun. Every option gets a sensible default.
2. **Core values: community, creativity, the common space.**
3. **Game-agnostic.** A small core plus one-click packs for each kind of play.
4. **Copyleft, never a cash spigot.** The tools this is built with learned from millions of people's work; Loomhall stays AGPL-3.0 and cannot be packaged and closed.
5. **Credit every creator.** All art and assets are properly sourced and credited, following Creative Commons and copyleft rules in letter and spirit.
6. **Bring your own tools.** Every outside service sits behind an adapter the project maintains, with a working default included.
7. **Trust-based in v1.** Privacy is enforced by the server's rules, not by encryption; zero-trust measures come later.
8. **The server is the only guard.** The front end is replaceable, so every protection lives in the back end.

## People and roles

One guild per install, sized for about 200 users. Roles are cumulative unless noted.

| Role | Can see | Can do |
| --- | --- | --- |
| Guest (not signed in) | Content marked guest/public, the public face of lore | Read; apply to join (applying requires a Discord sign-in) |
| Applicant | Same as guest, plus their own application thread | Submit an in-character application (e.g. pitch to a mercenary company) |
| Member | Guest + member content | Post as their characters, write stories and wiki pages, vote and comment on the community agreement, report content |
| Officer | Member + officer content | Approve applications, moderate, edit the agreement, veto proposals |
| Admin (an officer with admin privileges) | Officer content + the audit log | Reveal hidden character owners, change settings, manage packs, themes and adapters |
| Secret reader | A specific personal-secret block | Granted per block by its author; **officers and admins are not automatic readers** |

Game packs may rename these roles (for example "party" and "GM" for a D&D table) without changing what they can do.

## Milestones and features

Milestones are named, never numbered: **v1**, **v1.5**, **v2**. Features carry structural IDs (`F<N>`); tasks inside them are numbered in the plan (`F4-T1`). Requirements carry `FR<n>`; cross-cutting qualities carry `NFR<n>`. Both are cited by name in tests, plans and reviews and never renumbered.

| ID | Feature | Milestone |
| --- | --- | --- |
| F1 | Install, setup wizard, HTTPS and updates | v1 |
| F2 | Identity and membership (Discord login, applications, auto-join) | v1 |
| F3 | Characters (post as a character, hidden ownership, reveal, cross-space copies) | v1 |
| F4 | Visibility engine (four tiers, per-block, fail closed) | v1 |
| F5 | Editor (rich text over Markdown, mentions, visibility blocks, images) | v1 |
| F6 | Forum (IC and OOC areas, threads) | v1 |
| F7 | Wiki, lore and place pages | v1 |
| F8 | Story archive and tagging | v1 |
| F9 | Galleries with required credit | v1 |
| F10 | Spaces and packs (one-click packs, game vocabulary) | v1 |
| F11 | Calendar, RSVPs and Discord event sync | v1 |
| F12 | Notifications (bell, Discord DM, email) | v1 |
| F13 | Moderation and the community agreement | v1 |
| F14 | Audit log | v1 |
| F15 | Themes, accessibility and translation | v1 |
| F16 | Data adapters, backups and export | v1 |
| F17 | API, OpenAPI spec and generated clients | v1 (TypeScript client); v2 (Swift client) |
| F18 | Search (titles and tags) | v1 |
| F19 | Whiteboard pack (one editor at a time) | v1 |
| F20 | Federation with allied guilds (public content) + allied-reader tier | v1.5 |
| F21 | Live multi-person whiteboard | v1.5 |
| F22 | Import (new server, switching back ends) | v1.5 |
| F23 | NodeBB adapter pack | v1.5 |
| F24 | Optional AI-content screening, own spec | v1.5 |
| F25 | Full-text search | v1.5 |
| F26 | Wiki.js adapter (experimental) | v1.5 |

Features F20 to F26 have no requirements here. Each gets its own spec or decision rows when its milestone opens.

## Feature requirements

"Setting" means a per-guild option; the value after "default" ships out of the box.

### F1 Install, setup, HTTPS, updates

- **FR1** Install with one command on a rented VPS. The installer is a POSIX shell script with no prerequisite beyond a POSIX shell, and it checks everything it needs before acting. One-click marketplace images are a later goal.
- **FR2** Ships as one signed program. No Node, Docker or database server is required on the guild's server.
- **FR3** A first-run wizard asks for domain and first admin, confirms DNS, then turns on automatic HTTPS.
- **FR4** A built-in updater verifies signatures before applying an update and can roll back.

### F2 Identity and membership

- **FR5** Discord is the only sign-in. Guests read without signing in.
- **FR6** Join order is a setting, default **site first**: sign in with Discord, apply on the site, an officer approves, the site adds the person to the guild's Discord server and assigns the member role.
- **FR7** The application form and consent screen state clearly that approval adds the account to the Discord server.
- **FR8** Applications are in-character pitches; the form's prompts are configurable per guild.

### F3 Characters

- **FR9** One account, many characters. Players post, write and comment as a chosen character.
- **FR10** Out-of-character posts show the character by default; setting.
- **FR11** A player may hide which account owns a character. An admin can reveal it; every reveal is logged. The player is not notified by default; setting.
- **FR12** Character profiles are wiki pages. Every story or thread tagged with a character appears on its profile.
- **FR13** A character can be copied into another space and adapted. The copy is independent and shows "Adapted from \<original>"; the link follows the original's visibility, so viewers who can't see the original see "Adapted character" with no link.

### F4 Visibility engine

- **FR14** Four tiers applied **per block inside a page**: guest, member, officer, personal secret.
- **FR15** Personal secret: readable only by its author and the people the author allows. Officers and admins are not automatic readers.
- **FR16** The server removes every block a viewer may not see before sending anything. The same filter covers pages, search, previews, notifications, link previews, revision history, exports and the API.
- **FR17** Nested blocks take the stricter tier. A malformed or unknown block hides everything after it.
- **FR18** Links and mentions never reveal a hidden target.
- **FR19** Editing a page that contains blocks the editor can't see shows placeholders; the server keeps the hidden text intact on save.
- **FR20** Back ends that can't do per-block visibility use the fallback: the public text links to a separately restricted page or thread holding the hidden part.
- **FR21** New forum areas and wiki sections default to the **guest** tier. Officers can tighten any area at creation or later.

### F5 Editor

- **FR22** Rich-text editing that stores Markdown underneath, with a raw-Markdown mode.
- **FR23** Visibility blocks show as tinted, labeled boxes (never color alone).
- **FR24** Mentions of characters, places and pages; image upload with required alt text and stripped location data.
- **FR25** Full long-form writing works on phones.

### F6 Forum

- **FR26** Categories and threads; in-character and out-of-character are separated **by forum area**.
- **FR27** Category permissions follow the four tiers.

### F7 Wiki, lore and place pages

- **FR28** Wiki pages with full revision history.
- **FR29** Place pages for in-world locations, linked to the characters and stories set there.

### F8 Story archive and tagging

- **FR30** Stories with topic tags readers can browse; tags support synonyms managed by officers.
- **FR31** Stories carry the characters that appear in them.
- **FR32** Players can export their own work at any time.

### F9 Galleries

- **FR33** Personal and guild galleries.
- **FR34** Every upload requires creator, source and license; uploads without them are rejected.

### F10 Spaces and packs

- **FR35** An install holds multiple spaces (campaigns or worlds), separate by default.
- **FR36** Install the core, then add packs any time with one click. Packs stack (a WoW guild can add a D&D pack).
- **FR37** Everything optional is a pack: game vocabulary, maps, in-world timeline, character sheets, session planning, whiteboard, language packs.
- **FR38** Packs declare typed data schemas with migrations.
- **FR39** v1 ships these packs: whiteboard, maps, in-world timeline, character sheets, session planning, and WoW vocabulary.

### F11 Calendar and scheduling

- **FR40** Built-in event calendar with RSVPs.
- **FR41** Mirrors events to native Discord scheduled events and pulls Discord RSVPs back.
- **FR42** Raid-Helper adapter: reads events and sign-ups by default; accepts Premium webhooks when the guild has Premium. The site is the source of truth.

### F12 Notifications

- **FR43** Channels: on-site bell, Discord DM from the guild bot, optional email.
- **FR44** Each player picks their mix, default **bell + Discord DM**; optional per-character overrides.
- **FR45** A failed DM falls back to the bell.

### F13 Moderation and the community agreement

- **FR46** Every account accepts the community agreement at sign-up.
- **FR47** Officers approve applications; members' new posts can be held for approval (setting).
- **FR48** Reporting, post deletion (old versions kept), warnings and removal.
- **FR49** Admins can open a member's full post history to check patterns of behavior.
- **FR50** The agreement lives in its own wiki section showing created and last-updated dates.
- **FR51** Members propose changes and comment. Voting is a setting: anonymous thumbs up and down, both counts visible to everyone. Officers decide and can veto.
- **FR52** Votes are anonymous by design: the system records only that an account voted, never which way.
- **FR53** When the agreement changes, members are notified and must re-accept before posting.

### F14 Audit log

- **FR54** Logs character reveals, visibility changes, edits and deletes, membership changes, and theme and setting changes.
- **FR55** Readable by admins only. Retention is a setting, default 2 years.
- **FR56** Never records how anyone voted.

### F15 Themes, accessibility, translation

- **FR57** One theme styles the whole site. Built-in parts match exactly; adapter back ends get a close-match generated theme.
- **FR58** Ladder: pick a theme from a gallery; edit theme variables (colors, fonts, spacing); or replace the front end entirely.
- **FR59** Themes are shared through the `loomhall-themes` repo; authors choose their license and must credit sources.
- **FR60** Every theme passes automated WCAG 2.2 AA and colorblind checks before it's accepted.
- **FR61** Every user-facing string lives behind the translation layer. A build check fails on a hard-coded user-facing string.
- **FR62** v1 ships English. Right-to-left layout is proven in v1 with a pseudo-locale that flips direction and lengthens strings.
- **FR63** The translation layer accepts custom locale tags, so constructed and fictional languages (Klingon, Tolkien's Elvish, Thalassian, Krogan) can ship as language packs. A language pack carries its own license and credit. Community translations arrive as contributions under no milestone.

### F16 Data adapters, backups, export

- **FR64** Each slot has a default and bring-your-own options:

| Slot | Default | Bring your own |
| --- | --- | --- |
| Database | SQLite | Postgres, MySQL |
| File storage | Server disk | S3-compatible (R2, B2) |
| Backups | Server disk, scheduled | Google Drive, OneDrive, S3 |
| Email | Off | SMTP, Resend |

- **FR65** The project builds and maintains every adapter; guilds only point them at their own service.
- **FR66** Each guild registers its own Google, Microsoft and Discord apps through a guided wizard.
- **FR67** The application reaches storage only through a store interface. The default SQL implementation passes one shared test suite against both SQLite and Postgres, so switching is a configuration change.

### F17 API and clients

- **FR68** Everything the site does is available through a versioned API documented in OpenAPI.
- **FR69** The SvelteKit site is the reference front end; any other front end (React included) can be built on the same API.
- **FR70** A generated TypeScript client is published in v1. A generated Swift client for macOS and iOS is published in v2.

### F18 Search

- **FR71** Search by title and tag, filtered by what the searcher may see.

### F19 Whiteboard pack

- **FR72** Freeform board for brainstorming, running games and quick maps; one editor at a time in v1.
- **FR73** Visibility is set per board.

## Cross-cutting requirements

| ID | Area | Requirement |
| --- | --- | --- |
| **NFR1** | Cost | Runs for about 200 users on the cheapest VPS tier (1 GB RAM, about $5–7 a month); external services stay within free tiers by default |
| **NFR2** | Footprint | Core program idle under 100 MB resident, and under 300 MB during the NFR3 load, measured on the NFR1 server |
| **NFR3** | Speed | A forum thread or wiki page is served in under 200 ms at the 95th percentile with 50 simultaneous readers, on the NFR1 server |
| **NFR4** | Mobile | Every task works on a phone: reading, long-form writing, wiki editing, galleries, admin and moderation |
| **NFR5** | Screen-reader editing | A tester using TalkBack on Android and VoiceOver on iOS can, without sighted help, write a three-paragraph story containing one member-only block and one personal-secret block, and publish it. Scripted walkthrough, pass or fail per step |
| **NFR6** | Accessibility | WCAG 2.2 AA; colorblind-safe palettes; scalable fonts; screen-reader support; reduced-motion respected |
| **NFR7** | Translation | All interface text translatable, right-to-left languages included (mechanism in FR61 to FR63) |
| **NFR8** | Security | All visibility, ownership and permission checks enforced on the server; tokens encrypted at rest and kept out of logs; Discord bot gets minimum permissions, never Administrator |
| **NFR9** | Privacy | Trust-based in v1; no encryption of personal secrets; zero-trust deferred |
| **NFR10** | Licensing | Core and adapters AGPL-3.0; theme and language-pack authors choose their own license; no dependency under a non-open-source license (e.g. BSL, Commons Clause, watermark-required SDKs) |
| **NFR11** | Attribution | Every theme, image, language pack and asset carries creator, source and license; CI rejects entries without them |
| **NFR12** | Bring your own tools | Every outside service sits behind a project-maintained adapter with a working default (F16) |
| **NFR13** | Future-ready data | Every record carries a globally unique ID and an origin field naming the install that created it, so federation, live whiteboard and import in v1.5 need no renumbering |
| **NFR14** | No Node on the server | Node is used on build machines and CI only. The guild's server runs the single Go binary with the front end embedded |

## Stack

Baseline. Changing any of this is a `D<n>`.

- language: Go, current stable, pinned in `go.mod`; TypeScript for the front end; POSIX sh for the installer; Python 3 standard library for tooling only
- framework: Huma on `net/http`; SvelteKit with the static adapter, embedded in the binary; one SQL store through Bun over `database/sql`, with a pure-Go SQLite driver by default and Postgres and MySQL as bring-your-own
- runner: `go test ./...` for the back end; `vitest` for the front end; the store suite runs against SQLite and a Postgres service in CI
- layout: `cmd/loomhall` · `internal/` · `web/` (SvelteKit) · `packs/` · `adapters/` · `scripts/` · `docs/` — one repository; `loomhall-themes` holds themes only
- ci: GitHub Actions; `.githooks/pre-commit` locally

**Noted for tuning.** The data layer and driver are an area for later fine tuning. The C-backed SQLite driver, or another faster language for that layer, stays open. The register carries this as a standing risk fed by the trial's measurements.

## Hook intentions

Invariants that need a mechanism, not a sentence. The register's hooks table names the mechanism for each.

- Every `FR` is cited by at least one test; every test cites an `FR` or a `D`.
- No user-facing string escapes the translation layer (FR61).
- The visibility filter runs on every output path listed in FR16; a test enumerates the paths.
- The audit log schema has nowhere to store a vote's direction (FR52, FR56).
- An upload without creator, source and license is rejected (FR34).
- Every dependency's license is checked in CI (NFR10).
- The store suite runs against SQLite and Postgres on every pull request (FR67).
- The status page regenerates clean; no register ID appears in a commit subject; every file carries its Created and Edited stamp.

## Process

Loomhall is built under the Turing process from the `code-process-review` repository (cloned locally as `turing-review`).

- **Conventions:** a pinned copy of `conventions.md` is vendored into `docs/conventions.md` and imported by `CLAUDE.md`, together with `docs/conventions-local.md`. The vendored copy is never edited in place.
- **Specs and plans:** this spec is immutable once ratified. All changes, additions and deltas go in `docs/plans/00-register.md` as `D<n>` rows; the register wins where the two conflict.
- **Work chunks:** Features `F<N>`, Tasks `F<N>-T<M>`, Chores `C<N>`, Retrofits `F<N>b`. Gates are events crossed only with Marty's explicit yes, and the yes is written down when it is said.
- **First chores:** `C1` builds the harness (run, count, floor, mutate, pre-commit) and the repository scaffold. `C2` is the technical trial.
- **PRs:** conventional commits with the structural ID in scope, branches named `<feature-id>/<slug>`, rebase-and-merge.
- **Learning level: 7%.** The agent authors nearly everything; Marty reviews and logs every PR. Each feature plan carries about one small `owner: human` task, a variable or a small function that is easy to learn from, and every PR carries short notes on why the Go is written the way it is.
- **Workflow:** deep reasoning for spec review and planning, with clarifying questions asked one at a time; adversarial review at the end of every feature and at the start of every session; no feature is done without a passing check shown as evidence.
- **Technical trial (`C2`), first step after ratification:** two weeks proving the visibility engine with redacted editing (FR16, FR17, FR19), the editor on Android and iOS with a screen reader (NFR5), and memory and page speed on the NFR1 server (NFR2, NFR3). The VPS is chosen when the trial starts.

**Deep spec required** before build:

| Gate | Decisions that drive it |
| --- | --- |
| F1 Install and updates | Signing-key custody and location, key rotation after compromise, rollback against a database that has migrated forward |
| F2 Identity and membership | Discord scopes, consent wording, approve-then-join flow, token storage |
| F3 Characters | Ownership hiding, admin reveal and its logging, cross-space copies and origin-link visibility |
| F4 Visibility engine | Block syntax, tier nesting, redacted editing and merge-on-save, permission chaining, fallback for adapter back ends |
| F10 Spaces and packs | Pack format, schema migrations, how packs extend roles, vocabulary and languages |
| F13 Moderation and agreement | Anonymous-vote storage, re-accept flow, member-history access |
| F16 Data adapters | Store interface, SQLite and Postgres parity, backup format and restore |

**Claims about outside systems are proven before they are ratified.** Two are known now: that Discord's `guilds.join` scope lets the site add an approved applicant to the server (FR6), and that Raid-Helper's API serves event reads without Premium (FR42). Each probe is the first task of its feature's deep spec, and the decision row carries the probe's output.

## Out of scope for v1

Social-network features (walls, likes, friend feeds); multiple guilds per install; encryption of personal secrets; per-element whiteboard visibility; full-text search; import; federation; automated AI-content screening; Battle.net login; translations other than English; client libraries other than TypeScript.

## Decisions made at ratification

Formerly the open questions. Recorded here so the spec carries none.

- v1 packs: whiteboard, maps, in-world timeline, character sheets, session planning, WoW vocabulary (FR39).
- Data layer: SQLite first behind a store interface that makes Postgres a configuration change (FR64, FR67). Bring-your-own applies to every outside service (NFR12).
- Installer language: POSIX shell (FR1).
- Default tier for new forum areas and wiki sections: guest (FR21).
- Turing conventions: vendored at `C1` from the upstream clone.
- Repositories: `martin-gleason/loomhall` and `martin-gleason/loomhall-themes`, public from the first commit. Domain loomhall.app if purchased.

---

2026-10-22

\#AI/Claude
