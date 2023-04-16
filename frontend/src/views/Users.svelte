<script lang="ts">
  import { onMount } from "svelte";
  import AddUser from "../components/AddUser.svelte";
  import DelUser from "../components/DelUser.svelte";
  import ModifyUser from "../components/ModifyUser.svelte";
  import AddGroup from "../components/AddGroup.svelte";
  import DelGroup from "../components/DelGroup.svelte";
  import ModifyGroup from "../components/ModifyGroup.svelte";
  import { CheckAdmin } from "../../wailsjs/go/main/App.js";

  let showUserAdd = false;
  let showUserDel = false;
  let showUserMod = false;
  let showGroupAdd = false;
  let showGroupDel = false;
  let showGroupMod = false;
  let adminText: string;

  function toggleAddUser(): void {
    showUserAdd = !showUserAdd;
    if (showUserAdd) {
      showUserDel = false;
      showUserMod = false;
      showGroupAdd = false;
      showGroupDel = false;
      showGroupMod = false;
    }
  }

  function toggleDelUser(): void {
    showUserDel = !showUserDel;
    if (showUserDel) {
      showUserAdd = false;
      showUserMod = false;
      showGroupAdd = false;
      showGroupDel = false;
      showGroupMod = false;
    }
  }

  function toggleModUser(): void {
    showUserMod = !showUserMod;
    if (showUserMod) {
      showUserAdd = false;
      showUserDel = false;
      showGroupAdd = false;
      showGroupDel = false;
      showGroupMod = false;
    }
  }

  function toggleAddGroup(): void {
    showGroupAdd = !showGroupAdd;
    if (showGroupAdd) {
      showGroupDel = false;
      showGroupMod = false;
      showUserAdd = false;
      showUserDel = false;
      showUserMod = false;
    }
  }

  function toggleDelGroup(): void {
    showGroupDel = !showGroupDel;
    if (showGroupDel) {
      showGroupAdd = false;
      showGroupMod = false;
      showUserAdd = false;
      showUserDel = false;
      showUserMod = false;
    }
  }

  function toggleModGroup(): void {
    showGroupMod = !showGroupMod;
    if (showGroupMod) {
      showGroupAdd = false;
      showGroupDel = false;
      showUserAdd = false;
      showUserDel = false;
      showUserMod = false;
    }
  }

  onMount(() => {
    CheckAdmin().then((result) => {
      if (result === false) {
        adminText = "You are not an admin";
        alert(adminText);
      }
    });
  });
</script>

<main>
  <div class="section" id="users">
    <h1>Users</h1>
    <button class="btn" on:click={toggleAddUser}>Add User</button>
    <button class="btn" on:click={toggleDelUser}>Delete User</button>
    <button class="btn" on:click={toggleModUser}>Modify User</button>
    <div style="display: flex; flex-direction: column-reverse;">
      {#if showUserAdd}
        <AddUser on:clickOutside={toggleAddUser} />
      {:else if showUserDel}
        <DelUser on:clickOutside={toggleDelUser} />
      {:else if showUserMod}
        <ModifyUser on:clickOutside={toggleModUser} />
      {/if}
    </div>
  </div>

  <div class="section" id="groups">
    <h1>Groups</h1>
    <button class="btn" on:click={toggleAddGroup}>Create Group</button>
    <button class="btn" on:click={toggleDelGroup}>Delete Group</button>
    <button class="btn" on:click={toggleModGroup}>Modify Group</button>
    <div style="display: flex; flex-direction: column-reverse;">
      {#if showGroupAdd}
        <AddGroup on:clickOutside={toggleAddGroup} />
      {:else if showGroupDel}
        <DelGroup on:clickOutside={toggleDelGroup} />
      {:else if showGroupMod}
        <ModifyGroup on:clickOutside={toggleModGroup} />
      {/if}
    </div>
  </div>
</main>

<style>
  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    padding-top: 100px;
    padding-bottom: 30px;
  }

  .section {
    background-color: #f5f5f5;
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-right: 2rem;
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
