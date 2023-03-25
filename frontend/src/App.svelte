<script lang="ts">
  import { onMount, createEventDispatcher } from 'svelte';
  import Home from './components/Home.svelte';
  import Users from './components/Users.svelte';
  import Packages from './components/Packages.svelte';
  import System from './components/SystemInformation.svelte';

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
    <ul>
      <li><button on:click={() => currentView = 'home'}>Home</button></li>
      <li><button on:click={() => currentView = 'users'}>Users</button></li>
      <li><button on:click={() => currentView = 'packages'}>Packages</button></li>
      <li><button on:click={() => currentView = 'system'}>System</button></li>
    </ul>
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
    background-color: #30abdb;
    text-align: center;
  }

  .main {
    flex: 1;
    padding: 1rem;
  }
</style>
