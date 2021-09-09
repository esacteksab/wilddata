<!-- <script>
import { page } from "$app/stores";

    import { onMount } from "svelte";
    import Asset from "./OAsset.svelte";
    // define data holding variable
    export let assets = [];

onMount(async () => {
    await fetch(`http://localhost:5000/v1/orgs/${page.params.slug}/assets`)
    .then(r => r.json())
    .then(data => {
        assets = data;
        console.log('json:', data, 'assets:', assets)

    });
})

</script>

{#if assets}
    <ul>
    {#each assets as asset }
            <Asset {asset} />
    {/each}
    </ul>
{:else}
    <p class="loading">loading...</p>
{/if} -->

<script context="module">
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */
    import Assets from "./OAsset.svelte";
    export const assets = [] ;
	export async function load({ page, fetch, session, context }) {
		const url = `http://localhost:5000/v1/orgs/${page.params.slug}/assets`;
        console.log(url)
		const res = await fetch(url);

		if (res.ok) {
            console.log(res.json())
			return {
				props: {
					assets: await res.json()

				}
			};
		}

		return {
			status: res.status,
			error: new Error(`Could not load ${url}`)
		};
	}
</script>

<!-- {#if assets}
    <Asset { assets } />
    <ul>
    {#each assets as asset }
            <Asset {asset} />
    {/each}
    </ul>
{:else}
    <p class="loading">loading...</p>
{/if} -->

<Assets {assets} />