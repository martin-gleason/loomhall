// Created: 2026-10-22 12:30 -05:00 · Edited: 2026-10-22 12:30 -05:00
// The reference front end builds to static files that the Go binary embeds (FR2,
// NFR14). adapter-static, never adapter-auto: there is no Node on a guild's server.
import adapter from '@sveltejs/adapter-static';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [
		sveltekit({
			compilerOptions: {
				// Force runes mode for the project, except for libraries. Can be removed in svelte 6.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes('node_modules') ? undefined : true
			},
			adapter: adapter({ fallback: '404.html' })
		})
	],
	test: {
		include: ['src/**/*.test.ts'],
		environment: 'node'
	}
});
