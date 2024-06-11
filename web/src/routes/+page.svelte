<script lang="ts">
	import NavLink from '$lib/components/NavLink.svelte';
	import FancyHeader from '$lib/components/FancyHeader.svelte';
	import Icon from '@iconify/svelte';
	import { userStore } from '$lib/stores/user';

	$: isLoggedIn = Object.hasOwn(userStore, 'accessToken');

  let leaders: { username: string; score: number }[] = [
    { username: "John Smith", score: 1234 },
    { username: "Adam Scott", score: 1233 },
    { username: "Bill Willy", score: 1232 },
  ];
</script>

<div class="content-container">
	<main>
		<h1>salz</h1>
		<h2>A bot programming game</h2>

		<menu>
			<ul>
				<NavLink href="/matches"><Icon icon="tabler:video" /> Watch Games</NavLink>
				<NavLink href="/docs"><Icon icon="tabler:book-2" /> How to play</NavLink>
				{#if !isLoggedIn}
					<NavLink href="/login"><Icon icon="tabler:login" /> Login / Sign up</NavLink>
				{/if}
				{#if isLoggedIn}
					<NavLink href="/me"><Icon icon="tabler:user" /> Go to profile</NavLink>
				{/if}
			</ul>
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
          {leader.username} {leader.score}
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
		padding: 2em 3em;
	}

	main menu ul {
		padding: 0;
		list-style: none;
		margin-top: 2em;
	}

	main menu ul :global(li) {
		margin-top: 0.5em;
	}
	main menu ul :global(li a:link) {
		width: 65%;
	}

	aside {
		border: 1px solid gray;
		padding: 0.3em 1.5em 1em 0em;
	}

  aside a:link, aside a:visited {
    text-decoration: none;
    color: inherit;
  }

  aside .leaderboard-list {
    margin-top: 1em;
    list-style: none;
    padding-left: 1.5em;
  }
</style>
