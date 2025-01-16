<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import {
        GetProcessInfo,
        TerminateProcess,
        OpenDialogInfo,
        OpenDialogError,
        OpenDialogQuestion,
        CheckAdmin,
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

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
        sortDirection: string,
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
                val.toString().toLowerCase().includes(searchTerm.toLowerCase()),
            ),
        );
    };

    async function killProcess(pid: number) {
        try {
            const res = await OpenDialogQuestion(
                `Are you sure you want to kill process ${pid}?`,
            );
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
        await checkPrivileges();
        loading.set(true);
        try {
            interval = setInterval(async () => {
                processInfo = await GetProcessInfo();
                if (sortBy) {
                    processInfo = sortProcessInfo(
                        processInfo,
                        sortBy,
                        sortDirection,
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

    async function checkPrivileges() {
        try {
            const admin = await CheckAdmin();
            if (!admin) {
                dialog = openDialog(
                    dialog,
                    "Info",
                    `
                    <p style="color: var(--color-text-primary); font-size: 1rem;">
                        You are not running this application as root. You can see processes but you will not be able to kill them.
                    </p>
                    `,
                );
            }
        } catch (error) {
            console.error("Error checking admin privileges:", error);
        }
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
                    <th on:click={() => sortTable("User")}
                        >USER {currentIcon.User}</th
                    >
                    <th on:click={() => sortTable("Pid")}
                        >PID {currentIcon.Pid}</th
                    >
                    <th on:click={() => sortTable("CpuPercent")}
                        >%CPU {currentIcon.CpuPercent}</th
                    >
                    <th on:click={() => sortTable("MemPercent")}
                        >%MEM {currentIcon.MemPercent}</th
                    >
                    <th on:click={() => sortTable("VMS")}
                        >VSZ {currentIcon.VMS}</th
                    >
                    <th on:click={() => sortTable("RSS")}
                        >RSS {currentIcon.RSS}</th
                    >
                    <th on:click={() => sortTable("TTY")}
                        >TTY {currentIcon.TTY}</th
                    >
                    <th on:click={() => sortTable("Status")}
                        >STAT {currentIcon.Status}</th
                    >
                    <th on:click={() => sortTable("StartTime")}
                        >START {currentIcon.StartTime}</th
                    >
                    <th on:click={() => sortTable("Cmdline")}
                        >COMMAND {currentIcon.Cmdline}</th
                    >
                    <th>ACTION</th>
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
                                on:click={() => killProcess(proc.Pid)}
                            >
                                <span class="material-icons">delete</span>
                            </button>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}
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
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

    .main-container {
        padding: var(--spacing-lg);
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .search-container {
        padding-bottom: var(--spacing-md);
    }

    input[type="text"] {
        width: 100%;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        margin: var(--spacing-md) 0;
        font-size: 0.9rem;
    }

    thead tr {
        background-color: var(--color-bg-tertiary);
        color: var(--color-text-primary);
        text-align: left;
    }

    th,
    td {
        padding: var(--spacing-sm);
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    tbody tr:hover {
        background-color: var(--color-bg-secondary);
    }

    .command {
        max-width: 400px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .terminate-btn {
        background: none;
        border: none;
        padding: 0;
        margin: 0;
        cursor: pointer;
        color: var(--color-accent-red);
    }

    .material-icons {
        font-size: 1.5rem;
    }

    .loading {
        text-align: center;
    }

    .loading p {
        color: var(--color-text-primary);
        font-size: 1rem;
        margin-bottom: var(--spacing-md);
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--color-accent-blue);
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
