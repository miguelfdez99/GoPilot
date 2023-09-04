<script lang="ts">
    import {
        ModifyUser,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    let username: string = "";
    let password: string = "";
    let uid: string = "";
    let gid: string = "";
    let home: string = "";
    let shell: string = "/bin/bash";
    let addGroup: string = "";

    function dismiss() {
        dispatch('dismiss');
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
        <button type="submit" class="submit-button">Modify user</button>
        <button class="back-button" on:click={dismiss}>Back</button>
    </form>
</div>

<style>
    h2 {
      text-align: center;
      color: var(--main-color);
    }
    
    .container {
        position: relative;
        padding: 20px;
        max-width: 400px;
        margin: 0 auto;
    }

    .back-button {
        right: 0;
        padding: 10px 20px;
        border: none;
        background: var(--main-button-color);
        border-radius: 5px;
        font-size: 16px;
        color: #fff;
        transition: background-color 0.3s;
    }

    .form-control {
        padding: 30px;
        background-color: var(--main-bg-color2);
        border-radius: 5px;
    }

    .input-field {
        margin-bottom: 20px;
    }

    .input-field span {
        display: block;
        font-size: 14px;
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
        margin-bottom: 5px;
    }

    input[type="text"], input[type="password"], select {
        width: 100%;
        border: none;
        padding: 10px;
        border-radius: 5px;
        font-size: 16px;
        background: var(--main-input-color);
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
    }

    .submit-button {
        display: block;
        width: 100%;
        padding: 10px;
        border: none;
        background: var(--main-button-color);
        color: white;
        border-radius: 5px;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.3s;
    }

</style>
