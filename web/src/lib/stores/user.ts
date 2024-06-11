import { browser } from '$app/environment';
import { writable } from 'svelte/store';

export interface UserStore {
	username: string;
	accessToken: string;
}

const storage = browser && localStorage.getItem('userStore');

export const userStore = writable<UserStore>(storage ? JSON.parse(storage) : {});

userStore.subscribe((val) => browser && localStorage.setItem('userStore', JSON.stringify(val)));
