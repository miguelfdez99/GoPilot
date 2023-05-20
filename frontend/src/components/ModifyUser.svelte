<script lang="ts">
    import {
        ModifyUser,
    } from "../../wailsjs/go/backend/Backend";
    import { onMount } from "svelte";
    import {
        openDialog,
        closeDialog,
        checkCommand,
    } from "../functions/functions";
    import CustomDialog from "../dialogs/CustomDialog.svelte";

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    let username: string = "";
    let password: string = "";
    let uid: string = "";
    let gid: string = "";
    let home: string = "";
    let shell: string = "/bin/bash";
    let addGroup: string = "";

    onMount(async () => {
        await checkCommand("usermod", dialog);
    });

    async function modifyUser() {
        if (!username) {
            dialog = openDialog(dialog, "Error", "Username is required");
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
            dialog = openDialog(dialog, "Success", `Successfully modified user ${username}`);
            username = "";
            password = "";
            uid = "";
            gid = "";
            home = "";
            shell = "";
            addGroup = "";
        } catch (err) {
            dialog = openDialog(dialog, "Error", `${err}`);
        }
    }
</script>

<CustomDialog
    bind:show={dialog.showDialog}
    title={dialog.dialogTitle}
    message={dialog.dialogMessage}
    onClose={() => (dialog = closeDialog(dialog))}
    confirmButton={false}
/>

<form on:submit|preventDefault={modifyUser}>
    <label>
        <span>Username:</span>
        <input type="text" bind:value={username} required />
    </label>
    <label>
        <span>Password:</span>
        <input type="password" bind:value={password} />
    </label>
    <label>
        <span>User ID (UID):</span>
        <input type="text" bind:value={uid} />
    </label>
    <label>
        <span>Group ID (GID):</span>
        <input type="text" bind:value={gid} />
    </label>
    <label>
        <span>Home directory:</span>
        <input type="text" bind:value={home} />
    </label>
    <label>
        <span>Add to group:</span>
        <input type="text" bind:value={addGroup} />
    </label>
    <label>
        <span>Login shell:</span>
        <select bind:value={shell}>
            <option value="/bin/bash">/bin/bash</option>
            <option value="/bin/sh">/bin/sh</option>
            <option value="/usr/bin/zsh">/usr/bin/zsh</option>
        </select>
    </label>
    <button type="submit">Modify user</button>
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        margin: 20px;
        border: 1px solid #ccc;
        padding: 20px;
        border-radius: 5px;
        color: black;
    }

    label {
        display: flex;
        flex-direction: column;
        margin-bottom: 10px;
    }

    label span {
        font-size: 16px;
        font-weight: bold;
        margin-bottom: 5px;
    }

    input[type="text"],
    select {
        font-size: 16px;
        padding: 10px;
        border-radius: 5px;
        border: 1px solid #ccc;
        margin-bottom: 10px;
    }

    input[type="password"],
    select {
        font-size: 16px;
        padding: 10px;
        border-radius: 5px;
        border: 1px solid #ccc;
        margin-bottom: 10px;
    }

    button[type="submit"] {
        background-color: #007bff;
        color: #fff;
        font-size: 16px;
        padding: 10px;
        border-radius: 5px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    button[type="submit"]:hover {
        background-color: #0069d9;
    }
</style>
