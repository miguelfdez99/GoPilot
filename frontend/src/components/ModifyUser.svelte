<script lang="ts">
    import {
        ModifyUser,
        CommandExists,
    } from "../../wailsjs/go/backend/Backend.js";
    import { onMount } from "svelte";

    let username: string = "";
    let password: string = "";
    let uid: string = "";
    let gid: string = "";
    let home: string = "";
    let shell: string = "/bin/bash";
    let expire: string = "";
    let addGroup: string = "";


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
        await checkCommand("usermod");
    });

    async function modifyUser() {
        // Check if required fields are empty
        if (!username) {
            alert("Please fill out all fields");
            return;
        }

        // Create the user object
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
            // Call the Modify User function in the Go environment
            await ModifyUser(user.username, user);
            alert("User modified successfully!");
            // Clear the form
            username = "";
            password = "";
            uid = "";
            gid = "";
            home = "";
            shell = "";
            addGroup = "";
        } catch (err) {
            console.error(err);
            alert("Failed to modify user");
        }
    }
</script>

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
