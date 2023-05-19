<script lang="ts">
    import { DeleteUser, CommandExists } from '../../wailsjs/go/backend/Backend.js';
    import { onMount } from 'svelte';

    export let username: string = "";
    export let removeHomeDir: boolean = false;
    export let forceDelete: boolean = false;



    async function checkCommand(command: string) {
        try {
            const commandExists = await CommandExists(command);
            if (!commandExists) {
                alert(
                    `System command '${command}' required for this operation is not installed.` +
                        ` Please install it and try again.`
                );
            }
        } catch (err) {
            console.error(err);
            alert(
                `Failed to check if system command '${command}' is installed: ${err.message}`
            );
        }
    }

    onMount(async () => {
        await checkCommand("userdel");
    });

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
