<script lang="ts">
    import { onMount } from "svelte";
    import {
        Home,
        Users,
        Packages,
        Monitor,
        Cron,
        Backup,
        Process,
        Logs,
        Certs,
        Optimization,
        Services,
    } from "./imports";

    let currentView = "home";

    const navItems = [
        { id: "home", label: "Home", icon: "home", component: Home },
        { id: "backup", label: "Backup", icon: "save", component: Backup },
        { id: "certs", label: "Certs & Keys", icon: "key", component: Certs },
        { id: "cron", label: "Cron", icon: "clock", component: Cron },
        { id: "logs", label: "Logs", icon: "article", component: Logs },
        {
            id: "monitoring",
            label: "Monitoring",
            icon: "computer",
            component: Monitor,
        },
        {
            id: "optimization",
            label: "Optimization",
            icon: "pageview",
            component: Optimization,
        },
        {
            id: "packages",
            label: "Packages",
            icon: "inventory",
            component: Packages,
        },
        {
            id: "process",
            label: "Processes",
            icon: "memory",
            component: Process,
        },
        {
            id: "services",
            label: "Services",
            icon: "dns",
            component: Services,
        },
        { id: "users", label: "Users", icon: "person", component: Users },
    ];

    $: CurrentComponent = navItems.find(
        (item) => item.id === currentView,
    )?.component;

    onMount(() => {
        const event = new CustomEvent("changeView", { detail: currentView });
        dispatchEvent(event);
    });
</script>

<div class="app">
    <nav class="sidebar">
        <div class="logo">
            <h1>GoPilot</h1>
        </div>
        <div class="nav-items">
            {#each navItems as item}
                <button
                    class="nav-item"
                    class:active={currentView === item.id}
                    on:click={() => (currentView = item.id)}
                >
                    <span class="icon material-icons">{item.icon}</span>
                    <span>{item.label}</span>
                </button>
            {/each}
        </div>
    </nav>

    <main class="content">
        <svelte:component this={CurrentComponent} />
    </main>
</div>

<style>
    :global(:root) {
        /* Tokyo Night theme colors */
        --color-bg-primary: #1a1b26;
        --color-bg-secondary: #16161e;
        --color-bg-tertiary: #1f2335;
        --color-text-primary: #a9b1d6;
        --color-text-secondary: #787c99;
        --color-accent-blue: #7aa2f7;
        --color-accent-purple: #9d7cd8;
        --color-accent-cyan: #7dcfff;
        --color-accent-green: #9ece6a;
        --color-accent-orange: #ff9e64;
        --color-accent-red: #f7768e;

        /* Layout */
        --sidebar-width: 250px;
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

    .app {
        display: flex;
        min-height: 100vh;
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
    }

    .sidebar {
        position: fixed;
        width: var(--sidebar-width);
        height: 100vh;
        background-color: var(--color-bg-secondary);
        padding: var(--spacing-md);
        display: flex;
        flex-direction: column;
        gap: var(--spacing-lg);
    }

    .logo {
        padding: var(--spacing-md);
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    .logo h1 {
        color: var(--color-accent-blue);
        font-size: 1.5rem;
        font-weight: bold;
    }

    .nav-items {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-sm);
    }

    .nav-item {
        display: flex;
        align-items: center;
        gap: var(--spacing-md);
        padding: var(--spacing-md);
        color: var(--color-text-secondary);
        background: transparent;
        border: none;
        border-radius: 0.5rem;
        cursor: pointer;
        transition: all 0.2s ease-in-out;
        width: 100%;
        text-align: left;
    }

    .nav-item:hover {
        background-color: var(--color-bg-tertiary);
        color: var(--color-text-primary);
    }

    .nav-item.active {
        background-color: var(--color-bg-tertiary);
        color: var(--color-accent-blue);
    }

    .content {
        margin-left: var(--sidebar-width);
        padding: var(--spacing-lg);
        width: calc(100% - var(--sidebar-width));
    }

    .icon {
        width: 1.5rem;
        height: 1.5rem;
        display: inline-flex;
        align-items: center;
        justify-content: center;
    }
</style>
