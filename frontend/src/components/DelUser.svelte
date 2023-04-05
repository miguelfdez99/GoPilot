<script lang="ts">
    import { DeleteUser } from "../../wailsjs/go/main/App.js";

    export let username = "";
    export let removeHomeDir = false;
    export let forceDelete = false;

    async function handleSubmit() {
        try {
            await DeleteUser(username, removeHomeDir, forceDelete);
            alert("User deleted successfully!");
        } catch (err) {
            alert(err.message);
        }
    }
</script>

<template>
    <form on:submit|preventDefault={handleSubmit}>
        <h2>Delete User</h2>
        <label>
            Username:
            <input type="text" bind:value={username} />
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

    h2 {
        margin-top: 0;
        margin-bottom: 20px;
    }

    label {
        display: flex;
        align-items: center;
        margin-bottom: 10px;
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
