<script lang="ts">
	import { Temporal } from 'temporal-polyfill';

	import FancyHeader from '$lib/components/FancyHeader.svelte';
	import { userStore } from '$lib/stores/user';
	import RoundInitial from '$lib/components/RoundInitial.svelte';
	import Robot from '$lib/components/icons/Robot.svelte';
	import CheckeredPattern from '$lib/components/CheckeredPattern.svelte';
	import Icon from '@iconify/svelte';

	// Mocked bot results
	type Player = {
		id: string;
		username: string;
		profileImg?: string;
	};
	type Bot = {
		id: string;
		owner?: Player;
	};
	type MatchResult = {
		id: string;
		timestamp: Temporal.Instant;
		bot1: Bot;
		bot2: Bot;
		winner: 'bot1' | 'bot2';
	};
	$: botResults = [
		{
			id: 'match111',
			timestamp: Temporal.Instant.from('2024-07-19T00:31:19Z'),
			bot1: {
				id: 'abc123',
				owner: {
					id: 'me1',
					username: 'John Smith',
				},
			},
			bot2: {
				id: 'bar312',
				owner: {
					id: 'foobar1',
					username: 'Foo Bar 1',
				},
			},
			winner: 'bot1',
		},
		{
			id: 'match110',
			timestamp: Temporal.Instant.from('2024-07-19T00:26:07Z'),
			bot1: {
				id: 'abc123',
				owner: {
					id: 'me1',
					username: 'John Smith',
				},
			},
			bot2: {
				id: 'bar314',
				owner: {
					id: 'foobar2',
					username: 'Foo Bar 2',
				},
			},
			winner: 'bot2',
		},
		{
			id: 'match109',
			timestamp: Temporal.Instant.from('2024-07-17T16:59:07Z'),
			bot1: {
				id: 'abc123',
				owner: {
					id: 'me1',
					username: 'John Smith',
				},
			},
			bot2: {
				id: 'bar311',
				owner: {
					id: 'foobar3',
					username: 'Foo Bar 3',
				},
			},
			winner: 'bot1',
		},
	] as MatchResult[];

	const localTimezoneId = Temporal.Now.timeZoneId();

	$: buckets = botResults.reduce(
		(acc, cur) => {
			const zonedDt = cur.timestamp.toZonedDateTimeISO(localTimezoneId);
			const localizedDate = zonedDt.toPlainDate().toLocaleString();
			if (Object.hasOwn(acc, localizedDate)) {
				acc[localizedDate] = [...acc[localizedDate], cur];
			} else {
				acc[localizedDate] = [cur];
			}

			return acc;
		},
		{} as { [key: string]: MatchResult[] },
	);
</script>

<main>
	<FancyHeader headerLevel="1">
		<RoundInitial slot="icon" name={$userStore?.username} />
		{$userStore?.username}
	</FancyHeader>

	<div class="content-container">
		<article class="my-match-results">
			<h2>Bot Results</h2>
			{#each Object.entries(buckets) as bucket}
				<section class="match-result-group">
					<h3>
						<span>{bucket[0]}</span>
					</h3>
					{#each bucket[1] as matchResult}
						<article class="match-result">
							<time
								datetime={matchResult.timestamp
									.toZonedDateTimeISO(localTimezoneId)
									.toPlainTime()
									.toLocaleString()}
							>
								{matchResult.timestamp
									.toZonedDateTimeISO(localTimezoneId)
									.toPlainTime()
									.toLocaleString()}
							</time>

							<div class="match-result-summary">
								<figure class="bot-wrapper">
									<CheckeredPattern color={matchResult.winner === 'bot1' ? 'red' : 'blue'} />
									<Robot hasCrown={matchResult.winner === 'bot1'} />
									<RoundInitial name={matchResult.bot1.owner?.username} />
								</figure>
								vs
								<figure class="bot-wrapper">
									<CheckeredPattern color={matchResult.winner === 'bot2' ? 'red' : 'blue'} />
									<Robot hasCrown={matchResult.winner === 'bot2'} />
									<RoundInitial name={matchResult.bot2.owner?.username} />
								</figure>
							</div>
							<a href={`/matches?id=${matchResult.id}`}><Icon icon="tabler:video" /></a>
						</article>
					{/each}
				</section>
			{/each}
		</article>
	</div>
</main>

<style>
	main {
		border: 1px solid gray;
		padding: 0.3em 1.5em 1em 0em;
		border-radius: 5px 5px 10px 10px;
		min-height: 30vh;
		width: 60vw;
		margin: 2em auto 5em auto;
		background: white;
	}

	main :global(.fancy-header) {
		max-width: 60%;
	}

	main :global(.fancy-header .person-initials) {
		width: 1em;
	}

	.content-container {
		padding-top: 1.5em;
		padding-left: 1.5em;
	}

	.match-result-group + .match-result-group {
		margin-top: 2em;
	}

	.match-result-group > h3 {
		position: relative;
		padding-left: 1em;
		font-size: 1.6rem;
	}

	.match-result-group > h3::before {
		content: '';
		position: absolute;
		top: 50%;
		left: 0;
		width: 100%;
		height: 3px;
		background: black;
	}

	.match-result-group > h3 span {
		position: relative;
		padding: 0 1em;
		background: white;
		z-index: 1;
	}

	.match-result {
		display: grid;
		grid-template-columns: 1fr 2fr 1fr;
		column-gap: 1.5em;
		align-items: center;
		position: relative;
	}

	.match-result *:nth-child(3) {
		justify-self: center;
		font-size: 3rem;
	}

	.match-result + .match-result {
		margin-top: 1rem;
	}
	.match-result + .match-result::before {
		content: '';
		width: 70%;
		height: 1px;
		background: linear-gradient(
			0.25turn,
			transparent,
			15%,
			rgba(100, 100, 100, 0.5),
			85%,
			transparent
		);
		position: absolute;
		top: -0.5rem;
		left: 50%;
		transform: translateX(-50%);
	}

	.match-result-summary {
		display: flex;
		justify-content: space-around;
		align-items: center;
	}
	.bot-wrapper {
		position: relative;
		width: 120px;
		height: 120px;
		display: inline-flex;
		justify-content: center;
		align-items: center;
	}
	.bot-wrapper :global(.checkered-pattern) {
		position: absolute;
		top: 0;
		left: 0;
		width: inherit;
	}
	.bot-wrapper :global(.robot) {
		position: relative;
		width: 100px;
		z-index: 1;
	}
	.bot-wrapper :global(.person-initials) {
		position: absolute;
		z-index: 2;
		top: 65%;
		right: 10px;
	}
</style>
