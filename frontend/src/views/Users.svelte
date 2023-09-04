<script lang="ts">
  import { onMount } from "svelte";
  import AddUser from "../components/users/AddUser.svelte";
  import DelUser from "../components/users/DelUser.svelte";
  import ModifyUser from "../components/users/ModifyUser.svelte";
  import AddGroup from "../components/groups/AddGroup.svelte";
  import DelGroup from "../components/groups/DelGroup.svelte";
  import ModifyGroup from "../components/groups/ModifyGroup.svelte";
  import { CheckAdmin } from "../../wailsjs/go/backend/Backend";
  import { openDialog } from "../functions/functions";
  import CustomDialog from "../components/dialogs/CustomDialog.svelte";
  import infoIcon from "../assets/images/info.png";
  import { settings } from "../stores";

  let fontSize: string;
  let color: string;
  let fontFamily: string;
  let backgroundColor: string;
  let backgroundColor2: string;
  let inputColor: string;
  let buttonColor: string;
  let showInfoButton = false;
  settings.subscribe(($settings) => {
    fontSize = $settings.fontSize;
    color = $settings.color;
    fontFamily = $settings.fontFamily;
    backgroundColor = $settings.backgroundColor;
    backgroundColor2 = $settings.backgroundColor2;
    inputColor = $settings.inputColor;
    buttonColor = $settings.buttonColor;
    showInfoButton = $settings.showInfoButton;
  });

  $: {
    document.documentElement.style.setProperty("--main-font-size", fontSize);
    document.documentElement.style.setProperty("--main-color", color);
    document.documentElement.style.setProperty(
      "--main-font-family",
      fontFamily
    );
    document.documentElement.style.setProperty(
      "--main-bg-color",
      backgroundColor
    );
    document.documentElement.style.setProperty(
      "--main-bg-color2",
      backgroundColor2
    );
    document.documentElement.style.setProperty(
      "--main-input-color",
      inputColor
    );
    document.documentElement.style.setProperty(
      "--main-button-color",
      buttonColor
    );
  }

  let viewState = "default";
  let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

  async function checkPrivileges() {
  try {
    const admin = await CheckAdmin();
    if (!admin) {
      dialog = openDialog(
        dialog,
        "Warning",
        `
        <p style="color: ${color};">
           You are not running this application as root. You will not be able to add, delete or modify users and groups.
        </p>
        `
      );
    }
  } catch (error) {
    console.error("Error checking admin privileges:", error);
  }
}


  function setViewState(newViewState: string): void {
    viewState = newViewState;
  }

  function openInfo() {
    dialog = openDialog(
      dialog,
      "Info",
      `
        <p style="color: ${color};">
            This is the Users and Groups page. Here you can add, delete and modify users and groups.
        </p>
        <p style="color: ${color};">
            <b>Add User</b>: click on the "Add User" button. You will be prompted to enter the username, password, UID, group and shell for the new user.
        </p>
        <p style="color: ${color};">
            <b>Delete User</b>: click on the "Delete User" button. You will be prompted to enter the username of the user you want to delete.
        </p>
        <p style="color: ${color};">
            <b>Modify User</b>: click on the "Modify User" button. You will be prompted to enter the username, password, UID, group, shell and home for the user.
        </p>
        <p style="color: ${color};">
            <b>Create Group</b>: click on the "Create Group" button. You will be prompted to enter the name and GID for the new group.
        </p>
        <p style="color: ${color};">
            <b>Delete Group</b>: click on the "Delete Group" button. You will be prompted to enter the name of the group you want to delete.
        </p>
        <p style="color: ${color};">
            <b>Modify Group</b>: click on the "Modify Group" button. You will be prompted to enter the name and GID for the group.
        </p>
        `
    );
  }

  function onDialogClose() {
    dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
  }

  onMount(async () => {
    await checkPrivileges();
  });
</script>

<CustomDialog
  bind:show={dialog.showDialog}
  title={dialog.dialogTitle}
  message={dialog.dialogMessage}
  onClose={onDialogClose}
/>

{#if viewState === "default"}
  <main>
    {#if showInfoButton}
      <button
        type="button"
        class="info-button"
        title="Info"
        on:click={openInfo}
      >
        <img src={infoIcon} alt="Open Info" class="info-icon" />
      </button>
    {/if}
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
    background-color: var(--main-bg-color2);
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-right: 2rem;
  }

  h1 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    color: var(--main-color);
  }

  .btn {
    background: var(--main-button-color);
    color: #fff;
    border: none;
    border-radius: 4px;
    padding: 0.8rem 1.2rem;
    margin-right: 1rem;
    cursor: pointer;
    font-size: 1rem;
    transition: background-color 0.3s ease-in-out;
  }

  .info-button {
    position: absolute;
    top: 10px;
    right: 10px;
    border: none;
    background: var(--main-button-color);
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
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
