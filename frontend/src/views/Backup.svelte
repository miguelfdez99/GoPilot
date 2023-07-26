<script lang="ts">
    import { Backup, ScheduleBackup } from "../../wailsjs/go/backend/Backend";
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
        compressFile: false,
        schedule: "",
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

    const scheduleBackup = async () => {
        try {
            await ScheduleBackup(options, options.schedule);
            dialog = openDialog(dialog, "Success", `Backup scheduled successfully`);
        } catch (err) {
            dialog = openDialog(dialog, "Error", `Error scheduling backup: ${err.message}`);
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

<h1>Backups</h1>
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
        <label for="compressFile">Create a compress file:</label>
    </div>
    <div class="grid-item">
        <input type="checkbox" bind:checked={options.compressFile} />
    </div>

    <div class="grid-item">
        <label for="schedule">Schedule:</label>
    </div>
    <div class="grid-item">
        <input type="text" bind:value={options.schedule} placeholder="Cron schedule (e.g., * * * * *)"/>
    </div>


    <div class="grid-item">
        <button on:click={backup}>Backup Now</button>
    </div>
    <div class="grid-item">
        <button on:click={scheduleBackup}>Schedule Backup</button>
    </div>
</div>

<style>
    .grid-container {
        display: grid;
        gap: 1.5rem;
        padding: 2rem;
        background: #282828;
        box-shadow: 0 0 15px rgba(0,0,0,0.3);
        border-radius: 5px;
        grid-template-columns: 1fr 1fr;
        color: #ddd;
        width: 90%;
        max-width: 800px;
        margin: 2rem auto;
    }

    .grid-item {
        display: flex;
        flex-direction: column;
        align-items: start;
    }

    label {
        font-weight: 500;
        margin-bottom: 0.5em;
    }

    input[type="checkbox"] {
        margin: 0;
    }

    input[type="text"],
    select {
        padding: .7em;
        border: 0;
        border-radius: 4px;
        background: #333;
        color: #fff;
        width: 100%;
        box-sizing: border-box;
    }

    button {
        padding: .8em 1em;
        border: none;
        border-radius: 4px;
        background: #1abc9c;
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 100%;
        margin: 0 auto;
        grid-column: span 2;
    }

    button:hover {
        background-color: #16a085;
    }

    h1 {
        color: white;
    }
</style>

