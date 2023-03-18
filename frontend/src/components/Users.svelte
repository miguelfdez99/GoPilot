<script lang="ts">
  import { onMount } from 'svelte';

  let currentView = 'users';

  onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
  });
</script>

<h1>Users</h1>

<p>This is the Users view.</p>
