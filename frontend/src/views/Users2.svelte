<script lang="ts">
  import { onMount } from "svelte";
  import { createEventDispatcher } from "svelte";
  import { CreateUser2 } from "../../wailsjs/go/main/App.js";

  const dispatcher = createEventDispatcher();
  let currentView = "users2";

  let username = '';
  let password = '';
  let uid = '';
  let gid = '';
  let shell = "/bin/bash";

  async function createUser() {
  // Check if required fields are empty
  if (!username ) {
    alert('Please fill out all fields');
    return;
  }

  // Create the user object
  const user = {
    username,
    password,
    uid: parseInt(uid),
    gid: parseInt(gid),
    shell
  };

  try {
    // Call the Create User function in the Go environment
    await CreateUser2(user);
    alert('User created successfully!');
    // Clear the form
    username = '';
    password = '';
    uid = '';
    gid = '';
    shell = '';
  } catch (err) {
    console.error(err);
    alert('Failed to create user');
  }
}

onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
  });
</script>

<form on:submit|preventDefault={createUser}>
  <label>
    <span>Username:</span>
    <input type="text" bind:value={username} />
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
    <span>Login shell:</span>
    <select bind:value={shell}>
      <option value="/bin/bash">/bin/bash</option>
      <option value="/bin/sh">/bin/sh</option>
      <option value="/usr/bin/zsh">/usr/bin/zsh</option>
    </select>
  </label>
  <button type="submit">Create user</button>
</form>

<style>
  form {
    display: flex;
    flex-direction: column;
    margin: 20px;
    border: 1px solid #ccc;
    padding: 20px;
    border-radius: 5px;
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
