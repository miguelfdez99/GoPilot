<script lang="ts">
    import { DeleteGroup } from '../../../wailsjs/go/backend/Backend';
    import { onMount } from "svelte";
    import {
        openDialog,
        closeDialog,
        checkCommand,
    } from "../../functions/functions";
    import CustomDialog from "../dialogs/CustomDialog.svelte";

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    let name: string = "";

    onMount(async () => {
        await checkCommand("groupdel", dialog);
    });

    async function deleteGroup() {
        if (!name) {
            dialog = openDialog(dialog, "Error", "Group name is required");
            return;
        }

        try {
            await DeleteGroup(name);
            dialog = openDialog(dialog, "Success", `Successfully deleted group ${name}`);
            name = "";
        } catch (err) {
            dialog = openDialog(dialog, "Error", `Failed to delete group: ${name}`);
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
    <form on:submit|preventDefault={deleteGroup}>
      <label>
        <span>Group name:</span>
        <input type="text" bind:value={name} required />
      </label>
      <button type="submit">Delete Group</button>
    </form>
  </template>
<style>
    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 2rem;
        font-size: 1.2rem;
    }

    label {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-bottom: 1rem;
        color: black;
    }

    span {
        margin-bottom: 0.5rem;
    }

    input {
        padding: 0.5rem;
        font-size: 1rem;
        border-radius: 0.25rem;
        border: 1px solid #ccc;
    }

    button {
        padding: 0.5rem 1rem;
        font-size: 1rem;
        border-radius: 0.25rem;
        border: none;
        background-color: #007bff;
        color: #fff;
        cursor: pointer;
    }

    button:hover {
        background-color: #0069d9;
    }
</style>
