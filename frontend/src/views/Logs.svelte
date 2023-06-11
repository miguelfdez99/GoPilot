<script lang="ts">
    import { onMount } from "svelte";
    import { GetLogs } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import { writable } from "svelte/store";

    let logs = writable([]);
    let selectedLogType = writable("all");
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    let loading = writable(false);

    let fetchLogs = async (type) => {
        loading.set(true);
        try {
            const result = await GetLogs(type);
            logs.set(result.slice().reverse()); // Reverse the logs array
        } catch (error) {
            dialog = openDialog(
                dialog,
                "Error",
                `Failed to get logs: ${error}`
            );
        } finally {
            loading.set(false);
        }
    };

    selectedLogType.subscribe((value) => {
        fetchLogs(value);
    });

    onMount(() => {
        fetchLogs("all");
    });
</script>

<h1>Log viewer</h1>

<select bind:value={$selectedLogType}>
    <option value="all">All</option>
    <option value="system">System</option>
    <option value="application">Application</option>
</select>

{#if $loading}
    <div class="loading">
        <p>Loading...</p>
        <div class="spinner" />
    </div>
{:else}
    <div class="logs">
        {#each $logs as log}
            <pre>{log}</pre>
        {/each}
    </div>
{/if}

<style>
    h1 {
        color: white;
        font-size: 24px;
        margin-bottom: 10px;
    }

    select {
        color: white;
        font-size: 16px;
        margin-bottom: 20px;
    }

    .logs {
        border: 1px solid #dddbdb;
        padding: 0.5em;
        background-color: #030303fa;
        color: #070303;
        text-align: left;
    }

    .logs pre {
        white-space: pre-wrap;
        border-bottom: 1px solid #0a0909;
        background-color: rgb(44, 43, 43);
        padding: 0.5em;
        margin: 0;
        color: white;
    }

    .logs pre:last-child {
        border-bottom: none;
    }

    .loading {
        text-align: center;
    }

    .loading p {
        color: white;
        font-size: 16px;
        margin-bottom: 20px;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid #fff;
        border-radius: 50%;
        width: 30px;
        height: 30px;
        animation: spin 1s linear infinite;
        margin: 0 auto;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
