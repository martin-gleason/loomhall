# Loomhall reference front end

Created: 2026-10-22 12:45 -05:00 · Edited: 2026-10-22 11:38 -05:00

SvelteKit on the static adapter. `npm run build` writes static files that the Go binary embeds, so a guild's server never runs Node (NFR14). `npx vitest run` runs the tests; the harness counts them. The API contract is the OpenAPI document the Go server publishes (FR68); this app is one client of it, replaceable by any other (FR69).
