<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import {
        ListUfw,
        GetFirewallStatus,
        SetFirewallStatus,
        GetDefaultPolicy,
        SetDefaultPolicy,
        FetchTrafficData,
    } from "../../wailsjs/go/backend/Backend";

    export const firewallEnabled = writable(false);

    let rules = [];
    let error: Error | null = null;
    let incomingPolicy: string;
    let outgoingPolicy: string;
    let trafficData = undefined;
    let shouldSync = false;

    onMount(async () => {
        try {
            trafficData = await FetchTrafficData();
            let stringStatus = await GetFirewallStatus();
            firewallEnabled.set(stringStatus.toLowerCase().includes("active"));
            listFirewallRules();

            shouldSync = true;
        } catch (err) {
            console.error(err);
            error = err;
        }
    });

    firewallEnabled.subscribe((value) => {
        if (shouldSync) {
            setFirewallStatus(value);
        }
    });

    async function setFirewallStatus(status: boolean) {
        try {
            let stringStatus = status ? "active" : "inactive";
            await SetFirewallStatus(stringStatus);
        } catch (err) {
            console.error(err);
        }
    }

    async function listFirewallRules() {
        try {
            rules = await ListUfw();
            error = null;
        } catch (err) {
            rules = [];
            error = err;
        }
    }

    async function setDefaultPolicy(direction: string, policy: string) {
        try {
            await SetDefaultPolicy(direction, policy);
            if (direction === "incoming") {
                incomingPolicy = policy;
            } else {
                outgoingPolicy = policy;
            }
        } catch (err) {
            console.error(err);
        }
    }
</script>

<main>
    <h1>Firewall Settings</h1>

    <div class="settings">
        <div class="setting">
            <label
                >Status:
                <div class="switch">
                    <input type="checkbox" bind:checked={$firewallEnabled} />
                    <span class="slider round" />
                </div>
            </label>
        </div>

        <div class="setting">
            <label
                >Incoming:
                <select
                    bind:value={incomingPolicy}
                    on:change={() =>
                        setDefaultPolicy("incoming", incomingPolicy)}
                    disabled={!firewallEnabled}
                    class="select"
                >
                    <option value="allow">Allow</option>
                    <option value="deny">Deny</option>
                </select>
            </label>
        </div>

        <div class="setting">
            <label
                >Outgoing:
                <select
                    bind:value={outgoingPolicy}
                    on:change={() =>
                        setDefaultPolicy("outgoing", outgoingPolicy)}
                    disabled={!firewallEnabled}
                    class="select"
                >
                    <option value="allow">Allow</option>
                    <option value="deny">Deny</option>
                </select>
            </label>
        </div>
    </div>

    {#if error}
        <p class="error">{error.message}</p>
    {:else if rules.length > 0}
        <textarea readonly>{rules.join("\n")}</textarea>
    {:else}
        <p>No firewall rules found.</p>
    {/if}

    <!-- <form on:submit|preventDefault={addRule} class="form">
        <h2>Add Firewall Rule</h2>
        <label>
            Port:
            <input
                type="text"
                bind:value={port}
                required
                disabled={!firewallEnabled}
                class="input"
            />
        </label>
        <label>
            Protocol:
            <input
                type="text"
                bind:value={protocol}
                required
                disabled={!firewallEnabled}
                class="input"
            />
        </label>
        <button type="submit" disabled={!firewallEnabled} class="btn"
            >Add Rule</button
        >
    </form> -->

    {#if trafficData === undefined}
        <p>Loading data...</p>
    {:else if trafficData.length === 0}
        <p>No data to display</p>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Nº</th>
                    <th>Protocol</th>
                    <th>Local Address:Port</th>
                    <th>Peer Address:Port</th>
                    <th>Application</th>
                </tr>
            </thead>
            <tbody>
                {#each trafficData as data, i (i)}
                    <tr>
                        <td>{i + 1}</td>
                        <td>{data.netid}</td>
                        <td>{data.localAddr}</td>
                        <td>{data.peerAddr}</td>
                        <td>{data.application}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}

    {#if error}
        <p>Error: {error.message}</p>
    {/if}

    <button on:click={listFirewallRules} disabled={!firewallEnabled} class="btn"
        >Refresh</button
    >
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
        font-family: Arial, sans-serif;
        color: #333;
        padding: 20px;
        color: white;
    }

    h1,
    p,
    label,
    td,
    tr,
    th {
        color: white;
    }

    .settings {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 20px;
        margin-bottom: 20px;
    }

    .setting {
        flex: 1 1 200px;
    }

    .select {
        display: block;
        width: 100%;
        padding: 10px;
        font-size: 16px;
        border-radius: 5px;
        border: 1px solid #ccc;
        box-sizing: border-box;
        margin-top: 5px;
        color: white;
    }

    /* .form {
        width: 100%;
        max-width: 400px;
    } */

    .btn {
        display: inline-block;
        background-color: #007bff;
        color: white;
        padding: 10px 20px;
        margin-top: 10px;
        border-radius: 5px;
        text-decoration: none;
        font-size: 16px;
        line-height: 1;
        cursor: pointer;
        border: none;
    }

    .btn:disabled {
        background-color: #ccc;
        cursor: not-allowed;
        color: white;
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
</style>
