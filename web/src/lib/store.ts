import { browser } from '$app/environment';
import type { Writable } from 'svelte/store';
import { writable, get } from 'svelte/store';

export const localWritable =
	(key: string) =>
	<T>(initValue: T): Writable<T> => {
		const store = writable(initValue);
		if (!browser) return store;

		const storedValueStr = localStorage.getItem(key);
		if (storedValueStr != null) store.set(JSON.parse(storedValueStr));

		store.subscribe((val) => {
			if ([null, undefined].includes(val as null | undefined)) {
				localStorage.removeItem(key);
			} else {
				localStorage.setItem(key, JSON.stringify(val));
			}
		});

		window.addEventListener('storage', () => {
			const storedValueStr = localStorage.getItem(key);
			if (storedValueStr == null) return;

			const localValue: T = JSON.parse(storedValueStr);
			if (localValue !== get(store)) store.set(localValue);
		});

		return store;
	};
