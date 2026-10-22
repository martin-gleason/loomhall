// Created: 2026-10-22 12:30 -05:00 · Edited: 2026-10-22 12:30 -05:00
import { describe, expect, it } from 'vitest';
import { greet } from './hello';

describe('greet', () => {
	it('addresses the visitor by name', () => {
		expect(greet('Thrall')).toBe('Well met, Thrall.');
	});
});
