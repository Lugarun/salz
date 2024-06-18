<script lang="ts">
	import NavLink from '$lib/components/NavLink.svelte';
	import FancyHeader from '$lib/components/FancyHeader.svelte';
	import Logo from '$lib/components/Logo.svelte';
	import Icon from '@iconify/svelte';
	import { isLoggedIn } from '$lib/stores/user';

	let leaders: { username: string; score: number }[] = [
		{ username: 'John Smith', score: 1234 },
		{ username: 'Adam Scott', score: 1233 },
		{ username: 'Bill Willy', score: 1232 }
	];
</script>

<div class="content-container">
	<main>
		<figure class="logo-container">
			<Logo />
		</figure>
		<h1>salz</h1>
		<h2>A bot programming game</h2>

		<menu>
			<NavLink href="/matches"><Icon icon="tabler:video" /> Watch Games</NavLink>
			<NavLink href="/docs"><Icon icon="tabler:book-2" /> How to play</NavLink>
			{#if !$isLoggedIn}
				<NavLink href="/login"><Icon icon="tabler:login" /> Login / Sign up</NavLink>
			{/if}
			{#if $isLoggedIn}
				<NavLink href="/me"><Icon icon="tabler:user" /> Go to profile</NavLink>
			{/if}
		</menu>
	</main>
	<aside>
		<a href="/leaderboard">
			<FancyHeader headerLevel="3">
				<Icon slot="icon" icon="tabler:align-box-bottom-center" />
				Leaderboard
			</FancyHeader>
		</a>
		<ol class="leaderboard-list">
			{#each leaders as leader, i}
				<li>
					<Icon icon={`tabler:hexagon-number-${i + 1}`} />
					{leader.username}
					{leader.score}
				</li>
			{/each}
		</ol>
	</aside>
</div>

<style>
	.content-container {
		height: 100vh;
		position: relative;
		display: flex;
		justify-content: space-around;
		align-items: center;
	}

	main {
		max-width: 40vw;
		text-align: center;
		border: 1px solid gray;
		padding: 3em 3em 2em 3em;
		position: relative;
	}

	main .logo-container {
		--slant-px: 30px;
		padding: 0.7em 2.3em 0.2em 2em;
		position: absolute;
		top: -3em;
		left: 50%;
		transform: translateX(-50%);
		background: lightgray;
		clip-path: polygon(var(--slant-px) 0, 100% 0, calc(100% - var(--slant-px)) 100%, 0 100%);
	}

	main .logo-container :global(.logo) {
		width: 4em;
	}

	main menu {
		padding: 0;
		list-style: none;
		margin-top: 2em;
	}

	main menu :global(li) {
		margin-top: 0.5em;
	}
	main menu :global(li a:link) {
		width: 65%;
	}

	aside {
		border: 1px solid gray;
		padding: 0.3em 1.5em 1em 0em;
	}

	aside a:link,
	aside a:visited {
		text-decoration: none;
		color: inherit;
	}

	aside .leaderboard-list {
		margin-top: 1em;
		list-style: none;
		padding-left: 1.5em;
	}

	@media screen and (max-width: 767px) {
		.content-container {
			display: initial;
			padding: 1em 2em;
		}

		main {
			width: 80vw;
			max-width: initial;
			margin: 0 auto;
		}

		main menu :global(li a:link) {
			width: 90%;
		}
	}
</style>
