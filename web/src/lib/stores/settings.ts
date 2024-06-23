import { localWritable } from '$lib/store';

export interface Settings {
	apiBaseurl: string;
}

export const settingsStore = localWritable('settings')<Settings>({
	// apiBaseurl: 'https://api.salz.life',
	apiBaseurl: 'https://api.salz.life',
});

export const updateSettings = (newSettings: Partial<Settings>) =>
	settingsStore.update((store) => {
		return { ...store, ...newSettings };
	});
