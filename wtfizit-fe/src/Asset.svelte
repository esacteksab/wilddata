<script>
import Assets from "./Assets.svelte";

    // create a prop
    export let asset;
    //console.log(asset.Tags)
    console.log(asset.Tags[0])
</script>


<foo>
    <small>
        Name: <b>{asset.Name}</b> <br/>
        Org: <b>{asset.Org}</b> <br/>
        {#if asset.Tags}
        Tags: 
        <ul>
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
    </small><br/>
    <hr />
</foo>

<style>
  foo {
    margin: 0 0 1em 0;
  }
  h1 {
    font-size: 1.4em;
    margin: 0;
    display: block;
  }
</style>
