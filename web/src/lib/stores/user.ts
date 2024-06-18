import storage from '$lib/store';
import { derived } from 'svelte/store';

export interface UserStore {
	username: string;
	accessToken: string;
}

export const userStore = storage<UserStore | null>('userStore', { username: '', accessToken: '' });

export const isLoggedIn = derived(userStore, ($userStore) => Boolean($userStore?.accessToken));
