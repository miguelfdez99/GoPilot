<script lang="ts">
    import { ModifyGroup } from '../../wailsjs/go/backend/Backend';
    import { onMount } from "svelte";
    import {
        openDialog,
        closeDialog,
        checkCommand,
    } from "../functions/functions";
    import CustomDialog from "../dialogs/CustomDialog.svelte";

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    let name: string = "";
    let gid: string = "";

    onMount(async () => {
        await checkCommand("groupadd", dialog);
    });

    async function modifyGroup() {
        if (!name || !gid) {
            dialog = openDialog(dialog, "Error", "Group name and GID are required");
            return;
        }
        try {
            await ModifyGroup(name, parseInt(gid));
            dialog = openDialog(dialog, "Success", `Successfully modified group ${name}`);
            name = "";
            gid = "";
        } catch (err) {
            console.error(err);
            dialog = openDialog(dialog, "Error", `Failed to modified group: ${err}`);
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
    <form on:submit|preventDefault={modifyGroup}>
      <label>
        <span>Group name:</span>
        <input type="text" bind:value={name} required />
      </label>
      <label>
        <span>Group ID (GID):</span>
        <input type="text" bind:value={gid} />
      </label>
      <button type="submit">Modify Group</button>
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
