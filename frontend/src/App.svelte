<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte';
  import Home from './views/Home.svelte';
  import Users from './views/Users.svelte';
  import Packages from './views/Packages.svelte';
  import System from './views/SystemInformation.svelte';
  import Users2 from './views/Users2.svelte';
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
    <button on:click={() => currentView = 'home'} style="width: 100%">Home</button>
    <button on:click={() => currentView = 'users'} style="width: 100%">Users</button>
    <button on:click={() => currentView = 'packages'} style="width: 100%">Packages</button>
    <button on:click={() => currentView = 'system'} style="width: 100%">System</button>
    <button on:click={() => currentView = 'users2'} style="width: 100%">Users2</button>
    <button on:click={() => currentView = 'cron'} style="width: 100%">Cron</button>
    <button on:click={() => currentView = 'firewall'} style="width: 100%">Firewall</button>
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
    {:else if currentView === 'users2'}
      <Users2 />
    {:else if currentView === 'cron'}
      <Cron />
    {:else if currentView === 'firewall'}
      <Firewall />
    {/if}
  </div>
</div>

<style>
  .container {
    display: flex;
    height: 100vh;
  }

  .sidebar {
    width: 200px;
    height: 100%;
    background-color: #3e778d;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .sidebar button {
    margin-bottom: 0.5rem;
    padding: 0.5rem;
    background-color: #2c6278;
    border: none;
    color: white;
    font-weight: bold;
    cursor: pointer;
  }

  .sidebar button:hover {
    background-color: #245468;
  }

  .main {
    flex: 1;
    padding: 1rem;
  }
</style>
