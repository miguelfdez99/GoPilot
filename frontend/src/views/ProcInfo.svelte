<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import {
        GetProcessInfo,
        TerminateProcess,
        OpenDialogInfo,
        OpenDialogError,
        OpenDialogQuestion,
    } from "../../wailsjs/go/backend/Backend";
    import terminateIcon from "../assets/images/terminate.png";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
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

    let processInfo: string[] = [];
    let filteredProcessInfo = [];
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    let sortBy: string = "null";
    let sortDirection: string = "null";
    let searchTerm: string = "";
    let interval: any;

    let iconAsc = "↑";
    let iconDesc = "↓";

    let processList = [];

    let loading = writable(false);

    let currentIcon = {
        User: "",
        Pid: "",
        CpuPercent: "",
        MemPercent: "",
        VMS: "",
        RSS: "",
        TTY: "",
        Status: "",
        StartTime: "",
        Cmdline: "",
    };

    const sortTable = (key) => {
        if (sortBy === key) {
            sortDirection = sortDirection === "asc" ? "desc" : "asc";
        } else {
            sortBy = key;
            sortDirection = "asc";
        }

        for (let column in currentIcon) {
            currentIcon[column] = "";
        }

        currentIcon[key] = sortDirection === "asc" ? iconAsc : iconDesc;

        if (sortBy) {
            processInfo = sortProcessInfo(processInfo, sortBy, sortDirection);
            filterProcesses();
        }
    };

    const sortProcessInfo = (
        processInfo: string[],
        sortBy: string,
        sortDirection: string
    ) => {
        return [...processInfo].sort((a, b) => {
            if (a[sortBy] < b[sortBy]) return sortDirection === "asc" ? -1 : 1;
            if (a[sortBy] > b[sortBy]) return sortDirection === "asc" ? 1 : -1;
            return 0;
        });
    };

    const filterProcesses = () => {
        filteredProcessInfo = processInfo.filter((process) =>
            Object.values(process).some((val) =>
                val.toString().toLowerCase().includes(searchTerm.toLowerCase())
            )
        );
    };

    async function killProcess(pid: number) {
        try {
            const res = await OpenDialogQuestion(`Are you sure you want to kill process ${pid}?`);
            if (res === "Yes") {
                await TerminateProcess(pid);
                processList = processList.filter((proc) => proc.Pid !== pid);
                await OpenDialogInfo(`Successfully killed process ${pid}`);
            } else {
                await OpenDialogInfo(`Canceled killing process ${pid}`);
            }
        } catch (err) {
            await OpenDialogError(`Failed to kill process ${pid}: ${err}`);
        }
    }

    onMount(async () => {
        loading.set(true);
        try {
            interval = setInterval(async () => {
                processInfo = await GetProcessInfo();
                if (sortBy) {
                    processInfo = sortProcessInfo(
                        processInfo,
                        sortBy,
                        sortDirection
                    );
                }
                filterProcesses();
                loading.set(false);
            }, 1000);
        } catch (err) {
            await OpenDialogError(`Failed to get process info: ${err}`);
            loading.set(false);
        }
    });


    onDestroy(() => {
        clearInterval(interval);
    });

    $: if (searchTerm !== "") {
        filterProcesses();
    }

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
            <p style="color: ${color};>
                This view shows the list of processes running on the system.
            </p>
            <p style="color: ${color};>
                The process list is updated every second.
                The process list is sorted by the PID in ascending order by default.
                Click on a column header to sort the process list by that column.
            </p>
            <p style="color: ${color};>
                You can terminate a process by clicking on the terminate icon in the last column.
            </p>
            <p style="color: ${color};>
                <b>USER</b> - The user that started the process.
            </p>
            <p style="color: ${color};>
                <b>PID</b> - The process ID.
            </p>
            <p style="color: ${color};>
                <b>%CPU</b> - The percentage of the CPU time used by the process since the last update.
            </p>
            <p style="color: ${color};>
                <b>%MEM</b> - The percentage of the total RAM used by the process.
            </p>
            <p style="color: ${color};>
                <b>VSZ</b> - The total amount of virtual memory used by the process.
            </p>
            <p style="color: ${color};>
                <b>RSS</b> - The total amount of physical memory used by the process.
            </p>
            <p style="color: ${color};>
                <b>TTY</b> - The controlling terminal for the process.
            </p>
            <p style="color: ${color};>
                <b>STAT</b> - The state of the process.
            </p>
            <p style="color: ${color};>
                <b>START</b> - The time the process started.
            </p>
            <p style="color: ${color};>
                <b>COMMAND</b> - The command that started the process.
            </p>
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
<div class="search-container">
    <input type="text" bind:value={searchTerm} placeholder="Search..." />
