# Loomhall v1 Requirements Spec

Created: 2026-10-22 11:07 -05:00 · Edited: 2026-10-22 11:07 -05:00

Oct 22, 2026 · @Martin Gleason

This spec is the contract for Loomhall v1. It becomes immutable once ratified and copied to `docs/specs`; later changes go in the plan.

## Vision and principles

**Vision (Marty's words):** Loomhall is a digital space of community and collaboration for storytellers of all kinds. Built first for people who played WoW and used that game as a platform to tell all sorts of stories in Azeroth, its principles became a community engine for telling, retelling, scheduling, and playing in the worlds we love.

**Core use case:** a member shares a story they wrote with their guild (and, from v1.5, with allied guilds).

**Principles**

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

## Feature list

Features carry IDs; tasks inside them are numbered in the plan (`F4-T1`). Version tags say when a feature ships.

| ID | Feature | Version |
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
| F17 | API, OpenAPI spec and generated clients | v1 |
| F18 | Search (titles and tags) | v1 |
| F19 | Whiteboard pack (one editor at a time) | v1 |
| F20 | Federation with allied guilds (public content) + allied-reader tier | v1.5 |
| F21 | Live multi-person whiteboard | v1.5 |
| F22 | Import (new server, switching back ends) | v1.5 |
| F23 | NodeBB adapter pack | after v1 |
| F24 | Optional AI-content screening (Claude Haiku or Sonnet) | later, own spec |
| F25 | Full-text search | later |
| F26 | Wiki.js adapter (experimental) | later |

## Feature requirements

"Setting" means a per-guild option; the value after "default" ships out of the box.

### F1 Install, setup, HTTPS, updates

- Install with one command on a rented VPS; one-click marketplace images are a later goal. The installer checks its own prerequisites (including Python 3 if any Python tooling is used) before acting.
- Ships as one signed program; no Node, Docker or database server required on the guild's server.
- First-run wizard asks for domain and first admin, confirms DNS, then turns on automatic HTTPS.
- Built-in updater with signature checks and rollback.

### F2 Identity and membership

- Discord is the only sign-in. Guests read without signing in.
- Join order is a setting, default **site first**: sign in with Discord, apply on the site, an officer approves, the site adds the person to the guild's Discord server and assigns the member role.
- The application form and consent screen state clearly that approval adds the account to the Discord server.
- Applications are in-character pitches; the form's prompts are configurable per guild.

### F3 Characters

- One account, many characters. Players post, write and comment as a chosen character.
- Out-of-character posts show the character by default; setting.
- A player may hide which account owns a character. An admin can reveal it; every reveal is logged. The player is not notified by default; setting.
- Character profiles are wiki pages. Every story or thread tagged with a character appears on its profile.
- A character can be copied into another space and adapted. The copy is independent and shows "Adapted from \<original>"; the link follows the original's visibility, so viewers who can't see the original see "Adapted character" with no link.

### F4 Visibility engine

- Four tiers applied **per block inside a page**: guest, member, officer, personal secret.
- Personal secret: readable only by its author and the people the author allows. Officers and admins are not automatic readers.
- The server removes every block a viewer may not see before sending anything. The same filter covers pages, search, previews, notifications, link previews, revision history, exports and the API.
- Nested blocks take the stricter tier. A malformed or unknown block hides everything after it.
- Links and mentions never reveal a hidden target.
- Editing a page that contains blocks the editor can't see shows placeholders; the server keeps the hidden text intact on save.
- Back ends that can't do per-block visibility use the fallback: the public text links to a separately restricted page or thread holding the hidden part.

### F5 Editor

- Rich-text editing that stores Markdown underneath, with a raw-Markdown mode.
- Visibility blocks show as tinted, labeled boxes (never color alone).
- Mentions of characters, places and pages; image upload with required alt text and stripped location data.
- Full long-form writing works on phones.

### F6 Forum

- Categories and threads; in-character and out-of-character are separated **by forum area**.
- Category permissions follow the four tiers.

### F7 Wiki, lore and place pages

- Wiki pages with full revision history.
- Place pages for in-world locations, linked to the characters and stories set there.

### F8 Story archive and tagging

- Stories with topic tags readers can browse; tags support synonyms managed by officers.
- Stories carry the characters that appear in them.
- Players can export their own work at any time.

### F9 Galleries

- Personal and guild galleries.
- Every upload requires creator, source and license; uploads without them are rejected.

### F10 Spaces and packs

- An install holds multiple spaces (campaigns or worlds), separate by default.
- Install the core, then add packs any time with one click. Packs stack (a WoW guild can add a D&D pack).
- Everything optional is a pack: game vocabulary, maps, in-world timeline, character sheets, session planning, whiteboard.
- Packs declare typed data schemas with migrations.

### F11 Calendar and scheduling

- Built-in event calendar with RSVPs.
- Mirrors events to native Discord scheduled events and pulls Discord RSVPs back.
- Raid-Helper adapter: reads events and sign-ups by default; accepts Premium webhooks when the guild has Premium. The site is the source of truth.

### F12 Notifications

- Channels: on-site bell, Discord DM from the guild bot, optional email.
- Each player picks their mix, default **bell + Discord DM**; optional per-character overrides.
- A failed DM falls back to the bell.

### F13 Moderation and the community agreement

- Every account accepts the community agreement at sign-up.
- Officers approve applications; members' new posts can be held for approval (setting).
- Reporting, post deletion (old versions kept), warnings and removal.
- Admins can open a member's full post history to check patterns of behavior.
- The agreement lives in its own wiki section showing created and last-updated dates.
- Members propose changes and comment. Voting is a setting: anonymous thumbs up and down, both counts visible to everyone. Officers decide and can veto.
- Votes are anonymous by design: the system records only that an account voted, never which way.
- When the agreement changes, members are notified and must re-accept before posting.

### F14 Audit log

- Logs character reveals, visibility changes, edits and deletes, membership changes, and theme and setting changes.
- Readable by admins only. Retention is a setting, default 2 years.
- Never records how anyone voted.

### F15 Themes, accessibility, translation

- One theme styles the whole site. Built-in parts match exactly; adapter back ends get a close-match generated theme.
- Ladder: pick a theme from a gallery; edit theme variables (colors, fonts, spacing); or replace the front end entirely.
- Themes are shared through the `loomhall-themes` repo; authors choose their license and must credit sources.
- Every theme passes automated WCAG 2.2 AA and colorblind checks before it's accepted.
- The interface is fully translatable.

### F16 Data adapters, backups, export

| Slot | Default | Bring your own |
| --- | --- | --- |
| Database | SQLite | Postgres, MySQL |
| File storage | Server disk | S3-compatible (R2, B2) |
| Backups | Server disk, scheduled | Google Drive, OneDrive, S3 |
| Email | Off | SMTP, Resend |

- The project builds and maintains every adapter; guilds only point them at their own service.
- Each guild registers its own Google, Microsoft and Discord apps through a guided wizard.

### F17 API and clients

- Everything the site does is available through a versioned API documented in OpenAPI.
- The SvelteKit site is the reference front end; any other front end (React included) can be built on the same API.
- Generated API clients are published for TypeScript and other languages.

### F18 Search

- Search by title and tag, filtered by what the searcher may see.

### F19 Whiteboard pack

- Freeform board for brainstorming, running games and quick maps; one editor at a time in v1.
- Visibility is set per board.

## Cross-cutting requirements

| Area | Requirement |
| --- | --- |
| Cost | Runs for about 200 users on the cheapest VPS tier (1 GB RAM, about $5–7 a month); external services stay within free tiers by default |
| Footprint | Idle memory target under 100 MB for the core program, measured in the technical trial |
| Stack | Go back end (Huma, OpenAPI); SvelteKit reference front end shipped as prebuilt static files; Node tools rewritten in Go where needed so guilds never install Node; Python allowed only for tools, installers, helpers and theme scripts |
| Mobile | Every task works on a phone: reading, long-form writing, wiki editing, galleries, admin and moderation |
| Accessibility | WCAG 2.2 AA; colorblind-safe palettes; scalable fonts; screen-reader support; reduced-motion respected |
| Translation | All interface text translatable, right-to-left languages included |
| Security | All visibility, ownership and permission checks enforced on the server; tokens encrypted at rest and kept out of logs; Discord bot gets minimum permissions, never Administrator |
| Privacy | Trust-based in v1; no encryption of personal secrets; zero-trust deferred |
| Licensing | Core and adapters AGPL-3.0; theme authors choose their own license; no dependency under non-open-source licenses (e.g. BSL, Commons Clause, watermark-required SDKs) |
| Attribution | Every theme, image and asset carries creator, source and license; theme CI rejects entries without them |
| Bring your own tools | Every outside service sits behind a project-maintained adapter with a working default (F16) |
| Versioning | v1's data model must be ready for federation, live whiteboard and import in v1.5 |

**Name and places:** product **Loomhall** (one word); repos `martin-gleason/loomhall` and `martin-gleason/loomhall-themes`; domain loomhall.app if purchased.

## Process

Loomhall is built under Marty's Turing process (`code-process-review`).

- **Conventions:** a pinned copy of `conventions.md` from `code-process-review` is vendored into `loomhall/docs/conventions.md` and imported by the repo's CLAUDE.md. The vendored copy is never edited in place.
- **Specs and plans:** this spec is copied unchanged into `docs/specs` once ratified and is then immutable. All changes, additions and deltas go in `docs/plans`.
- **Work chunks:** Features (`F<N>`), Tasks (`F<N>-T<M>`), Chores (`C<N>`, done by Marty), Retrofits (`F<N>b`). Phases are tags, not containers. Gates are crossed only with Marty's explicit go-ahead.
- **PRs:** conventional commits with the feature ID in scope, branches named `<feature-id>/<slug>`, rebase-and-merge.
- **Learning level: 5%.** Claude authors; Marty reviews and logs every PR. Every PR carries short notes on why the Go is written the way it is, and occasionally hands Marty a small, easy piece of Go to write himself (for example, defining a variable or a small function).
- **Workflow:** deep reasoning for spec review and planning; adversarial review at the end of every feature and at the start of every session; no feature is done without a passing check shown as evidence.
- **First step after ratification:** a two-week technical trial proving the visibility engine with redacted editing, the editor on Android and iOS with a screen reader, and Go's memory use and page speed on a $5 server.

**Deep spec required** before build (decision-dense or security-critical gates):

| Gate | Decisions that drive it |
| --- | --- |
| F4 Visibility engine | Block syntax, tier nesting, redacted editing and merge-on-save, permission chaining, fallback for adapter back ends |
| F3 Characters | Ownership hiding, admin reveal and its logging, cross-space copies and origin-link visibility |
| F2 Identity and membership | Discord scopes, consent wording, approve-then-join flow, token storage |
| F13 Moderation and agreement | Anonymous-vote storage, re-accept flow, member-history access |
| F10 Spaces and packs | Pack format, schema migrations, how packs extend roles and vocabulary |
| F16 Data adapters | Adapter interface, SQLite and Postgres parity, backup format and restore |

## Out of scope and open questions

**Out of scope for v1:** social-network features (walls, likes, friend feeds); multiple guilds per install; encryption of personal secrets; per-element whiteboard visibility; full-text search; import; federation; automated AI-content screening; Battle.net login.

**Open questions**

- [x] Which packs ship with v1 besides the whiteboard (maps, timeline, character sheets, session planning, WoW vocabulary)?
- [ ] begin wtigh sqlite, but make sure the data layer is customizable. bring your own "X" applies to all aspects of this tool?
- [ ] Installer language: a POSIX shell script (no prerequisites) or Python?
- [ ] Access to `code-process-review`: this session can't see it, so the conventions copy is not yet vendored.
  - [ ] CC needs to pull the turing agent
- [ ] Default content tiers for new forum areas and wiki sections.

---

2026-10-22

\#AI/Claude
