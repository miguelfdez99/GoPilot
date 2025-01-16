<script lang="ts">
    import {
        ModifyUser,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let username: string = "";
    let password: string = "";
    let uid: string = "";
    let gid: string = "";
    let home: string = "";
    let shell: string = "/bin/bash";
    let addGroup: string = "";

    function dismiss() {
        dispatch("dismiss");
    }

    async function modifyUser() {
        if (!username) {
            await OpenDialogError("Please enter a username");
            return;
        }

        const user = {
            username,
            password,
            uid: parseInt(uid),
            gid: parseInt(gid),
            homeDirectory: home,
            groups: addGroup ? [addGroup] : [],
            shell,
        };

        try {
            await ModifyUser(user.username, user);
            await OpenDialogInfo(`Successfully modified user ${username}`);
            username = "";
            password = "";
            uid = "";
            gid = "";
            home = "";
            shell = "";
            addGroup = "";
        } catch (err) {
            await OpenDialogError(`Failed to modify user ${username}: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Modify User</h2>
    <form on:submit|preventDefault={modifyUser} class="form-control">
        <label class="input-field">
            <span>Username:</span>
            <input type="text" bind:value={username} required />
        </label>
        <label class="input-field">
            <span>Password:</span>
            <input type="password" bind:value={password} />
        </label>
        <label class="input-field">
            <span>User ID (UID):</span>
            <input type="text" bind:value={uid} />
        </label>
        <label class="input-field">
            <span>Group ID (GID):</span>
            <input type="text" bind:value={gid} />
        </label>
        <label class="input-field">
            <span>Home directory:</span>
            <input type="text" bind:value={home} />
        </label>
        <label class="input-field">
            <span>Add to group:</span>
            <input type="text" bind:value={addGroup} />
        </label>
        <label class="input-field">
            <span>Login shell:</span>
            <select bind:value={shell}>
                <option value="/bin/bash">/bin/bash</option>
                <option value="/bin/sh">/bin/sh</option>
                <option value="/usr/bin/zsh">/usr/bin/zsh</option>
            </select>
        </label>
        <button type="submit" class="submit-button">
            <span class="material-icons">edit</span>
            Modify User
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

    input[type="text"],
    input[type="password"],
    select {
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
