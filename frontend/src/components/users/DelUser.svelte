<script lang="ts">
    import { DeleteUser } from '../../../wailsjs/go/backend/Backend';
    import { onMount } from 'svelte';
    import {
        openDialog,
        closeDialog,
        checkCommand,
    } from "../../functions/functions";
    import CustomDialog from "../dialogs/CustomDialog.svelte";

    let username: string = "";
    let removeHomeDir: boolean = false;
    let forceDelete: boolean = false;

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    onMount(async () => {
        await checkCommand("userdel", dialog);
    });

    async function handleSubmit() {
        try {
            await DeleteUser(username, removeHomeDir, forceDelete);
            dialog = openDialog(dialog, "Success", `Successfully deleted user ${username}`);
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


<template>
    <form on:submit|preventDefault={handleSubmit}>
        <label>
            <span>Username:</span>
            <input type="text" bind:value={username} required />
        </label>
        <label>
            <input type="checkbox" bind:checked={removeHomeDir} />
            <span>Remove home directory</span>
        </label>
        <label>
            <input type="checkbox" bind:checked={forceDelete} />
            <span>Force delete</span>
        </label>
        <button type="submit">Delete User</button>
    </form>
</template>

<style>
    form {
        display: flex;
        flex-direction: column;
        margin: 20px;
        border: 1px solid #ccc;
        padding: 20px;
        border-radius: 5px;
        color: black;
        max-width: 400px;
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

    input[type="checkbox"] {
        margin-right: 10px;
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
        margin-top: 20px;
    }

    button[type="submit"]:hover {
        background-color: #0069d9;
    }
</style>
