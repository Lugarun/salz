<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import debounce from 'lodash/debounce';
	import cloneDeep from 'lodash/cloneDeep';

	import FancyHeader from '$lib/components/FancyHeader.svelte';
	import Button from '$lib/components/Button.svelte';

	import { updateSettings, settingsStore } from '$lib/stores/settings';

	const settingsClone = cloneDeep($settingsStore);

	const errorMessages = {
		otherApiUrl: '',
	};

	let apiUrlChoice =
		settingsClone.apiBaseurl === 'https://api.salz.life'
			? 'https://api.salz.life'
			: settingsClone.apiBaseurl.startsWith('http://localhost')
				? 'http://localhost'
				: '';
	let localhostPort = settingsClone.apiBaseurl.startsWith('http://localhost')
		? settingsClone.apiBaseurl.substring(settingsClone.apiBaseurl.lastIndexOf(':') + 1)
		: 8080;
	let otherApiUrl = apiUrlChoice === '' ? settingsClone.apiBaseurl : '';
	$: hasErrors = Object.values(errorMessages).some((v) => v !== '');
	$: {
		if (apiUrlChoice === '' && otherApiUrl === '') {
			errorMessages.otherApiUrl = 'Must not be empty';
		}
	}

	$: {
		switch (apiUrlChoice) {
			case 'https://api.salz.life':
				settingsClone.apiBaseurl = 'https://api.salz.life';
				break;
			case 'http://localhost':
				settingsClone.apiBaseurl = `http://localhost:${localhostPort}`;
				break;
			case '':
				settingsClone.apiBaseurl = otherApiUrl;
				break;
		}
	}

	function saveSettings() {
		updateSettings(settingsClone);
	}

	onMount(() => {
		const apiUrlInputBlurHandler = () => {
			try {
				if (otherApiUrl !== '') new URL(otherApiUrl);
				errorMessages.otherApiUrl = '';
			} catch {
				errorMessages.otherApiUrl = 'Not a valid URL';
			}
		};
		const debouncedApiUrlInputBlurHandler = debounce(apiUrlInputBlurHandler);

		const otherApiUrlInput = document.getElementById('custom-api-url');
		otherApiUrlInput?.addEventListener('blur', debouncedApiUrlInputBlurHandler);
		window.addEventListener('unload', () => {
			otherApiUrlInput?.removeEventListener('blur', debouncedApiUrlInputBlurHandler);
		});
	});
</script>

<main>
	<FancyHeader headerLevel="1">
		<Icon slot="icon" icon="tabler:settings" />
		Settings
	</FancyHeader>

	<form>
		<fieldset>
			<legend>API URL</legend>

			<div>
				<input
					type="radio"
					id="salz-life"
					name="apiUrl"
					bind:group={apiUrlChoice}
					value="https://api.salz.life"
					checked={apiUrlChoice === 'https://api.salz.life'}
				/>
				<label for="salz-life">https://api.salz.life</label>
			</div>

			<div>
				<input
					type="radio"
					id="localhost"
					name="apiUrl"
					bind:group={apiUrlChoice}
					value="http://localhost"
					checked={apiUrlChoice === 'http://localhost'}
				/>
				<label for="localhost">http://localhost:</label>
				<input
					type="number"
					id="localhost-port"
					name="apiUrl-locdalhost-port"
					bind:value={localhostPort}
					disabled={apiUrlChoice !== 'http://localhost'}
				/>
			</div>

			<div>
				<input
					type="radio"
					id="other"
					name="apiUrl"
					bind:group={apiUrlChoice}
					value=""
					checked={apiUrlChoice === ''}
				/>
				<input
					type="text"
					id="custom-api-url"
					name="custom-api-url"
					bind:value={otherApiUrl}
					placeholder="Other (please specify)"
					disabled={apiUrlChoice !== ''}
				/>
				{#if apiUrlChoice === ''}
					<span role="alert" class="error-msg">
						{errorMessages.otherApiUrl}
					</span>
				{/if}
			</div>
		</fieldset>

		<fieldset class="form-controls">
			<Button onClick={saveSettings} isDisabled={hasErrors}>Save</Button>
			{#if hasErrors}
				<span role="alert" class="error-msg">There are errors in the form.</span>
			{/if}
		</fieldset>
	</form>
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

	main > form {
		margin-top: 2em;
		padding-left: 1.5em;
	}

  fieldset {
    padding: 0 1rem 1rem 1rem;
  }

	fieldset + fieldset {
		margin-top: 1.5rem;
	}

	.form-controls {
		border: 0;
	}

	.error-msg {
		color: red;
	}
</style>
