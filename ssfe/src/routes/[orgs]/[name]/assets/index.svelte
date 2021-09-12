<script context="module">
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */

	export async function load({ page, fetch, session, context }) {
		const url = `http://localhost:5000/v1/orgs/${page.params.name}/assets`;
		const res = await fetch(url);

        // console.log(url)
		// console.log(res.json())

		if (res.ok) return { props: { oassets: await res.json()	} };
		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	}
</script>

<script>
	export let oassets;
</script>

<main>

	<h1 class="text-8xl">Orgs Assets</h1>

	{#each oassets as asset}
		<ul class="pb-2 pt-2">
			Asset Name: <b>{asset.Name}</b> <br/>
			{#if asset.Tags}
			Tags: 
			<ul class="px-2">
				{#each Object.keys(asset.Tags[0]) as key}
				<li>
				{#each Object.values(asset.Tags) as value}
				  <b>{key}:</b>
				{#each Object.values(value[key]) as tag}
					{tag}
				{/each}
				 <br />
				{/each}
				{/each}
			</ul>
			{/if}
		</ul>
	{/each}
</main>

<style style lang="postcss">
	main {
		/* @apply text-center; */
		@apply p-4;
		@apply mx-auto;
	}

	h1 {
		@apply text-red-600;
		@apply uppercase;
		@apply text-6xl;
		@apply font-thin;
		@apply leading-tight;
		@apply my-16 mx-auto;
		@apply max-w-xs;
	}

	p {
		@apply max-w-xs;
		@apply my-8 mx-auto;
		@apply leading-snug;
	}

	@screen sm {
		h1 {
			@apply max-w-none;
		}

		p {
			@apply max-w-none;
		}
	}
</style>
