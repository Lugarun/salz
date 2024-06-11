<script lang="ts">
	import { userStore } from '$lib/stores/user';
	import Icon from '@iconify/svelte';

	import NavLink from '$lib/components/NavLink.svelte';
	$: isLoggedIn = Object.hasOwn(userStore, 'accessToken');
</script>

<nav>
	<div class="left">
		<ul>
			<NavLink href="/"><Icon icon="tabler:assembly" /></NavLink>
		</ul>
	</div>
	<div class="right">
		<ul>
			<NavLink href="/leaderboard"><Icon icon="tabler:align-box-bottom-center" /></NavLink>
			<NavLink href="/matches"><Icon icon="tabler:video" /></NavLink>
			<NavLink href="/docs"><Icon icon="tabler:book-2" /></NavLink>
		</ul>

		<ul>
			{#if !isLoggedIn}
				<NavLink href="/login"><Icon icon="tabler:login" /></NavLink>
			{/if}
			{#if isLoggedIn}
				<NavLink href="/me"><Icon icon="tabler:user" /></NavLink>
				<NavLink href="/me/upload_bot"><Icon icon="tabler:hexagon-plus" /></NavLink>
			{/if}
		</ul>
	</div>
</nav>

<style>
	nav {
		padding: 1em 2em;
		display: flex;
		justify-content: space-between;
	}

	nav .left,
	nav .right {
		display: flex;
	}

	nav .right > * + * {
		margin-left: 1em;
	}

	nav ul {
		padding: 0;
		list-style-type: none;
		display: flex;
	}

	nav ul > :global(li + li) {
		margin-left: -10px;
	}
</style>
