<script lang="ts">
    import { onMount } from 'svelte';
    import { ListFirewallRules } from '../../wailsjs/go/main/App.js';

    let rules = [];
    let error = null;

    let currentView = 'cron';

    function handleViewChange(event: CustomEvent<string>) {
    currentView = event.detail;
    }


    onMount(() => {
        addEventListener('changeView', (event: CustomEvent) => {
            currentView = event.detail;
        });
        listFirewallRules()
    });

    async function listFirewallRules() {
        try {
            rules = await ListFirewallRules();
            error = null;
        } catch (err) {
            rules = [];
            error = err;
        }
    }
</script>

<main>
    <h1>Current Firewall Rules</h1>

    {#if error}
        <p class="error">{error.message}</p>
    {:else if rules.length > 0}
        <textarea readonly>{rules.join("\n")}</textarea>
    {:else}
        <p>No firewall rules found.</p>
    {/if}

    <button on:click={listFirewallRules}>Refresh</button>
</main>

<style>
    textarea {
        width: 100%;
        height: 200px;
        font-family: monospace;
        white-space: pre-wrap;
        overflow-y: scroll;
    }

    .error {
        color: red;
        font-weight: bold;
    }
</style>
