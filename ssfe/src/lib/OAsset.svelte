<script>
    // create a prop
    export let assets;
</script>


<foo class="py-4 text-lg">

        Name: <b><a href="assets/{assets}">{assets}</a></b> <br/>
        <!-- Org: <b>{asset.Org}</b> <br/>
        {#if asset.Tags}
        Tags: 
        <ul class="px-2 py-4">
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
        {/if} -->
<br/>
    <hr />
</foo>