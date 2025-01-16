<script lang="ts">
    import {
        Backup,
        ScheduleBackup,
        OpenDir,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
    import { onMount } from "svelte";
    import { checkCommand } from "../functions/functions";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let options = {
        sourceDir: "",
        destDir: "",
        exclude: [],
    };

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    onMount(async () => {
        await checkCommand("rsync");
    });

    const backup = async () => {
        try {
            await Backup(options);
            await OpenDialogInfo("Backup created successfully");
        } catch (err) {
            await OpenDialogError(`Error creating backup: ${err.message}`);
        }
    };

    const selectDir = async () => {
        const dir = await OpenDir();
        if (dir) {
            options.sourceDir = dir;
        }
    };

    const selectDestDir = async () => {
        const dir = await OpenDir();
        if (dir) {
            options.destDir = dir;
        }
    };

    function onDialogClose() {
        dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    }
</script>

<div class="backup-container">
    <h2>Backups</h2>
    <div class="grid-container">
        <!-- Source Directory -->
        <div class="grid-item">
            <label for="sourceDir">Source Directory:</label>
        </div>
        <div class="grid-item">
            <input type="text" bind:value={options.sourceDir} />
            <button
                class="open-btn"
                title="Select Source Directory"
                on:click={selectDir}
            >
                <span class="material-icons">folder_open</span>
            </button>
        </div>

        <!-- Destination Directory -->
        <div class="grid-item">
            <label for="destDir">Destination Directory:</label>
        </div>
        <div class="grid-item">
            <input type="text" bind:value={options.destDir} />
            <button
                class="open-btn"
                title="Select Destination Directory"
                on:click={selectDestDir}
            >
                <span class="material-icons">folder_open</span>
            </button>
        </div>

        <!-- Backup Button -->
        <div class="grid-item">
            <button on:click={backup}>
                <span class="material-icons">backup</span>
                Backup Now
            </button>
        </div>
    </div>
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

    .backup-container {
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        font-family: sans-serif;
    }

    h2 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .grid-container {
        display: grid;
        gap: var(--spacing-md);
        grid-template-columns: 1fr 1fr;
        align-items: center;
    }

    .grid-item {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
    }

    label {
        font-weight: bold;
        color: var(--color-text-primary);
    }

    input {
        flex: 1;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
    }

    button {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        padding: var(--spacing-sm) var(--spacing-md);
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    button:hover {
        background-color: var(--color-bg-tertiary);
    }

    .open-btn {
        padding: var(--spacing-sm);
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        cursor: pointer;
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-accent-blue);
    }
</style>
