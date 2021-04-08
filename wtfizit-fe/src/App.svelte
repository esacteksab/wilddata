<script>
  import Assets from "./Assets.svelte";
  import TailwindCSS from './TailwindCSS.svelte';
  let assets;
</script>

<TailwindCSS />

<h1 class="text-8xl">Assets</h1>
<main>
    <Assets {assets} />
</main>