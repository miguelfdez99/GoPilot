<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import {
        GetProcessInfo,
        TerminateProcess,
    } from "../../wailsjs/go/backend/Backend";
    import CustomDialog from "../dialogs/CustomDialog.svelte";
    import { openDialog, closeDialog } from "../functions/functions";

    let processInfo: string[] = [];
    let filteredProcessInfo = [];

    let sortBy: string = "null";
    let sortDirection: string = "null";
    let searchTerm: string = "";
    let interval: any;

    let iconAsc = "↑";
    let iconDesc = "↓";

    let processList = [];
    let showKillDialog: boolean = false;
    let dialogPid: number = 0;
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

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

    function confirmKillDialog(pid: number) {
        dialogPid = pid;
        showKillDialog = true;
    }

    function onDialogConfirm() {
        killProcess(dialogPid);
        showKillDialog = false;
    }

    function onDialogClose() {
        showKillDialog = false;
    }

    function killProcess(pid: number) {
        TerminateProcess(pid)
            .then(() => {
                processList = processList.filter((proc) => proc.Pid !== pid);
                dialog = openDialog(
                    dialog,
                    "Success",
                    `Successfully killed process ${pid}`
                );
            })
            .catch((err) => {
                dialog = openDialog(
                    dialog,
                    "Error",
                    `Failed to kill process ${pid}: ${err}`
                );
            });
    }

    onMount(() => {
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
        }, 1000);
    });

    onDestroy(() => {
        clearInterval(interval);
    });

    $: if (searchTerm !== "") {
        filterProcesses();
    }
</script>

<CustomDialog
    bind:show={showKillDialog}
    title="Kill Process"
    message={`Are you sure you want to kill process ${dialogPid}?`}
    onConfirm={onDialogConfirm}
    onClose={onDialogClose}
/>

<div class="search-container">
    <input type="text" bind:value={searchTerm} placeholder="Search..." />
</div>

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
                    <button on:click={() => confirmKillDialog(proc.Pid)}
                        >Terminate</button
                    >
                </td>
            </tr>
        {/each}
    </tbody>
</table>

<style>
    .command {
        max-width: 400px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: normal;
        word-wrap: break-word;
    }
    .search-container {
        padding-bottom: 1rem;
    }

    input[type="text"] {
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 0.25rem;
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
        background-color: #009879;
        color: #ffffff;
        text-align: left;
    }
    th,
    td {
        padding: 12px 15px;
        width: 10%;
        overflow: hidden;
        text-overflow: ellipsis;
    }
    tbody tr {
        border-bottom: 1px solid #dddddd;
    }
    tbody tr:nth-of-type(even) {
        background-color: #312c2c;
    }
    tbody tr:last-of-type {
        border-bottom: 2px solid #009879;
    }
</style>
