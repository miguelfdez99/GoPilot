<script lang="ts">
  import { onMount } from "svelte";
  import {
    ListPackages,
    RemovePackage,
    InstallPackage,
  } from "../../wailsjs/go/backend/Backend";
  import CustomDialog from "../components/dialogs/CustomDialog.svelte";
  import { openDialog, closeDialog } from "../functions/functions";
  import deleteIcon from "../assets/images/delete.png";

  let packages: string[] = [];
  let filteredPackages: string[] = [];
  let searchInput: string = "";
  let installInput: string = "";
  let showDeleteDialog: boolean = false;
  let dialogName: string = "";
  let showInstallDialog: boolean = false;
  let installDialogName: string = "";
  let dialog = { showDialog: false, dialogTitle: '', dialogMessage: '' };

  function listPackages() {
    ListPackages().then((result) => {
      packages = result;
      filteredPackages = packages;
    });
  }

  function confirmDeleteDialog(name: string) {
    dialogName = name;
    showDeleteDialog = true;
  }

  function onDialogConfirm() {
    deletePackage(dialogName);
    showDeleteDialog = false;
  }

  function onDialogClose() {
    showDeleteDialog = false;
  }

  function confirmInstallDialog(name: string) {
    installDialogName = name;
    showInstallDialog = true;
  }

  function onInstallDialogConfirm() {
    installPackage(installDialogName);
    showInstallDialog = false;
  }

  function onInstallDialogClose() {
    showInstallDialog = false;
  }

  function deletePackage(name: string) {
    RemovePackage(name)
      .then(() => {
        packages = packages.filter((pkg) => pkg !== name);
        filteredPackages = filteredPackages.filter((pkg) => pkg !== name);
        dialog = openDialog(dialog, "Success", `Successfully deleted ${name}`);
      })
      .catch((err) => {
        dialog = openDialog(dialog, "Error", `Failed to delete ${name}: ${err}`);
      });
  }

  function installPackage(name: string) {
    InstallPackage(name)
      .then(() => {
        packages.push(name);
        filteredPackages.push(name);
        dialog = openDialog(dialog, "Success", `Successfully installed ${name}`);
      })
      .catch((err) => {
        dialog = openDialog(dialog, "Error", `Failed to install ${name}: ${err}`);
      });
  }

  function handleSearch(event: Event) {
    const target = event.target as HTMLInputElement;
    searchInput = target.value.trim();
    filteredPackages = search();
  }

  function search() {
    return packages.filter((pkg) => {
      return pkg.toLowerCase().includes(searchInput.toLowerCase());
    });
  }

  onMount(() => {
    listPackages();
  });
</script>

<CustomDialog
  bind:show={showDeleteDialog}
  title="Delete Package"
  message={`Are you sure you want to delete ${dialogName}?`}
  onConfirm={onDialogConfirm}
  onClose={onDialogClose}
/>

<CustomDialog
  bind:show={showInstallDialog}
  title="Install Package"
  message={`Are you sure you want to install ${installDialogName}?`}
  onConfirm={onInstallDialogConfirm}
  onClose={onInstallDialogClose}
/>

<CustomDialog
  bind:show={dialog.showDialog}
  title={dialog.dialogTitle}
  message={dialog.dialogMessage}
  onClose={() => dialog = closeDialog(dialog)}
  confirmButton={false}
/>

<main>
  <div class="input-box" id="input">
    <h1>Packages</h1>
    <div class="input-container">
      <input
        id="install-input"
        type="text"
        bind:value={installInput}
        placeholder="Enter package name..."
      />
      <button on:click={() => confirmInstallDialog(installInput)}
        >Install</button
      >
    </div>

    <div class="search-box">
      <input
        type="text"
        placeholder="Search packages"
        on:input={handleSearch}
      />
    </div>
    <div>
      {#if filteredPackages.length > 0}
        <ul>
          {#each filteredPackages as pkg}
            <li>
              {pkg}
              <button
                class="delete-btn"
                on:click={() => confirmDeleteDialog(pkg)}
              >
                <img src={deleteIcon} alt="Delete" class="delete-icon" />
              </button>
            </li>
          {/each}
        </ul>
      {:else}
        <p>No packages found</p>
      {/if}
    </div>
  </div>
</main>

<style>
  ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  li {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  button {
    margin-left: 10px;
    width: 80px;
    height: 30px;
    font-size: 14px;
    background: #1abc9c;
  }

  button:hover {
        background-color: #16a085;
  }

  .delete-btn {
    float: right;
    width: 24px;
    height: 24px;
    padding: 0;
    margin: 0;
    border: none;
    background: none;
  }

  .delete-icon {
    width: 100%;
    height: 100%;
  }

  .input-container {
    display: flex;
    align-items: center;
  }

  #install-input {
    flex-grow: 1;
    margin-right: 1rem;
  }

  button {
    width: 6rem;
    height: 2.5rem;
    font-size: 1rem;
    padding-top: 0.5rem;
  }

  h1,
  li,
  input,
  p {
    color: white;
  }
</style>
