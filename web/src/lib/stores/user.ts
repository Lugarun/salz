import { localWritable } from '$lib/store';
import { derived } from 'svelte/store';

export interface UserStore {
	username: string;
	accessToken: string;
}

export const initialValue = { username: '', accessToken: '' };

export const userStore = localWritable('userStore')<UserStore>(initialValue);

export const isLoggedIn = derived(userStore, ($userStore) => Boolean($userStore?.accessToken));
