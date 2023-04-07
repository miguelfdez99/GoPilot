<script lang="ts">
  import { onMount } from 'svelte';
  import "@picocss/pico/css/pico.css"

  let currentView = 'home';

  onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
  });
</script>

<h1 class="text-2xl font-bold">Home</h1>

<p class="text-lg">This is the Home view.</p>

<h1 class="text-3xl font-bold underline">
  Hello world!
</h1>


