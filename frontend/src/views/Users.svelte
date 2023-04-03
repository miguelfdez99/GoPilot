<script lang="ts">

  import { onMount } from 'svelte';
  import AddUser from '../components/AddUser.svelte';
  import {PrintUsers, CreateUser, DeleteUser, CheckAdmin} from '../../wailsjs/go/main/App.js'

  let currentView = 'users';
  let showAddUser = false;
  let name: string
  let username: string
  let adminText: string

function printUser(): void {
      PrintUsers()
    }

    function createUser(): void {
      showAddUser = true;
      currentView = '';
    }

    function deleteUser(): void {
      DeleteUser()
    }

  function toggleAddUser(): void {
    showAddUser = !showAddUser;
  }


  onMount(() => {
    addEventListener('changeView', (event: CustomEvent) => {
      currentView = event.detail;
    });
    CheckAdmin().then(result => {
        if (result === false) {
          adminText = "You are not an admin"
          alert(adminText)
        }
    });
  });
</script>

<main>
  <div class="section" id="users">
    <h1>Users</h1>
    <button class="btn" on:click={printUser}>Print Users</button>
    <button class="btn" on:click={createUser}>Create User</button>
    <button class="btn" on:click={deleteUser}>Delete User</button>
    <button class="btn" on:click={toggleAddUser}>Add User</button>
    <div style="display: flex; flex-direction: column-reverse;">
      {#if showAddUser}
        <AddUser />
      {/if}
    </div>
  </div>

  <div class="section" id="groups">
    <h1>Groups</h1>
    <button class="btn" on:click={printUser}>Print Groups</button>
    <button class="btn" on:click={createUser}>Create Group</button>
    <button class="btn" on:click={deleteUser}>Delete Group</button>
    <button class="btn" on:click={createUser}>Add User to Group</button>
    <button class="btn" on:click={deleteUser}>Remove User from Group</button>
  </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

  .section {
    background-color: #f5f5f5;
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-bottom: 2rem;
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

  @media only screen and (max-width: 768px) {
    .section {
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
  }
</style>