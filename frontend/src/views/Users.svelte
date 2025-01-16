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
                    <p>
                        You are not running this application as root. You will not be able to add, delete, or modify users and groups.
                    </p>
                    `,
                );
            }
        } catch (error) {
            console.error("Error checking admin privileges:", error);
        }
    }

    function setViewState(newViewState: string): void {
        viewState = newViewState;
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
    <main class="users-container">
        <div class="section" id="users">
            <h1>Users</h1>
            <button class="btn" on:click={() => setViewState("addUser")}>
                <span class="material-icons">person_add</span>
                Add User
            </button>
            <button class="btn" on:click={() => setViewState("delUser")}>
                <span class="material-icons">person_remove</span>
                Delete User
            </button>
            <button class="btn" on:click={() => setViewState("modUser")}>
                <span class="material-icons">edit</span>
                Modify User
            </button>
        </div>

        <div class="section" id="groups">
            <h1>Groups</h1>
            <button class="btn" on:click={() => setViewState("addGroup")}>
                <span class="material-icons">group_add</span>
                Create Group
            </button>
            <button class="btn" on:click={() => setViewState("delGroup")}>
                <span class="material-icons">group_remove</span>
                Delete Group
            </button>
            <button class="btn" on:click={() => setViewState("modGroup")}>
                <span class="material-icons">edit</span>
                Modify Group
            </button>
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

    .users-container {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: var(--spacing-lg);
        padding: var(--spacing-lg);
    }

    .section {
        background-color: var(--color-bg-secondary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        text-align: center;
    }

    h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .btn {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        padding: var(--spacing-sm) var(--spacing-md);
        background-color: var(--color-bg-tertiary);
        border: 1px solid var(--color-bg-tertiary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
        margin-bottom: var(--spacing-sm);
    }

    .btn:hover {
        background-color: var(--color-bg-secondary);
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-accent-blue);
    }
</style>
