<script lang="ts">
	import { goto } from '$app/navigation';
	import { userStore, initialValue } from '$lib/stores/user';

	$userStore = initialValue;

  goto('/');
</script>
