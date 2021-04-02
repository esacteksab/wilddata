<script>
  import { onMount } from "svelte";
  import Asset from "./Asset.svelte";
  import Assets from "./Assets.svelte";
  let assets;
  let asset;
</script>

<h1>Assets</h1>
<main>
    <Assets {assets} />
</main>