<script lang="ts">
	import { onMount } from 'svelte';

  let matchId: string | null = null;

	onMount(() => {
		const urlParams = new URLSearchParams(window.location.search.replace('^?', ''));
    if (urlParams.has("id")) {
      matchId = urlParams.get("id");
    }
	});
</script>

Matches

{matchId}
