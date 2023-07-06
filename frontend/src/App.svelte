<script lang="ts">
  import { onMount } from 'svelte';
  import { Home, Users, Packages, Monitor, Cron, Firewall, Backup, Process, Networking, Logs , Certs} from './imports';
  import "@picocss/pico/css/pico.css"
  let currentView = 'home';

  const views = {
  'backup': Backup,
  'certs': Certs,
  'cron': Cron,
  'firewall': Firewall,
  'home': Home,
  'logs': Logs,
  'monitoring': Monitor,
  'networking': Networking,
  'packages': Packages,
  'process': Process,
  'users': Users,
};

  let CurrentComponent = views[currentView];

  onMount(() => {
    const event = new CustomEvent('changeView', { detail: currentView });
    dispatchEvent(event);
  });
</script>

<div>
  <div class="sidebar">
    <button on:click={() => (currentView = 'backup', CurrentComponent = views[currentView])}>Backup</button>
    <button on:click={() => (currentView = 'certs', CurrentComponent = views[currentView])}>Certificates</button>
    <button on:click={() => (currentView = 'cron', CurrentComponent = views[currentView])}>Cron</button>
    <button on:click={() => (currentView = 'firewall', CurrentComponent = views[currentView])}>Firewall</button>
    <button on:click={() => (currentView = 'home', CurrentComponent = views[currentView])}>Home</button>
    <button on:click={() => (currentView = 'logs', CurrentComponent = views[currentView])}>Logs</button>
    <button on:click={() => (currentView = 'monitoring', CurrentComponent = views[currentView])}>Monitoring</button>
    <button on:click={() => (currentView = 'networking', CurrentComponent = views[currentView])}>Networking</button>
    <button on:click={() => (currentView = 'packages', CurrentComponent = views[currentView])}>Packages</button>
    <button on:click={() => (currentView = 'process', CurrentComponent = views[currentView])}>Processes</button>
    <button on:click={() => (currentView = 'users', CurrentComponent = views[currentView])}>Users</button>
  </div>

  <div class="main">
    <svelte:component this={CurrentComponent} />
  </div>
</div>

<style>

  /* .container {
    display: flex;
    height: 100vh;
  } */

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
    padding: 2rem;
    margin-left: 200px;
    color: white;
  }
</style>
