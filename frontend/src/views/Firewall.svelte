<script lang="ts">
    import { onMount } from "svelte";
    import {
        ListFirewallRules,
        AddFirewallRule,
        RemoveFirewallRule,
        GetFirewallStatus, 
        SetFirewallStatus, 
        GetDefaultPolicy, 
        SetDefaultPolicy 
    } from '../../wailsjs/go/backend/Backend';

    let rules = [];
    let error = null;

    let port = "";
    let protocol = "";

    let firewallEnabled: boolean;
    let incomingPolicy: string;
    let outgoingPolicy: string;

    async function addRule() {
        try {
            await AddFirewallRule(port, protocol);
            listFirewallRules();
            port = "";
            protocol = "";
        } catch (err) {
            error = err;
        }
    }

    onMount(async () => {
        try {
            let status = await GetFirewallStatus();
            firewallEnabled = status === "active";
            incomingPolicy = await GetDefaultPolicy("incoming");
            outgoingPolicy = await GetDefaultPolicy("outgoing");
            listFirewallRules();
        } catch (err) {
            console.error(err);
        }
    });

    // Whenever firewallEnabled changes, this reactive statement will execute
    $: setFirewallStatus(firewallEnabled), firewallEnabled;

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
            rules = await ListFirewallRules();
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

    <div>
        Status: 
        <label class="switch">
            <input type="checkbox" bind:checked={firewallEnabled}>
          <span class="slider round"></span>
        </label>
    </div>

    <div>
        Incoming:
        <select bind:value={incomingPolicy} on:change={() => setDefaultPolicy('incoming', incomingPolicy)}>
            <option value="allow">Allow</option>
            <option value="deny">Deny</option>
        </select>
    </div>
    
    <div>
        Outgoing: 
        <select bind:value={outgoingPolicy} on:change={() => setDefaultPolicy('outgoing', outgoingPolicy)}>
            <option value="allow">Allow</option>
            <option value="deny">Deny</option>
        </select>
    </div>

    {#if error}
        <p class="error">{error.message}</p>
    {:else if rules.length > 0}
        <textarea readonly>{rules.join("\n")}</textarea>
    {:else}
        <p>No firewall rules found.</p>
    {/if}

    <form on:submit|preventDefault={addRule}>
        <h2>Add Firewall Rule</h2>
        <label>
            Port:
            <input type="text" bind:value={port} required />
        </label>
        <label>
            Protocol:
            <input type="text" bind:value={protocol} required />
        </label>
        <button type="submit">Add Rule</button>
    </form>
    <button on:click={listFirewallRules}>Refresh</button>
</main>

<style>
    textarea {
        width: 100%;
        height: 200px;
        font-family: monospace;
        white-space: pre-wrap;
        overflow-y: scroll;
        color: white;
    }

    .error {
        color: red;
        font-weight: bold;
    }

    h1, h2, label {
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
