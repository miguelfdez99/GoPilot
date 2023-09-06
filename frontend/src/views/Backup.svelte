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
    import openIcon from "../assets/images/open.png";
    import infoIcon from "../assets/images/info.png";
    import { settings } from '../stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    let showInfoButton: boolean;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
        showInfoButton = $settings.showInfoButton;
    });

    $: {
    document.documentElement.style.setProperty('--main-font-size', fontSize);
    document.documentElement.style.setProperty('--main-color', color);
    document.documentElement.style.setProperty('--main-font-family', fontFamily);
    document.documentElement.style.setProperty('--main-bg-color', backgroundColor);
    document.documentElement.style.setProperty('--main-bg-color2', backgroundColor2);
    document.documentElement.style.setProperty('--main-input-color', inputColor);
    document.documentElement.style.setProperty('--main-button-color', buttonColor);
  }


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

    const scheduleBackup = async () => {
        try {
            await ScheduleBackup(options, options.schedule);
            await OpenDialogInfo("Backup scheduled successfully");
        } catch (err) {
            await OpenDialogError(`Error scheduling backup: ${err.message}`);
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

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
        <div style="color: ${color}; font-size: ${fontSize};">
            Backups are created using the rsync command.
            <br />
            <br />
            <b>Source Directory</b> is the directory to be backed up.
            <br />
            <br />
            <b>Destination Directory</b> is the directory where the backup will be stored.
            <br />
            <br />
            <b>Exclude</b> is a list of files and directories to be excluded from the backup.
            <br />
            <br />
            <b>Compress Data</b> compresses the data before transferring it.
            <br />
            <br />
            <b>Links Option</b> is the option to preserve hard links.
            <br />
            <br />
            <b>Verify</b> verifies the transfer.
            <br />
            <br />
            <b>Compress File</b> compresses the file data during the transfer.
            <br />
            <br />
            <b>Schedule</b> is the schedule for the backup. It uses the cron format.
            <br />
            <br />
            More information about <a href="https://linux.die.net/man/1/rsync" style="color: inherit;">rsync</a>
        </div>
        `
        );
    }

    function onDialogClose() {
        dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    }
</script>

<CustomDialog
    bind:show={dialog.showDialog}
    title="Info"
    message={dialog.dialogMessage}
    onClose={onDialogClose}
/>

<div class="main-container">
    {#if showInfoButton}
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    {/if}
    <h2>Backups</h2>
    <div class="grid-container">
        <div class="grid-item">
            <label for="sourceDir">Source Directory:</label>
        </div>
        <div class="grid-item">
            <input type="text" bind:value={options.sourceDir} />
            <button
                class="open-btn"
                title="Select Source Directory"
                on:click={selectDir}
                ><img src={openIcon} alt="Open Dir" class="open-icon" /></button
            >
        </div>

        <div class="grid-item">
            <label for="destDir">Destination Directory:</label>
        </div>
        <div class="grid-item">
            <input type="text" bind:value={options.destDir} />
            <button
                class="open-btn"
                title="Select Destination Directory"
                on:click={selectDestDir}
                ><img src={openIcon} alt="Open Dir" class="open-icon" /></button
            >
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
            <input
                type="text"
                bind:value={options.schedule}
                placeholder="Cron schedule (e.g., * * * * *)"
            />
        </div>

        <div class="grid-item">
            <button on:click={backup}>Backup Now</button>
        </div>
        <div class="grid-item">
            <button on:click={scheduleBackup}>Schedule Backup</button>
        </div>
    </div>
</div>

<style>
    .main-container {
        background-color: var(--main-bg-color);
        font-size: var(--main-font-size);
        color: var(--main-color);
    }
    .info-button {
        position: absolute;
        top: 10px;
        right: 10px;
        border: none;
        background: var(--main-button-color);
        height: 40px;
        width: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 50%;
    }

    .info-icon {
        max-width: none;
    }

    .grid-container {
        display: grid;
        gap: 1.5rem;
        padding: 2rem;
        background: #282828;
        box-shadow: 0 0 15px rgba(0, 0, 0, 0.3);
        border-radius: 5px;
        grid-template-columns: 1fr 1fr;
        color: #ddd;
        width: 90%;
        max-width: 800px;
        margin: 2rem auto;
        background-color: var(--main-bg-color2);
    }

    .grid-item {
        display: flex;
        align-items: center;
        color: var(--main-color);
    }

    label {
        font-weight: 500;
        margin-bottom: 0.5em;
    }

    input[type="checkbox"] {
        margin: 0;
        appearance: auto;
        background: var(--main-input-color);
    }

    input[type="text"],
    select {
        padding: 0.7em;
        border: 0;
        border-radius: 4px;
        background: var(--main-input-color);
        color: #fff;
        width: 100%;
        box-sizing: border-box;
        color: var(--main-color);
    }

    button {
        padding: 0.8em 1em;
        border: none;
        border-radius: 4px;
        background: var(--main-button-color);
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 100%;
        margin: 0 auto;
        grid-column: span 2;
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

    h2 {
        color: var(--main-color);
    }
</style>
