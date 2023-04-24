<script lang="ts">
    import { ModifyGroup } from '../../wailsjs/go/backend/Backend.js';

    let name = "";
    let gid = "";

    async function modifyGroup() {
        if (!name || !gid) {
            alert("Please fill out all fields");
            return;
        }
        try {
            await ModifyGroup(name, parseInt(gid));
            alert("Group modified successfully!");
            name = "";
            gid = "";
        } catch (err) {
            console.error(err);
            alert("Failed to modify group");
        }
    }
</script>

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
