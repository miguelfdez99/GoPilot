<script lang="ts">
    import { onMount } from "svelte";
    import {
        GetLSCPU,
        GetSystemInfo,
    } from "../../../wailsjs/go/backend/Backend";
    import debianIcon from "../../assets/images/debian.png";
    import archIcon from "../../assets/images/arch.png";
    import fedoraIcon from "../../assets/images/fedora.png";
    import ubuntuIcon from "../../assets/images/ubuntu.png";
    import manjaroIcon from "../../assets/images/manjaro.png";
    import mintIcon from "../../assets/images/mint.png";
    import elementaryIcon from "../../assets/images/elementary.png";
    import { settings } from '../../stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
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

    let distroIcons = {
    "Debian": debianIcon,
    "Arch": archIcon,
    "Fedora": fedoraIcon,
    "Ubuntu": ubuntuIcon,
    "Manjaro": manjaroIcon,
    "Mint": mintIcon,
    "Elementary": elementaryIcon,
};

let distroIcon: string = "";
let osName: string = "";

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
        osName = systemInfo.osName;
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

    function findDistroIcon(osName: string): string {
    for (let distro in distroIcons) {
        if (osName.toLowerCase().includes(distro.toLowerCase())) {
            return distroIcons[distro];
        }
    }
    return "";
}

    $: distroIcon = findDistroIcon(osName);
</script>

<div class="summary-container">
    <div class="header-container">
        {#if distroIcon}
            <div class="os-info">
                <img class="os-icon" src={distroIcon} alt="Distro icon" />
                <h2>{osName}</h2>
            </div>
        {/if}
    </div>
    <div class="info-tables">
        <h3>System Info</h3>
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
        --font-family: "Roboto", sans-serif;
    }

    .summary-container {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 2em;
        font-family: var(--font-family);
        color: var(--main-color);
        margin: 2em;
    }

    .header-container {
        grid-column: 1 / -1;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
    }

    .os-info {
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .os-icon {
        width: 100px;
        margin-bottom: 10px;
        margin-right: 20px;
    }

    h2 {
        margin-top: 30px;
    }

    table {
        width: 100%;
        border-collapse: collapse;
    }

    h2, h3 {
        color: white;
    }

    th,
    td {
        padding: 0.5em;
        text-align: left;
        border-bottom: 1px solid var(--main-color);
        color: white;
    }

    th {
        background-color: var(--main-input-color);
    }
</style>
