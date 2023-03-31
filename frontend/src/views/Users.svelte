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