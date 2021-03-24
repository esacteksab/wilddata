<script>
  import { onMount } from "svelte";
  import Continent from "./Continent.svelte";
  import Continents from "./Continents.svelte";
  let continents;
  let continent;
</script>

<h1>The Seven Continents Svelte App</h1>
<main>
    <Continents {continents} />
</main>