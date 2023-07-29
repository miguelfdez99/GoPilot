<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetLogs, ExportLogs, OpenFile, OpenDialogInfo, OpenDialogError } from "../../wailsjs/go/backend/Backend";
    import { writable } from "svelte/store";
    import openIcon from "../assets/images/open.png";

    let filepath: string = "";
    let logs = writable([]);
    let selectedLogType = writable("all");
    let loading = writable(false);
    let searchString = writable("");
    let bootNumber = writable(0);
    const ActiveComponent = { NONE: 0, SEARCH: 1, EXPORT: 2, BOOT_NUMBER: 3 };
    let activeComponent = ActiveComponent.NONE;

    let fetchLogs = async (type, boot) => {
        loading.set(true);
        try {
            const result = await GetLogs(type, boot);
            const cleanedLogs = result.map(log => log.split(' ').slice(4).join(' ')); 
            logs.set(cleanedLogs.slice().reverse()); 
        } catch (error) {
            await OpenDialogError(`Failed to fetch logs: ${error}`);
        } finally {
            loading.set(false);
        }
};


    let exportLogs = async () => {
        loading.set(true);
        try {
            await ExportLogs($selectedLogType, $bootNumber, filepath);
            await OpenDialogInfo(`Logs exported successfully to ${filepath}`);
        } catch (error) {
            await OpenDialogError(`Failed to export logs: ${error}`);
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

    const selectFile = async () => {
        const file = await OpenFile();
        filepath = file;
    };
</script>

<div class="cont">
    <div class="control-buttons">
        <button
            on:click={() => activeComponent = activeComponent !== ActiveComponent.SEARCH ? ActiveComponent.SEARCH : ActiveComponent.NONE}
            title="Toggle search">🔍</button
        >
        <button
            on:click={() => activeComponent = activeComponent !== ActiveComponent.EXPORT ? ActiveComponent.EXPORT : ActiveComponent.NONE}
            title="Toggle export">💾</button
        >
        <button
            on:click={() => activeComponent = activeComponent !== ActiveComponent.BOOT_NUMBER ? ActiveComponent.BOOT_NUMBER : ActiveComponent.NONE}
            title="Toggle boot number">🔢</button
        >
    </div>

    <h1>Log viewer</h1>

    {#if activeComponent === ActiveComponent.SEARCH}
    <label >
        Search:
        <input bind:value={$searchString} placeholder="Search logs..." />
    </label>
    {/if}

    {#if activeComponent === ActiveComponent.EXPORT}
    <label>
        <p>Export:</p>
        <div class="input-group">
        <input
                type="text"
                bind:value={filepath}
                placeholder="Enter a file path"
            />
        <button class="open-btn" title="Choose File" on:click={selectFile}
                ><img src={openIcon} alt="Open Dir" class="open-icon" /></button
            >
        </div>
        <button on:click={exportLogs}>Export Logs</button>
    </label>
    
    {/if}

    {#if activeComponent === ActiveComponent.BOOT_NUMBER}
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

    label p {
        color: white;
    }

    .input-group {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 1em;
    }

    button.open-btn {
        background: none;
        width: 2.5rem;
        height: 2.5rem;
        padding: 0;
        margin: 0;
        margin-left: 10px;
        margin-bottom: 15px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    button.open-btn:hover {
        background-color: #131313;
    }

    input {
        padding: 0.7em;
        border: 0;
        border-radius: 4px;
        background: #333;
        color: #fff;
        width: 100%;
        box-sizing: border-box;
        margin-bottom: 0.5em;
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
        overflow: auto;
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
