<script lang="ts">
    import { onMount } from "svelte";
    import {
        GetLSCPU,
        GetSystemInfo,
    } from "../../../wailsjs/go/backend/Backend";

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
        return JSON.parse(response);
    }

    function parseSystemInfo(response: string) {
        return JSON.parse(response);
    }
</script>

<div class="summary">
    <header class="header">
        <h2>{systemInfo.osName}</h2>
    </header>

    <div class="info-grid">
        <div class="info-card">
            <h3>System Information</h3>
            <table>
                <tr><th>OS Name</th><td>{systemInfo.osName}</td></tr>
                <tr><th>Kernel Version</th><td>{systemInfo.kernelVer}</td></tr>
                <tr><th>Uptime</th><td>{systemInfo.uptime}</td></tr>
                <tr><th>Hostname</th><td>{systemInfo.hostname}</td></tr>
                <tr><th>DE</th><td>{systemInfo.desktopEnv}</td></tr>
                <tr><th>User</th><td>{systemInfo.currentUsername}</td></tr>
                <tr><th>Memory</th><td>{systemInfo.memory}</td></tr>
            </table>
        </div>

        <div class="info-card">
            <h3>CPU Information</h3>
            <table>
                <tr><th>Architecture</th><td>{cpuInfo.architecture}</td></tr>
                <tr><th>CPUs</th><td>{cpuInfo.cpus}</td></tr>
                <tr><th>Model Name</th><td>{cpuInfo.modelName}</td></tr>
                <tr
                    ><th>Threads Per Core</th><td>{cpuInfo.threadPerCore}</td
                    ></tr
                >
                <tr
                    ><th>Cores Per Socket</th><td>{cpuInfo.corePerSocket}</td
                    ></tr
                >
                <tr><th>Socket</th><td>{cpuInfo.socket}</td></tr>
                <tr><th>CPU Modes</th><td>{cpuInfo.cpuModes}</td></tr>
            </table>
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
        --sidebar-width: 250px;
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

    .summary {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-lg);
        padding: var(--spacing-lg);
    }

    .header {
        display: flex;
        align-items: center;
        gap: var(--spacing-md);
        padding: var(--spacing-lg);
        background-color: var(--color-bg-secondary);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .info-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
        gap: var(--spacing-lg);
    }

    .info-card {
        background-color: var(--color-bg-secondary);
        border-radius: 0.5rem;
        padding: var(--spacing-lg);
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    h3 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    table {
        width: 100%;
        border-collapse: collapse;
    }

    th,
    td {
        padding: var(--spacing-md);
        text-align: left;
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    th {
        color: var(--color-text-secondary);
        font-weight: 600;
        width: 40%;
    }

    td {
        color: var(--color-text-primary);
    }
</style>
