<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import {
        MonitorFile,
        MonitorDir,
        WatchList,
        MonitorSystemStats,
        OpenFile,
        OpenDir,
    } from "../../wailsjs/go/backend/Backend";
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

    interface WatchItem {
        path: string;
        type: string;
    }

    let filepath: string = "";
    let dirpath: string = "";
    let watchListStore = writable<Array<WatchItem>>([]);
    let cpuThreshold: number;
    let ramThreshold: number;
    let diskThreshold: number;
    let pollingInterval;
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    onMount(async () => {
        fetchWatchList();
        pollingInterval = setInterval(fetchWatchList, 1000);
    });

    onDestroy(() => {
        clearInterval(pollingInterval);
    });

    const startMonitoringFile = async () => {
        await MonitorFile(filepath);
        fetchWatchList();
    };

    const startMonitoringDir = async () => {
        await MonitorDir(dirpath);
        fetchWatchList();
    };

    const startMonitoringSystemStats = () => {
        let thresholds = {
            CPU: cpuThreshold,
            Memory: ramThreshold,
            Disk: diskThreshold,
        };
        MonitorSystemStats(thresholds);
    };

    async function fetchWatchList() {
        const newWatchList: { [key: string]: string } = await WatchList();
        const transformedWatchList: Array<WatchItem> = Object.entries(
            newWatchList
        ).map(([path, type]) => ({ path, type }));
        watchListStore.set(transformedWatchList);
    }

    const selectFile = async () => {
        const file = await OpenFile();
        filepath = file;
    };

    const selectDir = async () => {
        const dir = await OpenDir();
        dirpath = dir;
    };

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
            <p style="color: ${color}; font-size: ${fontSize};">This component is used for file, directory, and system stats monitoring.</p>
            <p style="color: ${color}; font-size: ${fontSize};">- Enter a file or directory path and click "Start Monitoring" to start the monitoring process. The file or directory will be monitored for changes and you will be alerted when a change occurs.</p>
            <p style="color: ${color}; font-size: ${fontSize};">- For system stats, enter the threshold values for CPU, RAM, and Disk Space. When these thresholds are exceeded, you will be alerted.</p>
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
    <h2>Alerts</h2>
    {#if showInfoButton}
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    {/if}

    <div class="monitor-container">
        <div class="monitor-section">
            <h1>File System Monitor</h1>

            <h2>File Monitor</h2>
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
            <button on:click={startMonitoringFile}>Start Monitoring File</button
            >

            <h2>Directory Monitor</h2>
            <div class="input-group">
                <input
                    type="text"
                    bind:value={dirpath}
                    placeholder="Enter a directory path"
                />
                <button
                    class="open-btn"
                    title="Choose Directory"
                    on:click={selectDir}
                    ><img
                        src={openIcon}
                        alt="Open Dir"
                        class="open-icon"
                    /></button
                >
            </div>
            <button on:click={startMonitoringDir}
                >Start Monitoring Directory</button
            >
        </div>

        <div class="monitor-section">
            <h2>System Stats Monitor</h2>
            <div class="stats-group">
                <input
                    type="number"
                    bind:value={cpuThreshold}
                    placeholder="Enter CPU threshold (in %)"
                />
                <input
                    type="number"
                    bind:value={ramThreshold}
                    placeholder="Enter RAM threshold (in %)"
                />
                <input
                    type="number"
                    bind:value={diskThreshold}
                    placeholder="Enter Disk Space threshold (in %)"
                />
            </div>            
            <button on:click={startMonitoringSystemStats}
                >Start Monitoring System Stats</button
            >

            <h2>Watch List</h2>
            <ul>
                {#each $watchListStore as { path, type }}
                    <li><strong>{type}</strong>: {path}</li>
                {/each}
            </ul>
        </div>
    </div>
</div>

<style>
    .main-container {
        background-color: var(--main-bg-color);
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

    .monitor-container {
        display: flex;
        justify-content: space-between;
        max-width: 95%;
        margin: 2rem auto;
    }

    .monitor-section {
        flex: 1;
        padding: 1rem;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
        border-radius: 10px;
        margin: 1rem;
        background-color: var(--main-bg-color2);
    }

    h1,
    h2 {
        color: var(--main-color);
        margin-bottom: 1em;
        font-family: var(--main-font-family);
    }

    .input-group {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 1em;
    }

    input {
        padding: 0.7em;
        border: 0;
        border-radius: 4px;
        background: var(--main-input-color);
        color: var(--main-color);
        width: 100%;
        box-sizing: border-box;
        margin-bottom: 0.5em;
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

    ul {
        padding: 0;
        list-style: none;
        border: 1px;
        border-radius: 4px;
        margin-bottom: 1em;
    }

    li {
        padding: 0.8em;
        background-color: var(--main-bg-color);
        color: var(--main-color);
        list-style: none;
    }

    li strong {
        color: var(--main-button-color);
    }
</style>
