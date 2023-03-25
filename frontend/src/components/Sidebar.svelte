<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  function handleButtonClick(view: string) {
    dispatch('changeView', { detail: view });
  }
</script>

<ul>
  <li><button on:click={() => handleButtonClick('home')}>Home</button></li>
  <li><button on:click={() => handleButtonClick('users')}>Users</button></li>
  <li><button on:click={() => handleButtonClick('packages')}>Packages</button></li>
</ul>