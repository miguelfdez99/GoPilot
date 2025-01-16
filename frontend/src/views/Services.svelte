<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import {
        GetAllServices,
        GetRunningServices,
        EnableService,
        DisableService,
        StartService,
        StopService,
        CheckAdmin,
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let serviceStore = writable([]);
    let runningServiceStore = writable([]);
    let error: Error | null = null;
    let selectedTable = "all";
    let loading = writable(false);
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    onMount(async () => {
        await checkPrivileges();
        loading.set(true);
        try {
            const services = await GetAllServices();
            serviceStore.set(
                services.map((service) => ({
                    service,
                    startupStatus: service.StartupStatus === "enabled",
                    startupStatusChanged: false,
                }))
            );

            const runningServices = await GetRunningServices();
            runningServiceStore.set(
                runningServices.map((service) => ({
                    service,
                    activeState: service.ActiveState === "active",
                    activeStateChanged: false,
                }))
            );
        } catch (err) {
            console.error(err);
            error = err;
        } finally {
            loading.set(false);
        }
    });

    serviceStore.subscribe(async (services) => {
        for (let service of services) {
            if (service.startupStatusChanged) {
                if (service.startupStatus) {
                    await EnableService(service.service.Name);
                } else {
                    await DisableService(service.service.Name);
                }
                service.startupStatusChanged = false;
            }
        }
    });

    runningServiceStore.subscribe(async (services) => {
        for (let service of services) {
            if (service.activeStateChanged) {
                if (service.activeState) {
                    await StartService(service.service.Name);
                } else {
                    await StopService(service.service.Name);
                }
                service.activeStateChanged = false;
            }
        }
    });

    function onDialogClose() {
        dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    }

    async function checkPrivileges() {
        try {
            const admin = await CheckAdmin();
            if (!admin) {
                dialog = openDialog(
                    dialog,
                    "Warning",
                    `
                    <p>
                        You are not running this application as root. You will not be able to manage services.
                    </p>
                    `
                );
            }
        } catch (error) {
            console.error("Error checking admin privileges:", error);
        }
    }
</script>

<div class="services-container">
    <h1>Services</h1>

    <!-- Table Selector -->
    <select bind:value={selectedTable}>
        <option value="all">Startup Services</option>
        <option value="running">Running Services</option>
    </select>

    <!-- Loading State -->
    {#if $loading}
        <div class="loading">
            <p>Loading...</p>
            <div class="spinner" />
        </div>
    {:else if selectedTable === "all"}
        <!-- Startup Services Table -->
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Startup Status</th>
                </tr>
            </thead>
            <tbody>
                {#each $serviceStore as serviceObject (serviceObject.service.Name)}
                    <tr>
                        <td>{serviceObject.service.Name}</td>
                        <td>
                            <label class="switch">
                                <input
                                    type="checkbox"
                                    bind:checked={serviceObject.startupStatus}
                                    on:change={() =>
                                        (serviceObject.startupStatusChanged = true)}
                                />
                                <span class="slider round" />
                            </label>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {:else}
        <!-- Running Services Table -->
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Active State</th>
                    <th>Description</th>
                </tr>
            </thead>
            <tbody>
                {#each $runningServiceStore as serviceObject (serviceObject.service.Name)}
                    <tr>
                        <td>{serviceObject.service.Name}</td>
                        <td>
                            <label class="switch">
                                <input
                                    type="checkbox"
                                    bind:checked={serviceObject.activeState}
                                    on:change={() =>
                                        (serviceObject.activeStateChanged = true)}
                                />
                                <span class="slider round" />
                            </label>
                        </td>
                        <td>{serviceObject.service.Description}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}

    <!-- Error Message -->
    {#if error}
        <p class="error">{error.message}</p>
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

    .services-container {
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        font-family: sans-serif;
    }

    h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    select {
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
        margin-bottom: var(--spacing-md);
    }

    table {
        width: 100%;
        border-collapse: collapse;
        margin-bottom: var(--spacing-md);
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

    .switch {
        position: relative;
        display: inline-block;
        width: 60px;
        height: 34px;
    }

    .switch input {
        opacity: 0;
        width: 0;
        height: 0;
    }

    .slider {
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: var(--color-bg-tertiary);
        transition: 0.4s;
    }

    .slider:before {
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: var(--color-text-primary);
        transition: 0.4s;
    }

    input:checked + .slider {
        background-color: var(--color-accent-blue);
    }

    input:focus + .slider {
        box-shadow: 0 0 1px var(--color-accent-blue);
    }

    input:checked + .slider:before {
        transform: translateX(26px);
    }

    .slider.round {
        border-radius: 34px;
    }

    .slider.round:before {
        border-radius: 50%;
    }

    .loading {
        text-align: center;
        padding: var(--spacing-lg);
    }

    .loading p {
        color: var(--color-text-primary);
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

    .error {
        color: var(--color-accent-red);
        text-align: center;
        margin-top: var(--spacing-md);
    }
</style>
