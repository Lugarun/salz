<script lang="ts">
	import { getInitials } from '$lib/name-wranglers';

	export let name;
</script>

<svg
  class="person-initials"
  version="1.1"
	xmlns="http://www.w3.org/2000/svg"
	preserveAspectRatio="xMinYMin meet"
  viewBox="0 0 120 120"
>
  <circle cx="60" cy="60" r="60" />
  <text x="50%" y="50%" text-anchor="middle" dominant-baseline="middle">
    {getInitials(name)}
  </text>
</svg>

<style>
svg {
  width: 2rem;
}

svg circle {
  fill: gray;
}

svg text {
  fill: white;
  font-family: sans-serif;
  font-size: 4rem;
}
</style>
