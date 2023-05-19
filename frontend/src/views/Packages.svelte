<script lang="ts">
  import { onMount } from "svelte";
  import {
    ListPackages,
    RemovePackage,
    InstallPackage,
  } from '../../wailsjs/go/backend/Backend';
  import deleteIcon from "../assets/images/delete.png";

  let packages: string[] = [];
  let filteredPackages: string[] = [];
  let searchInput: string = "";
  let installInput: string = "";

  function listPackages() {
    ListPackages().then((result) => {
      packages = result;
      filteredPackages = packages;
    });
  }

  function confirmDelete(name: string) {
    if (confirm(`Are you sure you want to delete ${name}?`)) {
      deletePackage(name);
    }
  }

  function deletePackage(name: string) {
    RemovePackage(name)
      .then(() => {
        packages = packages.filter((pkg) => pkg !== name);
        filteredPackages = filteredPackages.filter((pkg) => pkg !== name);
        alert(`Successfully deleted ${name}`);
      })
      .catch((err) => {
        alert(`Failed to delete ${name}: ${err}`);
      });
  }

  function confirmInstall(name: string) {
    if (confirm(`Are you sure you want to install ${name}?`)) {
      installPackage(name);
    }
  }

  function installPackage(name: string) {
    InstallPackage(name)
      .then(() => {
        packages.push(name);
        filteredPackages.push(name);
        alert(`Successfully installed ${name}`);
      })
      .catch((err) => {
        alert(`Failed to install ${name}: ${err}`);
      });
  }

  function handleSearch(event: Event) {
    const target = event.target as HTMLInputElement;
    searchInput = target.value.trim();
    filteredPackages = search();
  }

  function handleInstall() {
  const packageName = installInput.trim();
  if (packageName !== "") {
    confirmInstall(packageName);
    installInput = "";
  }
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
      <button on:click={handleInstall}>Install</button>
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
              <button class="delete-btn" on:click={() => confirmDelete(pkg)}>
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
