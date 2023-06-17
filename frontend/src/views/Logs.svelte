<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetLogs, ExportLogs } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import { writable } from "svelte/store";

    let showSearch = false;
    let showExport = false;
    let showBootNumber = false;
    let logs = writable([]);
    let selectedLogType = writable("all");
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    let loading = writable(false);
    let searchString = writable("");
    let bootNumber = writable(0);

    let fetchLogs = async (type, boot) => {
    loading.set(true);
    try {
        const result = await GetLogs(type, boot);
        const cleanedLogs = result.map(log => log.split(' ').slice(4).join(' ')); // Remove the timestamp and hostname
        logs.set(cleanedLogs.slice().reverse()); // Reverse the logs array1
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


    let exportLogs = async () => {
        loading.set(true);
        try {
            const filepath = "/tmp/prueba-logs.txt";
            await ExportLogs($selectedLogType, filepath, $bootNumber);
            dialog = openDialog(
                dialog,
                "Success",
                `Logs exported to: ${filepath}`
            );
        } catch (error) {
            dialog = openDialog(
                dialog,
                "Error",
                `Failed to export logs: ${error}`
            );
        } finally {
            loading.set(false);
        }
    };

    searchString.subscribe((value) => {
        if (value === "") {
            fetchLogs($selectedLogType, $bootNumber);
        } else {
            logs.set($logs.filter((log) => log.includes(value)));
        }
    });

    // Call fetchLogs when selectedLogType or bootNumber changes
    let unsubscribeSelectedLogType = selectedLogType.subscribe((value) => {
        fetchLogs(value, $bootNumber);
    });
    let unsubscribeBootNumber = bootNumber.subscribe((value) => {
        fetchLogs($selectedLogType, value);
    });

    onMount(() => {
        fetchLogs("all", 0);
    });

    onDestroy(() => {
        unsubscribeSelectedLogType();
        unsubscribeBootNumber();
    });
</script>

<div class="cont">
    <div class="control-buttons">
        <button
            on:click={() => (showSearch = !showSearch)}
            title="Toggle search">🔍</button
        >
        <button
            on:click={() => (showExport = !showExport)}
            title="Toggle export">💾</button
        >
        <button
            on:click={() => (showBootNumber = !showBootNumber)}
            title="Toggle boot number">🔢</button
        >
    </div>

    <h1>Log viewer</h1>

    {#if showSearch}
        <input bind:value={$searchString} placeholder="Search logs..." />
    {/if}

    {#if showExport}
        <button on:click={exportLogs}>Export Logs</button>
    {/if}

    {#if showBootNumber}
        <label>
            Boot number:
            <input type="number" bind:value={$bootNumber} />
        </label>
    {/if}

    <label>
        Log type:
        <select bind:value={$selectedLogType}>
            <option value="all">All</option>
            <option value="important">Important</option>
            <option value="system">System</option>
            <option value="hardware">Hardware</option>
            <option value="application">Application</option>
        </select>
    </label>

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
</div>

<style>
    h1 {
        color: white;
        font-size: 24px;
        margin-bottom: 10px;
    }

    select {
        color: white;
        background-color: #303030;
        font-size: 16px;
        margin-bottom: 20px;
    }

    button {
        background-color: #303030;
        border-color: #030303fa;
    }

    .cont {
        position: relative;
    }

    .control-buttons {
        position: absolute;
        right: 0;
        top: 0;
        display: flex;
        gap: 1em;
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
