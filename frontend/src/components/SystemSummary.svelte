<script lang="ts">
    import { onMount } from "svelte";
    import {
        GetLSCPU,
        GetSystemInfo,
    } from "../../wailsjs/go/backend/Backend";

    let cpuInfo = {
        architecture: "",
        cpus: "",
        modelName: "",
        threadPerCore: "",
        corePerSocket: "",
        socket: "",
        cpuModes: "",
    };

    let systemInfo = {
        osName: "",
        kernelVer: "",
        uptime: "",
        hostname: "",
        desktopEnv: "",
        currentUsername: "",
        memory: "",
    };

    onMount(async () => {
        const lscpuResponse = await GetLSCPU();
        cpuInfo = parseResponse(lscpuResponse);
        const systemInfoResponse = await GetSystemInfo();
        systemInfo = parseSystemInfo(systemInfoResponse);
    });

    function parseResponse(response: string) {
        const data = JSON.parse(response);
        return {
            architecture: data.architecture,
            cpus: data.cpus,
            modelName: data.modelName,
            threadPerCore: data.threadPerCore,
            corePerSocket: data.corePerSocket,
            socket: data.socket,
            cpuModes: data.cpuModes,
        };
    }

    function parseSystemInfo(response: string) {
        const data = JSON.parse(response);
        return {
            osName: data.osName,
            kernelVer: data.kernelVer,
            uptime: data.uptime,
            hostname: data.hostname,
            desktopEnv: data.desktopEnv,
            currentUsername: data.currentUsername,
            memory: data.memory,
        };
    }
</script>

<div class="summary-container">
    <h2>System Summary</h2>
    <div>
        <h3>System Info</h3>
        <table>
            <tr><th>OS Name</th><td>{systemInfo.osName}</td></tr>
            <tr><th>Kernel Version</th><td>{systemInfo.kernelVer}</td></tr>
            <tr><th>Uptime</th><td>{systemInfo.uptime}</td></tr>
            <tr><th>Hostname</th><td>{systemInfo.hostname}</td></tr>
            <tr><th>Desktop Environment</th><td>{systemInfo.desktopEnv}</td></tr>
            <tr><th>User</th><td>{systemInfo.currentUsername}</td></tr>
            <tr><th>Memory</th><td>{systemInfo.memory}</td></tr>
        </table>
    </div>
    <div>
        <h3>CPU Info</h3>
        <table>
            <tr><th>Architecture</th><td>{cpuInfo.architecture}</td></tr>
            <tr><th>CPUs</th><td>{cpuInfo.cpus}</td></tr>
            <tr><th>Model Name</th><td>{cpuInfo.modelName}</td></tr>
            <tr><th>Threads Per Core</th><td>{cpuInfo.threadPerCore}</td></tr>
            <tr><th>Cores Per Socket</th><td>{cpuInfo.corePerSocket}</td></tr>
            <tr><th>Socket</th><td>{cpuInfo.socket}</td></tr>
            <tr><th>CPU Modes</th><td>{cpuInfo.cpuModes}</td></tr>
        </table>
    </div>
</div>

<style>
    :root {
        --main-color: #414a4c;
        --secondary-color: #1966da;
        --font-family: "Roboto", sans-serif;
    }

    .summary-container {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 2em;
        font-family: var(--font-family);
        color: var(--main-color);
        margin: 2em;
    }

    h2 {
        grid-column: span 2;
        border-bottom: 2px solid var(--main-color);
        padding-bottom: 0.5em;
        color: white;
    }

    h3 {
        color: white;
    }

    table {
        width: 100%;
        border-collapse: collapse;
    }

    th,
    td {
        padding: 0.5em;
        text-align: left;
        border-bottom: 1px solid var(--main-color);
        color: white;
    }

    th {
        background-color: var(--secondary-color);
    }
</style>