</div>

{#if $loading}
        <div class="loading">
            <p>Loading...</p>
            <div class="spinner" />
        </div>
    {:else}

<table>
    <thead>
        <tr>
            <th on:click={() => sortTable("User")}>USER {currentIcon.User}</th>
            <th on:click={() => sortTable("Pid")}>PID {currentIcon.Pid}</th>
            <th on:click={() => sortTable("CpuPercent")}
                >%CPU {currentIcon.CpuPercent}</th
            >
            <th on:click={() => sortTable("MemPercent")}
                >%MEM {currentIcon.MemPercent}</th
            >
            <th on:click={() => sortTable("VMS")}>VSZ {currentIcon.VMS}</th>
            <th on:click={() => sortTable("RSS")}>RSS {currentIcon.RSS}</th>
            <th on:click={() => sortTable("TTY")}>TTY {currentIcon.TTY}</th>
            <th on:click={() => sortTable("Status")}
                >STAT {currentIcon.Status}</th
            >
            <th on:click={() => sortTable("StartTime")}
                >START {currentIcon.StartTime}</th
            >
            <th on:click={() => sortTable("Cmdline")}
                >COMMAND {currentIcon.Cmdline}</th
            >
        </tr>
    </thead>
    <tbody>
        {#each filteredProcessInfo as proc (proc.Pid)}
            <tr>
                <td>{proc.User}</td>
                <td>{proc.Pid}</td>
                <td>{proc.CpuPercent.toFixed(2)}</td>
                <td>{proc.MemPercent.toFixed(2)}</td>
                <td>{proc.VMS}</td>
                <td>{proc.RSS}</td>
                <td>{proc.TTY}</td>
                <td>{proc.Status}</td>
                <td>{new Date(proc.StartTime).toLocaleString()}</td>
                <td class="command">{proc.Cmdline}</td>
                <td>
                    <button 
                    class="terminate-btn"
                    on:click={() => killProcess(proc.Pid)}>
                        <img src={terminateIcon} alt="Terminate" class="terminate-icon" />
                    </button>
                </td>
            </tr>
        {/each}
    </tbody>
</table>
{/if}
</div>

<style>
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
    
    .command {
        max-width: 400px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: normal;
        word-wrap: break-word;
    }
    .search-container {
        padding-bottom: 1rem;
        max-width: 98%;
    }

    input[type="text"] {
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 0.25rem;
        background: var(--main-input-color);
        color: var(--main-color);
    }
    table {
        table-layout: auto;
        width: 100%;
        border-collapse: collapse;
        margin: 1rem 0;
        font-size: 0.9em;
        font-family: sans-serif;
        min-width: 400px;
        box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
    }
    thead tr {
        background-color: var(--main-input-color);
        color: var(--main-color);
        text-align: left;
    }
    th,
    td {
        padding: 12px 15px;
        width: 10%;
        overflow: hidden;
        text-overflow: ellipsis;
        color: var(--main-color);
    }
    tbody tr {
        border-bottom: 1px solid #dddddd;
    }
    tbody tr:nth-of-type(even) {
        background-color:  var(--main-input-color);
    }
    tbody tr:last-of-type {
        border-bottom: 2px solid var(--main-input-color);
    }

    .terminate-btn {
    width: 24px;
    height: 24px;
    padding: 0;
    margin: 0;
    background: var(--main-button-color);
    display: flex; 
    justify-content: center;
    align-items: center;
    border: none;
}

.terminate-icon {
    width: 16px;
    height: 16px;
}

.loading {
        text-align: center;
    }

    .loading p {
        color: var(--main-color);;
        font-size: 16px;
        margin-bottom: 20px;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--main-color);;
        border-radius: 50%;
        width: 30px;
        height: 30px;
        animation: spin 1s linear infinite;
        margin: 0 auto;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
