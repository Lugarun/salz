<script lang="ts">
	import { isLoggedIn } from '$lib/stores/user';
	import Icon from '@iconify/svelte';

	import NavLink from '$lib/components/NavLink.svelte';
	import Logo from './Logo.svelte';
</script>

<nav>
	<div class="left">
		<ul>
			<NavLink href="/"><Logo /></NavLink>
		</ul>
	</div>
	<div class="right">
		<ul>
			<NavLink href="/leaderboard"><Icon icon="tabler:align-box-bottom-center" /></NavLink>
			<NavLink href="/matches"><Icon icon="tabler:video" /></NavLink>
			<NavLink href="/docs"><Icon icon="tabler:book-2" /></NavLink>
		</ul>

		<ul>
      <NavLink href="/settings"><Icon icon="tabler:settings" /></NavLink>
			{#if !$isLoggedIn}
				<NavLink href="/login"><Icon icon="tabler:login" /></NavLink>
			{/if}
			{#if $isLoggedIn}
				<NavLink href="/me"><Icon icon="tabler:user" /></NavLink>
				<NavLink href="/me/upload_bot"><Icon icon="tabler:hexagon-plus" /></NavLink>
				<NavLink href="/logout"><Icon icon="tabler:logout" /></NavLink>
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
