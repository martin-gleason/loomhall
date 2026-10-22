// Created: 2026-10-22 12:30 -05:00 · Edited: 2026-10-22 12:30 -05:00
// The smallest function worth a test: it proves the front-end runner exists (C1-T2).
// Real strings go through the translation layer (FR61); this one is a placeholder
// the harness can count and will be deleted when F15 lands.
export function greet(name: string): string {
	return `Well met, ${name}.`;
}
