<script lang="ts">
    import { onMount } from "svelte";
    import {
        ListCronJobs,
        RemoveAllCronJobs,
        RemoveCronJob,
        OpenDialogInfo,
        OpenDialogError,
        OpenDialogQuestion,
    } from "../../wailsjs/go/backend/Backend";
    import AddCronJob from "../components/AddCronJob.svelte";
    import { createEventDispatcher } from "svelte";
    import { checkCommand } from "../functions/functions";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let jobs = [];
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    let viewState: string = "default";

    function setViewState(newViewState: string): void {
        viewState = newViewState;
    }

    async function getCronJobs() {
        jobs = await ListCronJobs();
    }

    async function removeAllCronJobs() {
        try {
            const res = await OpenDialogQuestion(
                "Are you sure you want to delete all jobs?"
            );
            if (res === "Yes") {
                await RemoveAllCronJobs();
                await OpenDialogInfo("Successfully deleted all jobs");
                jobs = [];
            } else {
                await OpenDialogInfo("Canceled job deletion");
            }
        } catch (err) {
            await OpenDialogError(`Failed to delete all jobs: ${err}`);
        }
    }

    onMount(async () => {
        await checkCommand("crontab");
        getCronJobs();
    });

    function deleteCron(name: string) {
        RemoveCronJob(name)
            .then(() => {
                OpenDialogInfo(`Successfully deleted ${name}`);
                jobs = jobs.filter((job) => job.Command !== name);
            })
            .catch((err) => {
                OpenDialogError(`Failed to delete ${name}: ${err}`);
            });
    }
</script>

<div class="cron-container">
    <h1>Cron Jobs</h1>

    {#if viewState === "default"}
        <div class="actions">
            <button on:click={() => setViewState("addCronJob")}>
                <span class="material-icons">add</span>
                Add Cron Job
            </button>
            <button on:click={removeAllCronJobs}>
                <span class="material-icons">delete_forever</span>
                Remove All Cron Jobs
            </button>
        </div>

        <h2>Crontab List</h2>
        {#if jobs.length < 1}
            <p>No jobs found</p>
        {:else}
            <ul class="job-list">
                {#each jobs as job}
                    <li>
                        <span class="job-schedule">{job.Schedule}</span>
                        <span class="job-command">{job.Command}</span>
                        <button
                            class="delete-btn"
                            on:click={() => deleteCron(job.Command)}
                        >
                            <span class="material-icons">delete</span>
                        </button>
                    </li>
                {/each}
            </ul>
        {/if}
    {:else if viewState === "addCronJob"}
        <AddCronJob
            on:cronJobAdded={() => {
                getCronJobs();
                setViewState("default");
            }}
            on:dismiss={() => setViewState("default")}
        />
    {/if}
</div>

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

    .cron-container {
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        font-family: sans-serif;
    }

    h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    h2 {
        color: var(--color-text-primary);
        margin-bottom: var(--spacing-md);
        font-size: 1.25rem;
    }

    .actions {
        display: flex;
        gap: var(--spacing-md);
        margin-bottom: var(--spacing-lg);
    }

    button {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        padding: var(--spacing-sm) var(--spacing-md);
        background-color: var(--color-bg-secondary);
        border: 1px solid var(--color-bg-tertiary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    button:hover {
        background-color: var(--color-bg-tertiary);
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-accent-blue);
    }

    .job-list {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .job-list li {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: var(--spacing-sm);
        border-bottom: 1px solid var(--color-bg-tertiary);
    }

    .job-list li:last-child {
        border-bottom: none;
    }

    .job-schedule {
        color: var(--color-text-secondary);
        font-family: monospace;
    }

    .job-command {
        flex: 1;
        margin: 0 var(--spacing-md);
    }

    .delete-btn {
        background: none;
        border: none;
        cursor: pointer;
        padding: var(--spacing-sm);
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .delete-btn .material-icons {
        color: var(--color-accent-red);
    }

    .delete-btn:hover .material-icons {
        color: var(--color-accent-orange);
    }
</style>
