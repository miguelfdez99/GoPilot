<script lang="ts">

  import { onMount } from 'svelte';

  import {PrintUsers, CreateUser, DeleteUser, CreateUserWithDir, DeleteUserWithDir, CheckAdmin} from '../../wailsjs/go/main/App.js'

  let currentView = 'users';
  let name: string
  let username: string
  let adminText: string

function printUser(): void {
      PrintUsers()
    }

    function createUser(): void {
      CreateUser()
    }

    function deleteUser(): void {
      DeleteUser()
    }

    function createUserWithDir(name: string): void {
      CreateUserWithDir(username)
    }

    function deleteUserWithDir(name: string): void {
      DeleteUserWithDir(username)
    }

    function checkAdmin(): void {
      CheckAdmin()
    }


  onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
    CheckAdmin().then(result => {
        if (result === false) {
          adminText = "You are not an admin"
        }
    });
  });
</script>

<main>
  <div class="input-box" id="input">
    <h1>Users</h1>
    <button class="btn" on:click={printUser}>Print Users</button>
    <button class="btn" on:click={createUser}>Create User</button>
    <button class="btn" on:click={deleteUser}>Delete User</button>
    <input autocomplete="off" bind:value={username} class="input" id="username" type="text"/>
    <button class="btn" on:click={() => createUserWithDir(username)}>Create User With Dir</button>
    <button class="btn" on:click={() => deleteUserWithDir(username)}>Delete User With Dir</button>

    <button class="btn" on:click={checkAdmin}>Check Admin</button>

    {#if adminText === "You are not an admin"}
      <h1>{adminText}</h1>
    {/if}

  </div>
</main>

<style>
main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.input-box {
  background-color: #f5f5f5;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h1 {
  font-size: 2rem;
  margin-bottom: 1.5rem;
  color: #333;
}

.btn {
  background-color: #333;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 0.8rem 1.2rem;
  margin-right: 1rem;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.3s ease-in-out;
}

.btn:hover {
  background-color: #555;
}

.input {
  padding: 0.8rem 1.2rem;
  border: none;
  border-radius: 4px;
  margin-right: 1rem;
  font-size: 1rem;
}

@media only screen and (max-width: 768px) {
  .input-box {
    padding: 1rem;
  }

  h1 {
    font-size: 1.5rem;
  }

  .btn {
    font-size: 0.9rem;
    padding: 0.6rem 1rem;
    margin-right: 0.5rem;
  }

  .input {
    font-size: 0.9rem;
    padding: 0.6rem 1rem;
    margin-right: 0.5rem;
  }
}
</style>