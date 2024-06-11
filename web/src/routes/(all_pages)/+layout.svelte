<script lang="ts">
	import Nav from '$lib/components/Nav.svelte';
</script>

<Nav />
<div class="content-container"><slot></slot></div>

<style>
  .content-container {
    padding: 0 2em;
  }
</style>
