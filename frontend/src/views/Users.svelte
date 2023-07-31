<script lang="ts">
  import AddUser from "../components/users/AddUser.svelte";
  import DelUser from "../components/users/DelUser.svelte";
  import ModifyUser from "../components/users/ModifyUser.svelte";
  import AddGroup from "../components/groups/AddGroup.svelte";
  import DelGroup from "../components/groups/DelGroup.svelte";
  import ModifyGroup from "../components/groups/ModifyGroup.svelte";
  import { openDialog } from "../functions/functions";
  import CustomDialog from "../components/dialogs/CustomDialog.svelte";
  import infoIcon from "../assets/images/info.png";

  let viewState = "default";
  let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

  function setViewState(newViewState: string): void {
    viewState = newViewState;
  }

  function openInfo() {
    dialog = openDialog(
      dialog,
      "Info",
      `
            <p>
            This is the Users and Groups page. Here you can add, delete and modify users and groups.
            </p>
            <p>
              <b>Add User</b>: click on the "Add User" button. You will be prompted to enter the username, password, UID, group and shell for the new user.
            </p>
            <p>
              <b>Delete User</b>: click on the "Delete User" button. You will be prompted to enter the username of the user you want to delete.
            </p>
            <p>
              <b>Modify User</b>: click on the "Modify User" button. You will be prompted to enter the username, password, UID, group, shell and home for the user.
            </p>
            <p>
              <b>Create Group</b>: click on the "Create Group" button. You will be prompted to enter the name and GID for the new group.
            </p>
            <p>
              <b>Delete Group</b>: click on the "Delete Group" button. You will be prompted to enter the name of the group you want to delete.
            </p>
            <p>
              <b>Modify Group</b>: click on the "Modify Group" button. You will be prompted to enter the name and GID for the group.
            </p>
        `
    );
  }

  function onDialogClose() {
    dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
  }
</script>

<CustomDialog
  bind:show={dialog.showDialog}
  title="Info"
  message={dialog.dialogMessage}
  onClose={onDialogClose}
/>

{#if viewState === "default"}
  <main>
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
      <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    <div class="section" id="users">
      <h1>Users</h1>
      <button class="btn" on:click={() => setViewState("addUser")}
        >Add User</button
      >
      <button class="btn" on:click={() => setViewState("delUser")}
        >Delete User</button
      >
      <button class="btn" on:click={() => setViewState("modUser")}
        >Modify User</button
      >
    </div>

    <div class="section" id="groups">
      <h1>Groups</h1>
      <button class="btn" on:click={() => setViewState("addGroup")}
        >Create Group</button
      >
      <button class="btn" on:click={() => setViewState("delGroup")}
        >Delete Group</button
      >
      <button class="btn" on:click={() => setViewState("modGroup")}
        >Modify Group</button
      >
    </div>
  </main>
{:else if viewState === "addUser"}
  <AddUser on:dismiss={() => setViewState("default")} />
{:else if viewState === "delUser"}
  <DelUser on:dismiss={() => setViewState("default")} />
{:else if viewState === "modUser"}
  <ModifyUser on:dismiss={() => setViewState("default")} />
{:else if viewState === "addGroup"}
  <AddGroup on:dismiss={() => setViewState("default")} />
{:else if viewState === "delGroup"}
  <DelGroup on:dismiss={() => setViewState("default")} />
{:else if viewState === "modGroup"}
  <ModifyGroup on:dismiss={() => setViewState("default")} />
{/if}

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

  .info-button {
    position: absolute;
    top: 10px;
    right: 10px;
    border: none;
    background: #5a5858;
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
  }

  .info-button:hover {
    background: #7a7979;
  }

  .info-icon {
    max-width: none;
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
