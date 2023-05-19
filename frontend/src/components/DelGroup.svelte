<script lang="ts">
    import { DeleteGroup } from '../../wailsjs/go/backend/Backend';
    import { checkCommand } from "../functions/functions";
    import { onMount } from "svelte";

    let name: string = "";

    onMount(async () => {
        await checkCommand("groupdel");
    });

    async function deleteGroup() {
        // Check if required fields are empty
        if (!name) {
            alert("Please fill out all fields");
            return;
        }

        // Call the CreateGroup function in the Go environment
        try {
            await DeleteGroup(name);
            alert("Group deleted successfully!");
            // Clear the form
            name = "";
        } catch (err) {
            console.error(err);
            alert("Failed to delete group");
        }
    }
</script>

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
