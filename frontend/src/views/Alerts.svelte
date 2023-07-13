<script lang="ts">
    import {
        MonitorFile,
        MonitorDir,
        WatchList,
        MonitorSystemStats,
    } from "../../wailsjs/go/backend/Backend";
    import { onMount } from "svelte";

    let filepath: string = "";
    let dirpath: string = "";
    let watchList = {};
    let cpuThreshold: number;
    let ramThreshold: number;
    let diskThreshold: number;

    onMount(async () => {
        watchList = await WatchList();
    });

    const startMonitoringFile = () => {
        MonitorFile(filepath);
        fetchWatchList();
    };

    const startMonitoringDir = () => {
        MonitorDir(dirpath);
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
        watchList = await WatchList();
    }
</script>

<div>
    <h1>File System Monitor</h1>
    <div>
        <h2>File Monitor</h2>
        <input
            type="text"
            bind:value={filepath}
            placeholder="Enter a file path"
        />
        <button on:click={startMonitoringFile}>Start Monitoring File</button>
    </div>
    <div>
        <h2>Directory Monitor</h2>
        <input
            type="text"
            bind:value={dirpath}
            placeholder="Enter a directory path"
        />
        <button on:click={startMonitoringDir}>Start Monitoring Directory</button
        >
    </div>
    <h2>Watch List</h2>
    <ul>
        {#each Object.entries(watchList) as [path, type]}
            <li>
                <strong>{type}</strong>: {path}
            </li>
        {/each}
    </ul>

    <div>
        <h2>System Stats Monitor</h2>
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
        <button on:click={startMonitoringSystemStats}
            >Start Monitoring System Stats</button
        >
    </div>
</div>

<style>
    div {
        padding: 20px;
        background-color: white;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
        border-radius: 10px;
    }
    h1,
    h2 {
        color: #333;
    }
    input {
        margin-bottom: 10px;
        padding: 10px;
        width: 100%;
        box-sizing: border-box;
        border: 1px solid #ddd;
        border-radius: 5px;
    }
    button {
        padding: 10px 20px;
        border: none;
        background-color: #007bff;
        color: white;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }
    button:hover {
        background-color: #0056b3;
    }
    ul {
        padding: 0;
        list-style: none;
    }
    li {
        padding: 10px;
        border-bottom: 1px solid #ddd;
    }
    li:last-child {
        border-bottom: none;
    }
    li strong {
        color: #007bff;
    }
</style>
