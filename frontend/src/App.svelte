<script lang="ts">
  import { onMount } from 'svelte';
  import Home from './views/Home.svelte';
  import Users from './views/Users.svelte';
  import Packages from './views/Packages.svelte';
  import Monitor from './views/Monitoring.svelte';
  import Cron from './views/Cron.svelte';
  import Firewall from './views/Firewall.svelte';
  import Backup from './views/Backup.svelte';
  import Process from './views/ProcInfo.svelte';
  import "@picocss/pico/css/pico.css"
  let currentView = 'home';

  const views = {
    'home': Home,
    'users': Users,
    'packages': Packages,
    'monitoring': Monitor,
    'cron': Cron,
    'firewall': Firewall,
    'backup': Backup,
    'process': Process
  };

  let CurrentComponent = views[currentView];

  onMount(() => {
    const event = new CustomEvent('changeView', { detail: currentView });
    dispatchEvent(event);
  });
</script>

<div class="container">
  <div class="sidebar">
    <button on:click={() => (currentView = 'home', CurrentComponent = views[currentView])}>Home</button>
    <button on:click={() => (currentView = 'users', CurrentComponent = views[currentView])}>Users</button>
    <button on:click={() => (currentView = 'packages', CurrentComponent = views[currentView])}>Packages</button>
    <button on:click={() => (currentView = 'monitoring', CurrentComponent = views[currentView])}>Monitoring</button>
    <button on:click={() => (currentView = 'cron', CurrentComponent = views[currentView])}>Cron</button>
    <button on:click={() => (currentView = 'firewall', CurrentComponent = views[currentView])}>Firewall</button>
    <button on:click={() => (currentView = 'backup', CurrentComponent = views[currentView])}>Backup</button>
    <button on:click={() => (currentView = 'process', CurrentComponent = views[currentView])}>Processes</button>
  </div>

  <div class="main">
    <svelte:component this={CurrentComponent} />
  </div>
</div>

<style>

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
    background-color: #1f1f1f;
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
    background-color: #1f1f1f;
    border: none;
    color: #f1f1f1;
    font-weight: bold;
    cursor: pointer;
    transition: all;
  }

  .sidebar button:hover {
    background-color: #3d3d3d;
  }

  .main {
    flex: 1;
    padding: 1rem;
    margin-left: 200px;
    color: white;
  }
</style>
