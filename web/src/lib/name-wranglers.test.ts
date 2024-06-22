import { describe, expect, it } from 'vitest';
import { getInitials } from './name-wranglers';

describe('getInitials', () => {
	it("converts 'JohnSmith' to 'J'", () => expect(getInitials('JohnSmith')).toBe('J'));
	it("converts 'John Smith' to 'JS'", () => expect(getInitials('John Smith')).toBe('JS'));
	it("converts 'john' to 'j'", () => expect(getInitials('john')).toBe('j'));
	it("converts '风生水起' to '风'", () => expect(getInitials('风生水起')).toBe('风'));
	it("only returns at most 2 characters", () => expect(getInitials('John Smith The Third')).toBe('JS'));
	it("doesn't handle emojis well (which is fine)", () => expect(getInitials('😻')).not.toBe('😻'));
});
