<script lang="ts">
    import {
        CreateGroup,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let name: string = "";
    let gid: string = "";

    function dismiss() {
        dispatch("dismiss");
    }

    async function createGroup() {
        if (!name) {
            await OpenDialogError("Please enter a group name");
            return;
        }

        const group = {
            name,
            gid: gid ? parseInt(gid) : undefined,
        };

        try {
            await CreateGroup(group.name, group.gid);
            await OpenDialogInfo(`Successfully created group ${name}`);
            name = "";
            gid = "";
        } catch (err) {
            await OpenDialogError(`Failed to create group ${name}: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Add Group</h2>
    <form on:submit|preventDefault={createGroup} class="form-control">
        <label class="input-field">
            <span>Group name:</span>
            <input type="text" bind:value={name} required />
        </label>
        <label class="input-field">
            <span>Group ID (GID):</span>
            <input type="text" bind:value={gid} />
        </label>
        <button type="submit" class="submit-button">
            <span class="material-icons">group_add</span>
            Create Group
        </button>
        <button class="back-button" on:click={dismiss}>
            <span class="material-icons">arrow_back</span>
            Back
        </button>
    </form>
</div>

<style>
    .container {
        background-color: var(--color-bg-secondary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        max-width: 400px;
        margin: 0 auto;
    }

    h2 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
        text-align: center;
    }

    .form-control {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-md);
    }

    .input-field {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-sm);
    }

    .input-field span {
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    input[type="text"] {
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
    }

    .submit-button,
    .back-button {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        padding: var(--spacing-sm) var(--spacing-md);
        background-color: var(--color-bg-tertiary);
        border: 1px solid var(--color-bg-tertiary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .submit-button:hover,
    .back-button:hover {
        background-color: var(--color-bg-secondary);
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-accent-blue);
    }
</style>
