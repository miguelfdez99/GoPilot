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
        CheckAdmin
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
    import infoIcon from "../assets/images/info.png";
    import { settings } from "../stores";

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
        document.documentElement.style.setProperty(
            "--main-font-size",
            fontSize
        );
        document.documentElement.style.setProperty("--main-color", color);
        document.documentElement.style.setProperty(
            "--main-font-family",
            fontFamily
        );
        document.documentElement.style.setProperty(
            "--main-bg-color",
            backgroundColor
        );
        document.documentElement.style.setProperty(
            "--main-bg-color2",
            backgroundColor2
        );
        document.documentElement.style.setProperty(
            "--main-input-color",
            inputColor
        );
        document.documentElement.style.setProperty(
            "--main-button-color",
            buttonColor
        );
    }

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

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
        <p style="color: ${color}; font-size: ${fontSize};">
            This page allows you to manage the services on your system.
            You can enable/disable services to start on boot and start/stop
            services.
        </p>
        `
        );
    }

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
        <p style="color: ${color}; font-size: ${fontSize};">
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

<CustomDialog
    bind:show={dialog.showDialog}
    title="Info"
    message={dialog.dialogMessage}
    onClose={onDialogClose}
/>

<div>
    {#if showInfoButton}
        <button
            type="button"
            class="info-button"
            title="Info"
            on:click={openInfo}
        >
            <img src={infoIcon} alt="Open Info" class="info-icon" />
        </button>
    {/if}
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
        background-color: var(--main-input-color);
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
        text-align: left;
    }
    th,
    td {
        padding: 12px 15px;
        overflow: hidden;
        text-overflow: ellipsis;
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
    }
    tbody tr {
        border-bottom: 1px solid #dddddd;
    }
    tbody tr:nth-of-type(even) {
        background-color: var(--main-input-color2);
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
        color: var(--main-color);
        font-family: var(--main-font-family);
    }

    select {
        background: var(--main-input-color);
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
    }

    .loading {
        text-align: center;
    }

    .loading p {
        color: var(--main-color);
        font-size: 16px;
        margin-bottom: 20px;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--main-color);
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
