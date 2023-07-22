<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import { GetAllServices, EnableService, DisableService, StartService, StopService } from "../../wailsjs/go/backend/Backend";

    let serviceStore = writable([]);
    let error: Error | null = null;

    onMount(async () => {
        try {
            const services = await GetAllServices();
            serviceStore.set(services.map(service => {
                return {
                    service,
                    startupStatus: service.StartupStatus === "enabled",
                    startupStatusChanged: false,
                    runningStatus: service.RunningStatus === "enabled",
                    runningStatusChanged: false
                };
            }));
        } catch (err) {
            console.error(err);
            error = err;
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

            if (service.runningStatusChanged) {
                if (service.runningStatus) {
                    await StartService(service.service.Name);
                } else {
                    await StopService(service.service.Name);
                }
                service.runningStatusChanged = false;
            }
        }
    });
</script>

<div>
    <h1>Services</h1>

    <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Startup Status</th>
                <th>Running Status</th>
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
                            <span class="slider round"></span>
                        </label>
                    </td>
                    <td>
                        <label class="switch">
                            <input
                                type="checkbox"
                                bind:checked={serviceObject.runningStatus}
                                on:change={() =>
                                    (serviceObject.runningStatusChanged = true)}
                            />
                            <span class="slider round"></span>
                        </label>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>

    {#if error}
        <p class="error">{error.message}</p>
    {/if}
</div>

<style>
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
        transition: .4s;
    }
    .slider:before {
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: .4s;
    }
    input:checked + .slider {
        background-color: #2196F3;
    }
    input:focus + .slider {
        box-shadow: 0 0 1px #2196F3;
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
</style>
