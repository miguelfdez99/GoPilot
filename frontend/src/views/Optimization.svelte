<script lang="ts">
    import {
        DeleteTempFiles,
        RemoveUnusedPackages,
        CleanCachePackages,
        DuplicatedFiles,
        RemoveOldLogs,
        OpenDir,
        OpenDialogError,
        OpenDialogInfo,
    } from "../../wailsjs/go/backend/Backend";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
    import openicon from "../assets/images/open.png";

    let tempfilesdirpath: string = "";
    let expiredays: number;
    let duplicatedfilesdirpath: string = "";
    let oldlogsdays: string = "";
    let dialog = { showdialog: false, dialogtitle: "", dialogmessage: "" };

    const startDeletingTempFiles = () => {
        if (!tempfilesdirpath || !expiredays) {
            OpenDialogError(
                "Both directory path and number of days are required.",
            );
            return;
        }
        try {
            DeleteTempFiles(tempfilesdirpath, expiredays);
            OpenDialogInfo("Temp files removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing temp files: ${err.message}`);
        }
    };

    const startRemovingUnusedPackages = () => {
        try {
            RemoveUnusedPackages();
            OpenDialogInfo("Unused packages removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing unused packages: ${err.message}`);
        }
    };

    const startCleaningCachePackages = () => {
        try {
            CleanCachePackages();
            OpenDialogInfo("Cache packages cleaned successfully");
        } catch (err) {
            OpenDialogError(`Error cleaning cache packages: ${err.message}`);
        }
    };

    const startDeletingDuplicatedFiles = () => {
        if (!duplicatedfilesdirpath) {
            OpenDialogError("Directory path is required.");
            return;
        }
        try {
            DuplicatedFiles(duplicatedfilesdirpath);
            OpenDialogInfo("Duplicated files removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing duplicated files: ${err.message}`);
        }
    };

    const startRemovingOldLogs = () => {
        if (!oldlogsdays) {
            OpenDialogError("Number of days is required.");
            return;
        }
        try {
            RemoveOldLogs(oldlogsdays);
            OpenDialogInfo("Old logs removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing old logs: ${err.message}`);
        }
    };

    const selectDir = async () => {
        const dir = await OpenDir();
        if (dir) {
            tempfilesdirpath = dir;
        }
    };

    const selectPath = async () => {
        const dir = await OpenDir();
        if (dir) {
            duplicatedfilesdirpath = dir;
        }
    };

    function onDialogClose() {
        dialog = { showdialog: false, dialogtitle: "", dialogmessage: "" };
    }
</script>

<CustomDialog
    show={dialog.showdialog}
    title="Info"
    message={dialog.dialogmessage}
    onclose={onDialogClose}
/>

<div class="clean-system-container">
    <div class="clean-system-section">
        <h2>Clean System</h2>

        <div class="input-group">
            <div class="flex-input">
                <input
                    type="text"
                    bind:value={tempfilesdirpath}
                    placeholder="Enter a directory path for temporary files"
                />
                <button
                    class="open-btn"
                    title="Select path"
                    on:click={selectDir}
                >
                    <img src={openicon} alt="Open Dir" class="open-icon" />
                </button>
            </div>
            <input
                type="number"
                bind:value={expiredays}
                placeholder="Enter the number of days to keep temporary files"
            />
            <button on:click={startDeletingTempFiles}
                >Delete Temporary Files</button
            >
        </div>

        <div class="input-group">
            <div class="flex-input">
                <input
                    type="text"
                    bind:value={duplicatedfilesdirpath}
                    placeholder="Enter a directory path to find duplicated files"
                />
                <button
                    class="open-btn"
                    title="Select path"
                    on:click={selectPath}
                >
                    <img src={openicon} alt="Open Dir" class="open-icon" />
                </button>
            </div>
            <button on:click={startDeletingDuplicatedFiles}
                >Delete Duplicated Files</button
            >
        </div>

        <div class="input-group">
            <input
                type="text"
                bind:value={oldlogsdays}
                placeholder="Enter the number of days to keep old logs"
            />
            <button on:click={startRemovingOldLogs}>Remove Old Logs</button>
        </div>
    </div>

    <div class="clean-system-section">
        <h2>Package Cleaning</h2>

        <button on:click={startRemovingUnusedPackages}
            >Remove Unused Packages</button
        >
        <button on:click={startCleaningCachePackages}
            >Clean Cache Packages</button
        >
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

    .clean-system-container {
        display: flex;
        justify-content: space-between;
        gap: var(--spacing-lg);
        padding: var(--spacing-lg);
        background-color: var(--color-bg-primary);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .clean-system-section {
        flex: 1;
        padding: var(--spacing-lg);
        background-color: var(--color-bg-secondary);
        border-radius: 0.5rem;
    }

    h2 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .input-group {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-md);
    }

    .flex-input {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
    }

    input {
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    button {
        padding: var(--spacing-sm) var(--spacing-md);
        border: none;
        border-radius: 0.5rem;
        background-color: var(--color-accent-blue);
        color: var(--color-bg-primary);
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    button:hover {
        background-color: var(--color-accent-cyan);
    }

    .open-btn {
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        padding: var(--spacing-sm);
        border-radius: 0.5rem;
        cursor: pointer;
    }

    .open-btn:hover {
        background-color: var(--color-bg-tertiary);
    }
</style>
