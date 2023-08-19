<script lang="ts">
  import { onMount } from 'svelte';
  import { Home, Users, Packages, Monitor, Cron, Backup, Process, Logs , Certs, Alerts, Optimization, Services, Settings} from './imports';
  import "@picocss/pico/css/pico.css"
  import { settings } from './stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
    });

    $: {
    document.documentElement.style.setProperty('--main-font-size', fontSize);
    document.documentElement.style.setProperty('--main-color', color);
    document.documentElement.style.setProperty('--main-font-family', fontFamily);
    document.documentElement.style.setProperty('--main-bg-color', backgroundColor);
    document.documentElement.style.setProperty('--main-bg-color2', backgroundColor2);
  }
  let currentView = 'home';

  const views = {
  'alerts': Alerts,
  'backup': Backup,
  'certs': Certs,
  'cron': Cron,
  'home': Home,
  'logs': Logs,
  'monitoring': Monitor,
  'optimization': Optimization,
  'packages': Packages,
  'process': Process,
  'services': Services,
  'users': Users,
  'settings': Settings
};

  let CurrentComponent = views[currentView];

  onMount(() => {
    const event = new CustomEvent('changeView', { detail: currentView });
    dispatchEvent(event);
  });
</script>

<div>
  <div class="sidebar">
    <button on:click={() => (currentView = 'alerts', CurrentComponent = views[currentView])}>Alerts</button>
    <button on:click={() => (currentView = 'backup', CurrentComponent = views[currentView])}>Backup</button>
    <button on:click={() => (currentView = 'certs', CurrentComponent = views[currentView])}>Certs & Keys</button>
    <button on:click={() => (currentView = 'cron', CurrentComponent = views[currentView])}>Cron</button>
    <button on:click={() => (currentView = 'home', CurrentComponent = views[currentView])}>Home</button>
    <button on:click={() => (currentView = 'logs', CurrentComponent = views[currentView])}>Logs</button>
    <button on:click={() => (currentView = 'monitoring', CurrentComponent = views[currentView])}>Monitoring</button>
    <button on:click={() => (currentView = 'optimization', CurrentComponent = views[currentView])}>Optimization</button>
    <button on:click={() => (currentView = 'packages', CurrentComponent = views[currentView])}>Packages</button>
    <button on:click={() => (currentView = 'process', CurrentComponent = views[currentView])}>Processes</button>
    <button on:click={() => (currentView = 'services', CurrentComponent = views[currentView])}>Services</button>
    <button on:click={() => (currentView = 'users', CurrentComponent = views[currentView])}>Users</button>
    <button on:click={() => (currentView = 'settings', CurrentComponent = views[currentView])}>Settings</button>
  </div>

  <div class="main">
    <svelte:component this={CurrentComponent} />
  </div>
</div>

<style>
  .sidebar {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    width: 200px;
    background-color: var(--main-bg-color2);
    color: var(--main-color);
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
    background-color: var(--main-bg-color2);
    color: var(--main-color);
    border: none;
    font-weight: bold;
    cursor: pointer;
    transition: all;
  }

  .sidebar button:hover {
    background-color: var(--main-bg-color);
  }

  .main {
    flex: 1;
    padding: 2rem;
    margin-left: 200px;
    background-color: var(--main-bg-color);
  }
</style>
