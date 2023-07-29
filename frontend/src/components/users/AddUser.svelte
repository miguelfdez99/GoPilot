<script lang="ts">
    import {
        CreateUser,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let username: string = "";
    let password: string = "";
    let uid: string = "";
    let gid: string = "";
    let shell: string = "/bin/bash";

    function dismiss() {
        dispatch("dismiss");
    }

    async function createUser() {
        if (!username) {
            await OpenDialogError("Please enter a username");
            return;
        }

        const user = {
            username,
            password,
            uid: parseInt(uid),
            gid: parseInt(gid),
            shell,
        };

        try {
            await CreateUser(user);
            await OpenDialogInfo(`Successfully created user ${username}`);
            username = "";
            password = "";
            uid = "";
            gid = "";
            shell = "";
        } catch (err) {
            await OpenDialogError(`Failed to create user ${username}: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Add User</h2>
    <form on:submit|preventDefault={createUser} class="form-control">
        <label class="input-field">
            <span>Username:</span>
            <input type="text" bind:value={username} />
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
            <span>Login shell:</span>
            <select bind:value={shell}>
                <option value="/bin/bash">/bin/bash</option>
                <option value="/bin/sh">/bin/sh</option>
                <option value="/usr/bin/zsh">/usr/bin/zsh</option>
            </select>
        </label>
        <button type="submit" class="submit-button">Create user</button>
        <button class="back-button" on:click={dismiss}>Back</button>
    </form>
  </div>
  
  <style>
    h2 {
      text-align: center;
      color: #fff;
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
      background-color: #333;
      border-radius: 5px;
      font-size: 16px;
      color: #fff;
      transition: background-color 0.3s;
    }
  
    .back-button:hover {
      background-color: #555;
    }
  
    .form-control {
      padding: 30px;
      background-color: #222;
      border-radius: 5px;
    }
  
    .input-field {
      margin-bottom: 20px;
    }
  
    .input-field span {
      display: block;
      font-size: 14px;
      color: #ccc;
      margin-bottom: 5px;
    }
  
    input[type="text"], input[type="password"], select {
      width: 100%;
      border: none;
      padding: 10px;
      border-radius: 5px;
      font-size: 16px;
      background-color: #333;
      color: #fff;
    }
  
    .submit-button {
      display: block;
      width: 100%;
      padding: 10px;
      border: none;
      background: #1abc9c;
      color: white;
      border-radius: 5px;
      font-size: 16px;
      cursor: pointer;
      transition: background-color 0.3s;
    }
  
    .submit-button:hover {
        background-color: #16a085;
    }
  </style>
