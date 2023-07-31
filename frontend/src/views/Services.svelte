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
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
    import infoIcon from "../assets/images/info.png";

    let serviceStore = writable([]);
    let runningServiceStore = writable([]);
    let error: Error | null = null;
    let selectedTable = "all";
    let loading = writable(false);
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    onMount(async () => {
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

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
            <p>
                This page allows you to manage the services on your system.
                You can enable/disable services to start on boot and start/stop
                services.
            </p>`
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

<div>
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    <h1>Services</h1>

    <select bind:value={selectedTable}>
        <option value="all">Startup Services</option>
        <option value="running">Running Services</option>
    </select>

    {#if $loading}
        <div class="loading">
            <p>Loading...</p>
            <div class="spinner" />
        </div>
    {:else if selectedTable === "all"}
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

    {#if error}
        <p class="error">{error.message}</p>
    {/if}
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
    div {
        padding-bottom: 1rem;
    }
    table {
        width: 100%;
        border-collapse: collapse;
        margin: 1rem 0;
        font-size: 0.9em;
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
        overflow: hidden;
        text-overflow: ellipsis;
        color: white;
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
        background-color: #ccc;
        transition: 0.4s;
    }
    .slider:before {
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: 0.4s;
    }
    input:checked + .slider {
        background-color: #2196f3;
    }
    input:focus + .slider {
        box-shadow: 0 0 1px #2196f3;
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

    h1,
    option {
        color: white;
    }

    .loading {
        text-align: center;
    }

    .loading p {
        color: white;
        font-size: 16px;
        margin-bottom: 20px;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid #fff;
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
