<script lang="ts">
  import { onMount } from 'svelte';

  let currentView = 'home';

  onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
  });
</script>

<h1>Home</h1>

<p>This is the Home view.</p>
<h1 class="text-3xl font-bold underline">
  Hello world!
</h1>

