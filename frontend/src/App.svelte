<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte';
  import Home from './views/Home.svelte';
  import Users from './views/Users.svelte';
  import Packages from './views/Packages.svelte';
  import System from './views/SystemInformation.svelte';
  import Cron from './views/Cron.svelte';
  import Firewall from './views/Firewall.svelte';

  let currentView = 'home';
  const dispatch = createEventDispatcher();

  function handleViewChange(event: CustomEvent<string>) {
    currentView = event.detail;
  }

  onMount(() => {
    const event = new CustomEvent('changeView', { detail: currentView });
    dispatchEvent(event);
  });

</script>

<div class="container">
  <div class="sidebar">
    <button on:click={() => currentView = 'home'}>Home</button>
    <button on:click={() => currentView = 'users'}>Users</button>
    <button on:click={() => currentView = 'packages'}>Packages</button>
    <button on:click={() => currentView = 'system'}>System</button>
    <button on:click={() => currentView = 'cron'}>Cron</button>
    <button on:click={() => currentView = 'firewall'}>Firewall</button>
  </div>

  <div class="main">
    {#if currentView === 'home'}
      <Home />
    {:else if currentView === 'users'}
      <Users />
    {:else if currentView === 'packages'}
      <Packages />
    {:else if currentView === 'system'}
      <System />
    {:else if currentView === 'cron'}
      <Cron />
    {:else if currentView === 'firewall'}
      <Firewall />
    {/if}
  </div>
</div>

<style>
  @import url("https://unpkg.com/picocss@4.1.1/dist/pico.min.css");

  .container {
    display: flex;
    height: 100vh;
  }

  .sidebar {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    width: 200px;
    background-color: #0d6efd;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    overflow-y: auto;
    z-index: 100;
  }

  .sidebar button {
    margin-bottom: 0.5rem;
    padding: 0.5rem;
    background-color: #0d6efd;
    border: none;
    color: white;
    font-weight: bold;
    cursor: pointer;
    transition: all 0.3s;
  }

  .sidebar button:hover {
    background-color: #025ee4;
  }

  .sidebar button.active {
    background-color: #025ee4;
  }

  .main {
    flex: 1;
    padding: 1rem;
    margin-left: 200px;
  }
</style>
