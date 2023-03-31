<script>
    import { createEventDispatcher } from 'svelte';
    import {CreateUser2} from '../../wailsjs/go/main/App.js'

    const dispatcher = createEventDispatcher();

    let username = '';
    let password = '';
    let uid = '';
    let gid = '';
    let homeDirectory = '';
    let shell = '';

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

  </script>

  <form on:submit|preventDefault={createUser}>
    <div>
      <label for="username">Username:</label>
      <input type="text" id="username" bind:value={username} required />
    </div>

    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" bind:value={password}/>
    </div>

    <div>
      <label for="uid">UID:</label>
      <input type="number" id="uid" bind:value={uid}/>
    </div>

    <div>
      <label for="gid">GID:</label>
      <input type="number" id="gid" bind:value={gid}/>
    </div>

    <div>
      <label for="homeDirectory">Home directory:</label>
      <input type="text" id="homeDirectory" bind:value={homeDirectory}/>
    </div>

    <div>
      <label for="shell">Shell:</label>
      <input type="text" id="shell" bind:value={shell}/>
    </div>

    <button type="submit">Create User</button>
  </form>


<style>
    form {
      display: flex;
      flex-direction: column;
      align-items: center;
      margin-top: 2rem;
      font-size: 1.2rem;
    }

    label {
      display: block;
      margin-bottom: 0.5rem;
    }

    input {
      padding: 0.5rem;
      margin-bottom: 1rem;
      border-radius: 0.25rem;
      border: none;
      box-shadow: 0px 0px 5px 0px rgba(0, 0, 0, 0.2);
    }

    button {
      padding: 0.5rem 1rem;
      background-color: #007bff;
      color: #fff;
      border: none;
      border-radius: 0.25rem;
      cursor: pointer;
      transition: background-color 0.3s ease-in-out;
    }

    button:hover {
      background-color: #0069d9;
    }
  </style>