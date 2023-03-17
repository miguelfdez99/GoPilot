<script lang="ts">
    import {Greet, PrintUsers, CreateUser, DeleteUser, CreateUserWithDir, DeleteUserWithDir, CheckAdmin} from '../../wailsjs/go/main/App.js'
    import { onMount } from 'svelte';

    let resultText: string = "Please enter your name below 👇"
    let name: string
    let username: string
    let adminText: string

    function greet(): void {
      Greet(name).then(result => resultText = result)
    }

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

    function checkAdmin(): boolean {
      CheckAdmin()
    }

    onMount(() => {
      CheckAdmin().then(result => {
        if (result === false) {
          adminText = "You are not an admin"
        }
      })
    })


</script>

<main>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Greet</button>
    <button class="btn" on:click={() => name = ""}>Clear</button>
    <h1>This is a TEST</h1>
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


</style>
