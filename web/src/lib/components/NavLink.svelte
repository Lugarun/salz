<script lang="ts">
	export let href: string;
</script>

<li>
	<a {href}>
		<slot></slot>
	</a>
</li>

<style>
	a:link,
	a:visited {
		text-decoration: none;
		color: inherit;
	}

	a:link {
		--slant-px: 10px;
		padding: 0.6em 1em;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		background: lightgray;
		transition: background 0.2s ease-in-out;
		clip-path: polygon(var(--slant-px) 0, 100% 0, calc(100% - var(--slant-px)) 100%, 0 100%);
	}

	a:hover {
		cursor: pointer;
		background: rgb(183, 183, 183);
	}

	a:focus {
		background: rgb(154, 154, 154);
	}
</style>
