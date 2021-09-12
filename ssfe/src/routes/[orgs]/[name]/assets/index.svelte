<script context="module">
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */

	export async function load({ page, fetch, session, context }) {
		const url = `http://localhost:5000/v1/orgs/foo/assets`;
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
import Asset from "$lib/Asset.svelte";

	export let oassets;
</script>

<main>
	{#each oassets as asset}
		<ul class="pb-2 pt-2">
			Name: <b><a href="assets/{asset.Name}">{asset.Name}</a></b> <br/>
			Org: <b>{asset.Org}</b> <br/>
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

