<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import {
        GetLogs,
        ExportLogs,
        OpenFile,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
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
    let fetchTimeout: number;

    let fetchLogs = async (type, boot) => {
        if (fetchTimeout) {
            clearTimeout(fetchTimeout);
        }

        fetchTimeout = setTimeout(async () => {
            loading.set(true);
            try {
                const result = await GetLogs(type, boot);
                logs.set(result.slice().reverse());
            } catch (error) {
                if (error.includes("Invalid boot number")) {
                    await OpenDialogError(
                        `The provided boot number is not valid: ${boot}`,
                    );
                } else {
                    await OpenDialogError(`Failed to fetch logs: ${error}`);
                }
            } finally {
                loading.set(false);
            }
        }, 300);
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
            on:click={() =>
                (activeComponent =
                    activeComponent !== ActiveComponent.SEARCH
                        ? ActiveComponent.SEARCH
                        : ActiveComponent.NONE)}
            title="Toggle search"
        >
            <span class="material-icons">search</span>
        </button>
        <button
            on:click={() =>
                (activeComponent =
                    activeComponent !== ActiveComponent.EXPORT
                        ? ActiveComponent.EXPORT
                        : ActiveComponent.NONE)}
            title="Toggle export"
        >
            <span class="material-icons">save</span>
        </button>
        <button
            on:click={() =>
                (activeComponent =
                    activeComponent !== ActiveComponent.BOOT_NUMBER
                        ? ActiveComponent.BOOT_NUMBER
                        : ActiveComponent.NONE)}
            title="Toggle boot number"
        >
            <span class="material-icons">numbers</span>
        </button>
    </div>

    <h1>Log Viewer</h1>

    {#if activeComponent === ActiveComponent.SEARCH}
        <label>
            <p>Search:</p>
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
                <button
                    class="open-btn"
                    title="Choose File"
                    on:click={selectFile}
                    ><img
                        src={openIcon}
                        alt="Open Dir"
                        class="open-icon"
                    /></button
                >
            </div>
            <button class="export-button" on:click={exportLogs}
                >Export Logs</button
            >
        </label>
    {/if}

    {#if activeComponent === ActiveComponent.BOOT_NUMBER}
        <label>
            <p>Boot number:</p>
            <input type="number" bind:value={$bootNumber} />
        </label>
    {/if}

    <label>
        <p>Log type:</p>
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
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

    .cont {
        position: relative;
        padding: var(--spacing-lg);
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        font-family: sans-serif;
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .control-buttons {
        position: absolute;
        right: var(--spacing-lg);
        top: var(--spacing-lg);
        display: flex;
        gap: var(--spacing-sm);
    }

    button {
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        color: var(--color-text-primary);
        padding: var(--spacing-sm);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    button:hover {
        background-color: var(--color-bg-tertiary);
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-text-primary);
    }

    label p {
        color: var(--color-text-primary);
        margin-bottom: var(--spacing-sm);
    }

    .input-group {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-md);
    }

    input {
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        flex: 1;
    }

    .open-btn {
        padding: var(--spacing-sm);
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        cursor: pointer;
    }

    .open-icon {
        width: 1.2rem;
        height: 1.2rem;
    }

    .export-button {
        background-color: var(--color-accent-blue);
        color: var(--color-bg-primary);
        border: none;
        padding: var(--spacing-sm) var(--spacing-md);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .export-button:hover {
        background-color: var(--color-accent-cyan);
    }

    select {
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
        cursor: pointer;
        appearance: none; /* Remove default arrow */
        width: 100%;
        padding-right: 2rem; /* Space for custom arrow */
    }

    select:hover {
        border-color: var(--color-accent-blue);
    }

    select:focus {
        outline: none;
        border-color: var(--color-accent-blue);
        box-shadow: 0 0 0 2px rgba(122, 162, 247, 0.2);
    }

    .logs {
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        padding: var(--spacing-md);
        max-height: 400px;
        overflow-y: auto;
    }

    .logs pre {
        white-space: pre-wrap;
        margin: 0;
        padding: var(--spacing-sm) 0;
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    .logs pre:last-child {
        border-bottom: none;
    }

    .loading {
        text-align: center;
        padding: var(--spacing-lg);
    }

    .loading p {
        color: var(--color-text-primary);
        margin-bottom: var(--spacing-md);
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--color-accent-blue);
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
