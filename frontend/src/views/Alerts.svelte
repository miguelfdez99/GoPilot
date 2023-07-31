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
            <p>This component is used for file, directory, and system stats monitoring.</p>
            <p>- Enter a file or directory path and click "Start Monitoring" to start the monitoring process. The file or directory will be monitored for changes and you will be alerted when a change occurs.</p>
            <p>- For system stats, enter the threshold values for CPU, RAM, and Disk Space. When these thresholds are exceeded, you will be alerted.</p>
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
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>

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
                    placeholder="Enter RAM threshold (in GB)"
                />
                <input
                    type="number"
                    bind:value={diskThreshold}
                    placeholder="Enter Disk Space threshold (in GB)"
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
    .info-button {
        position: absolute;
        top: 10px;
        right: 10px;
        border: none;
        background: #5a5858;
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
        background-color: #282828;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
        border-radius: 10px;
        color: #ddd;
        margin: 1rem;
    }

    h1,
    h2 {
        color: #ddd;
        margin-bottom: 1em;
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
        background: #333;
        color: #fff;
        width: 100%;
        box-sizing: border-box;
        margin-bottom: 0.5em;
    }

    button {
        padding: 0.8em 1em;
        border: none;
        border-radius: 4px;
        background: #1abc9c;
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 100%;
        margin-bottom: 1em;
    }

    button:hover {
        background-color: #16a085;
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
        border: 1px solid #333;
        border-radius: 4px;
        margin-bottom: 1em;
    }

    li {
        padding: 0.8em;
        border-bottom: 1px solid #444;
        background: #333;
    }

    li:last-child {
        border-bottom: none;
    }

    li strong {
        color: #1abc9c;
    }
</style>
