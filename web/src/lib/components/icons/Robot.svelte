<script lang="ts">
	export let hasCrown = false;

	const crownCenter = [20, 105];
	const crownSize = 4;
</script>

<svg
	version="1.1"
	viewBox="0 0 120 120"
	xmlns="http://www.w3.org/2000/svg"
	preserveAspectRatio="xMinYMin meet"
	class="robot"
>
	<circle cx="50%" cy="20" r="10" class="antenna" />

	<!-- replace these two lines with a rect so that we can fill it -->
	<rect x="57" y="30" width="6" height="10" class="antenna-neck" />

	<rect x="15" y="40" width="90" height="70" rx="2" class="head" />
	<ellipse cx="30%" cy="67" rx="12" ry="17" class="right-eye" />
	<ellipse cx="70%" cy="67" rx="12" ry="17" class="left-eye" />
	<rect x="45" y="90" width="30" height="13" class="mouth" />
	<line x1="50%" y1="90" x2="50%" y2="103" />
	<line x1="52" y1="90" x2="52" y2="103" />
	<line x1="68" y1="90" x2="68" y2="103" />

	{#if hasCrown}
		<polygon
			points={`
        ${crownCenter[0] - 3 * crownSize} ${crownCenter[1] + 3 * crownSize},
        ${crownCenter[0] - 4 * crownSize} ${crownCenter[1] - 1 * crownSize},
        ${crownCenter[0] - 2 * crownSize} ${crownCenter[1]},
        ${crownCenter[0]} ${crownCenter[1] - 3 * crownSize},
        ${crownCenter[0] + 2 * crownSize} ${crownCenter[1]},
        ${crownCenter[0] + 4 * crownSize} ${crownCenter[1] - 1 * crownSize},
        ${crownCenter[0] + 3 * crownSize} ${crownCenter[1] + 3 * crownSize}
      `}
			class="crown"
		/>
	{/if}
</svg>

<style>
	svg {
		width: 120px;
	}

	svg * {
		fill: transparent;
		stroke: black;
		stroke-width: 1;
	}

	.antenna,
	.antenna-neck,
	.head {
		fill: white;
	}

	.crown {
		fill: gold;
	}
</style>
