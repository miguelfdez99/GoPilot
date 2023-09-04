<script lang="ts">
  import { onMount } from "svelte";
  import { writable } from "svelte/store";
  import {
    ListPackages,
    RemovePackage,
    InstallPackage,
    OpenDialogInfo,
    OpenDialogError,
    OpenDialogQuestion,
    CheckAdmin
  } from "../../wailsjs/go/backend/Backend";
  import { openDialog } from "../functions/functions";
  import CustomDialog from "../components/dialogs/CustomDialog.svelte";
  import deleteIcon from "../assets/images/delete.png";
  import { settings } from '../stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
    });

    $: {
    document.documentElement.style.setProperty('--main-font-size', fontSize);
    document.documentElement.style.setProperty('--main-color', color);
    document.documentElement.style.setProperty('--main-font-family', fontFamily);
    document.documentElement.style.setProperty('--main-bg-color', backgroundColor);
    document.documentElement.style.setProperty('--main-bg-color2', backgroundColor2);
    document.documentElement.style.setProperty('--main-input-color', inputColor);
    document.documentElement.style.setProperty('--main-button-color', buttonColor);
  }

  let packages: string[] = [];
  let filteredPackages: string[] = [];
  let searchInput: string = "";
  let installInput: string = "";
  let loading = writable(false);

  async function listPackages() {
    loading.set(true);
    try {
        const result = await ListPackages();
        packages = result;
        filteredPackages = packages;
    } catch (error) {
        await OpenDialogError(`Failed to list packages: ${error}`);
    } finally {
        loading.set(false);
    }
}

  async function deletePackage(name: string) {
      try {
          const res = await OpenDialogQuestion(`Are you sure you want to remove ${name}?`);
          if (res === "Yes") {
            await RemovePackage(name);
            packages = packages.filter((pkg) => pkg !== name);
            filteredPackages = filteredPackages.filter((pkg) => pkg !== name);
            await OpenDialogInfo(`Successfully deleted ${name}`);
          } else {
              await OpenDialogInfo(`Canceled package deletion`);
          }
      } catch (err) {
          await OpenDialogError(`Failed to delete ${name}: ${err}`);
      }
  }

  async function installPackage(name: string) {
      try {
          const res = await OpenDialogQuestion(`Are you sure you want to install ${name}?`);
          if (res === "Yes") {
              await InstallPackage(name);
              packages.push(name);
              filteredPackages.push(name);
              await OpenDialogInfo(`Successfully installed ${name}`);
          } else {
              await OpenDialogInfo(`Canceled package installation`);
          }
      } catch (err) {
          await OpenDialogError(`Failed to install ${name}: ${err}`);
      }
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

  let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

  async function checkPrivileges() {
  try {
    const admin = await CheckAdmin();
    if (!admin) {
      dialog = openDialog(
        dialog,
        "Warning",
        `
        <p style="color: ${color}; font-size: ${fontSize};">
           You are not running this application as root. You will not be able to install or remove packages.
        </p>
        `
      );
    }
  } catch (error) {
    console.error("Error checking admin privileges:", error);
  }
}

  function onDialogClose() {
    dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
  }

  onMount(async () => {
    await checkPrivileges();
    listPackages();
  });
</script>

<CustomDialog
  bind:show={dialog.showDialog}
  title={dialog.dialogTitle}
  message={dialog.dialogMessage}
  onClose={onDialogClose}
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
      <button on:click={() => installPackage(installInput)}
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
    {#if $loading}
        <div class="loading">
            <p>Loading...</p>
            <div class="spinner" />
        </div>
    {:else}
    <div>
      {#if filteredPackages.length > 0}
        <ul>
          {#each filteredPackages as pkg}
            <li>
              {pkg}
              <button
                class="delete-btn"
                on:click={() => deletePackage(pkg)}
              >
                <img src={deleteIcon} alt="Delete" class="delete-icon" />
              </button>
            </li>
          {/each}
        </ul>
      {/if}
    </div>
    {/if}
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
    background: var(--main-button-color);
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


  li,
  input,
  p {
    color: var(--main-color);
    font-size: var(--main-font-size);
    font-family: var(--main-font-family);
  }

  h1 {
    color: var(--main-color);
    font-family: var(--main-font-family);
  }

  input {
    background: var(--main-input-color);
  }

  .loading {
        text-align: center;
    }

    .loading p {
        color: var(--main-color);;
        font-size: 16px;
        margin-bottom: 20px;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--main-color);;
        border-radius: 50%;
        width: 30px;
        height: 30px;
        animation: spin 1s linear infinite;
        margin: 0 auto;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
