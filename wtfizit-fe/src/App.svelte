<script>
  import { onMount } from "svelte";
  import Asset from "./Asset.svelte";
  import Assets from "./Assets.svelte";
  let assets;
  let asset;
</script>

<h1>The Seven Assets Svelte App</h1>
<main>
    <Assets {assets} />
</main>