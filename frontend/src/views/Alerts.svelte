<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { writable } from 'svelte/store';
    import {
        MonitorFile,
        MonitorDir,
        WatchList,
        MonitorSystemStats,
    } from "../../wailsjs/go/backend/Backend";

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
        const transformedWatchList: Array<WatchItem> = Object.entries(newWatchList).map(([path, type]) => ({ path, type }));
        watchListStore.set(transformedWatchList);
    }

</script>


<div class="monitor-container">

    <div class="monitor-section">
        <h1>File System Monitor</h1>

        <h2>File Monitor</h2>
        <div class="input-group">
            <input type="text" bind:value={filepath} placeholder="Enter a file path" />
            <button on:click={startMonitoringFile}>Start Monitoring File</button>
        </div>

        <h2>Directory Monitor</h2>
        <div class="input-group">
            <input type="text" bind:value={dirpath} placeholder="Enter a directory path" />
            <button on:click={startMonitoringDir}>Start Monitoring Directory</button>
        </div>
    </div>

    <div class="monitor-section">
        <h2>System Stats Monitor</h2>
        <div class="stats-group">
            <input type="number" bind:value={cpuThreshold} placeholder="Enter CPU threshold (in %)" />
            <input type="number" bind:value={ramThreshold} placeholder="Enter RAM threshold (in GB)" />
            <input type="number" bind:value={diskThreshold} placeholder="Enter Disk Space threshold (in GB)" />
        </div>
        <button on:click={startMonitoringSystemStats}>Start Monitoring System Stats</button>
    
        <h2>Watch List</h2>
        <ul>
            {#each $watchListStore as { path, type }}
                <li><strong>{type}</strong>: {path}</li>
            {/each}
        </ul>
        
        
    </div>

</div>

<style>
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
    
    h1, h2 {
        color: #ddd;
        margin-bottom: 1em;
    }

    .input-group {
        display: flex;
        flex-direction: column;
        margin-bottom: 1em;
    }

    input {
        padding: .7em;
        border: 0;
        border-radius: 4px;
        background: #333;
        color: #fff;
        width: 100%;
        box-sizing: border-box;
        margin-bottom: 0.5em;
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
        margin-bottom: 1em;
    }

    button:hover {
        background-color: #16a085;
    }

    ul {
        padding: 0;
        list-style: none;
        border: 1px solid #333;
        border-radius: 4px;
        margin-bottom: 1em;
    }

    li {
        padding: .8em;
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
