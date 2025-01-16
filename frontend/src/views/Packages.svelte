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
        CheckAdmin,
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

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
            const res = await OpenDialogQuestion(
                `Are you sure you want to remove ${name}?`,
            );
            if (res === "Yes") {
                await RemovePackage(name);
                packages = packages.filter((pkg) => pkg !== name);
                filteredPackages = filteredPackages.filter(
                    (pkg) => pkg !== name,
                );
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
            const res = await OpenDialogQuestion(
                `Are you sure you want to install ${name}?`,
            );
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
                    <p style="color: var(--color-text-primary); font-size: 1rem;">
                        You are not running this application as root. You will not be able to install or remove packages.
                    </p>
                    `,
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
                                    <span class="material-icons">delete</span>
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
    :global(:root) {
        /* Tokyo Night theme colors */
        --color-bg-primary: #1a1b26;
        --color-bg-secondary: #16161e;
        --color-bg-tertiary: #1f2335;
        --color-text-primary: #a9b1d6;
        --color-text-secondary: #787c99;
        --color-accent-blue: #7aa2f7;
        --color-accent-purple: #9d7cd8;
        --color-accent-cyan: #7dcfff;
        --color-accent-green: #9ece6a;
        --color-accent-orange: #ff9e64;
        --color-accent-red: #f7768e;

        /* Layout */
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

    ul {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    li {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-left: 5px;
        padding: var(--spacing-sm) 0;
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    li:last-child {
        border-bottom: none;
    }

    button {
        padding: var(--spacing-sm) var(--spacing-md);
        border: none;
        border-radius: 0.5rem;
        background-color: var(--color-accent-blue);
        color: var(--color-bg-primary);
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    button:hover {
        background-color: var(--color-accent-cyan);
    }

    .delete-btn {
        background: none;
        border: none;
        padding: 0;
        margin: 0;
        cursor: pointer;
    }

    .material-icons {
        color: var(--color-accent-red);
        font-size: 1.5rem;
    }

    .input-container {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-md);
    }

    #install-input {
        flex-grow: 1;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    .search-box input {
        width: 100%;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .loading {
        text-align: center;
    }

    .loading p {
        color: var(--color-text-primary);
        font-size: 1rem;
        margin-bottom: var(--spacing-md);
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.3);
        border-top: 4px solid var(--color-accent-blue);
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
