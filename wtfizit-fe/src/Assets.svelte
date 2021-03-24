<script>
    import { onMount } from "svelte";
    import Asset from "./Asset.svelte";
    // define data holding variable
    export let assets = [];

onMount(async () => {
    await fetch(`http://localhost:3000/v1/artifacts`)
    .then(r => r.json())
    .then(data => {
        assets = data;
        console.log('json:', data)
    });
})

</script>

{#if assets}
    {#each assets as asset }
    <ul>
        <li>
            <Asset {asset} />
        </li>
    </ul>
    {/each}
{:else}
    <p class="loading">loading...</p>
{/if}