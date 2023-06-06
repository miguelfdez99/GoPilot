<script lang="ts">
    import { Backup } from "../../wailsjs/go/backend/Backend";
    import { onMount } from "svelte";
    import { openDialog, closeDialog, checkCommand } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let dialog = { showDialog: false, dialogTitle: '', dialogMessage: '' };

    let options = {
        sourceDir: "",
        destDir: "",
        exclude: [],
        compressData: false,
        linksOption: "",
        verify: false,
        schedule: "",
        compressFile: false,
    };

    onMount(async () => {
        await checkCommand("rsync", dialog);
    });

    const backup = async () => {
        try {
            await Backup(options);
            dialog = openDialog(dialog, "Success", `Backup created successfully`);
        } catch (err) {
            dialog = openDialog(dialog, "Error", `Error creating backup: ${err.message}`);
        }
    };
</script>

<CustomDialog
  bind:show={dialog.showDialog}
  title={dialog.dialogTitle}
  message={dialog.dialogMessage}
  onClose={() => dialog = closeDialog(dialog)}
  confirmButton={false}
/>

<div class="grid-container">
    <div class="grid-item">
        <label for="sourceDir">Source Directory:</label>
    </div>
    <div class="grid-item">
        <input type="text" bind:value={options.sourceDir} />
    </div>

    <div class="grid-item">
        <label for="destDir">Destination Directory:</label>
    </div>
    <div class="grid-item">
        <input type="text" bind:value={options.destDir} />
    </div>

    <div class="grid-item">
        <label for="exclude">Exclude (comma-separated paths):</label>
    </div>
    <div class="grid-item">
        <input type="text" bind:value={options.exclude} />
    </div>

    <div class="grid-item">
        <label for="compressData">Compress Data:</label>
    </div>
    <div class="grid-item">
        <input type="checkbox" bind:checked={options.compressData} />
    </div>

    <div class="grid-item">
        <label for="linksOption">Symbolic Links Option:</label>
    </div>
    <div class="grid-item">
        <select bind:value={options.linksOption}>
            <option value="">Default</option>
            <option value="-l">-l (copy links as links)</option>
            <option value="-L"
                >-L (transform symlink into referent file/dir)</option
            >
            <option value="-k">-k (copy symlinks as regular files)</option>
        </select>
    </div>

    <div class="grid-item">
        <label for="verify">Verify Backup:</label>
    </div>
    <div class="grid-item">
        <input type="checkbox" bind:checked={options.verify} />
    </div>

    <div class="grid-item">
        <label for="schedule">Schedule Backup (in seconds):</label>
    </div>
    <div class="grid-item">
        <input type="text" bind:value={options.schedule} />
    </div>

    <div class="grid-item">
        <label for="compressFile">Create a compress file:</label>
    </div>
    <div class="grid-item">
        <input type="checkbox" bind:checked={options.compressFile} />
    </div>

    <div class="grid-item">
        <button on:click={backup}>Backup</button>
    </div>
</div>

<style>
    .grid-container {
        display: grid;
        gap: 1rem;
        padding: 1rem;
        border: 1px solid #ccc;
        border-radius: 0.5rem;
        grid-template-columns: max-content 1fr;
    }

    .grid-item {
        display: grid;
        grid-template-columns: max-content;
        align-items: center;
    }

    label {
        font-weight: bold;
    }

    input[type="checkbox"] {
        margin: 0;
    }

    input[type="text"],
    select {
        padding: 0.75rem;
        border: 1px solid #ccc;
        border-radius: 0.25rem;
        color: white;
        width: 100%;
        box-sizing: border-box;
    }

    button {
        padding: 0.75rem;
        border: none;
        background-color: #4caf50;
        color: white;
        border-radius: 0.25rem;
        cursor: pointer;
        width: 100%;
        margin: 0 auto;
        grid-column: span 2;
    }

    button:hover {
        background-color: #3e8e41;
    }
</style>
