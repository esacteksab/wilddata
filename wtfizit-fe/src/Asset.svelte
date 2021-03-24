<script>
    // create a prop
    export let asset;
</script>


<foo>
    <h1>Asset </h1>
    <small>
        Name: <b>{asset.Name}</b> <br/>
        Org: <b>{asset.Org}</b> <br/>
        {#if asset.Tags}
        Tags: 
            {#each asset.Tags as tag}
            <ul>
            <li>
                <b>md5sum: </b>{tag.md5sum}</li>
            
            <li>
                <b>sha256sum: </b>{tag.sha256sum}</li>
            </ul>
            {/each}
            {/if}
    </small><br/>
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
