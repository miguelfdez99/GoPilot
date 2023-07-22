<script lang="ts">
    import { onMount } from "svelte";
    import {
        ListCronJobs,
        RemoveAllCronJobs,
        RemoveCronJob,
    } from "../../wailsjs/go/backend/Backend";
    import AddCronJob from "../components/AddCronJob.svelte";
    import deleteIcon from "../assets/images/delete.png";
    import { openDialog, closeDialog, checkCommand  } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    let showDeleteAllDialog: boolean = false;

    let jobs = [];
    let showAddCronJob: boolean = false;
    let showDeleteDialog: boolean = false;
    let dialogName: string = "";

    function confirmDeleteDialog(name: string) {
        dialogName = name;
        showDeleteDialog = true;
    }

    function onDialogConfirm() {
        deleteCron(dialogName);
        showDeleteDialog = false;
    }

    function onDialogClose() {
        showDeleteDialog = false;
    }

    function confirmDeleteAllDialog() {
    showDeleteAllDialog = true;
  }

  function onDeleteAllDialogConfirm() {
    removeAllCronJobs();
    showDeleteAllDialog = false;
  }

  function onDeleteAllDialogClose() {
    showDeleteAllDialog = false;
  }

    async function getCronJobs() {
        jobs = await ListCronJobs();
    }

    function removeAllCronJobs() {
    RemoveAllCronJobs()
      .then(() => {
        dialog = openDialog(dialog,'Success', 'Successfully deleted all jobs');
        jobs = [];
      })
      .catch((err) => {
        dialog = openDialog(dialog,'Error', `Failed to delete all jobs: ${err}`);
      });
  }

    onMount(async () => {
        await checkCommand("crontab", dialog);
        getCronJobs();
    });

    function toggleAddCronJob() {
        showAddCronJob = !showAddCronJob;
    }

    function deleteCron(name: string) {
        RemoveCronJob(name)
            .then(() => {
                dialog = openDialog(dialog,'Success', `Successfully deleted ${name}`);
                jobs = jobs.filter((job) => job.Command !== name);
            })
            .catch((err) => {
                dialog = openDialog(dialog,'Error', `Failed to delete ${name}: ${err}`);
            });
    }

    $: addButtonText = showAddCronJob ? "Hide Add Cron Job" : "Add Cron Job";
</script>

<CustomDialog
    bind:show={showDeleteDialog}
    title="Delete Cron Job"
    message={`Are you sure you want to delete ${dialogName}?`}
    onConfirm={onDialogConfirm}
    onClose={onDialogClose}
/>

<CustomDialog
  bind:show={dialog.showDialog}
  title={dialog.dialogTitle}
  message={dialog.dialogMessage}
  onClose={() => dialog = closeDialog(dialog)}
  confirmButton={false}
/>

<CustomDialog
  bind:show={showDeleteAllDialog}
  title="Delete All Cron Jobs"
  message="Are you sure you want to delete all cron jobs?"
  onConfirm={onDeleteAllDialogConfirm}
  onClose={onDeleteAllDialogClose}
/>

<div class="container">
    <h1>Cron Jobs</h1>

    <button on:click={toggleAddCronJob}>{addButtonText}</button>

    {#if showAddCronJob}
        <AddCronJob on:cronJobAdded={() => getCronJobs()} />
    {/if}

    <button on:click={confirmDeleteAllDialog}>Remove all cron jobs</button>

    <h2>Crontab List</h2>
    {#if jobs.length < 1}
        <p>No jobs found</p>
    {:else}
        <ul>
            {#each jobs as job}
                <li>
                    {job.Schedule}
                    {job.Command}
                    <button
                        class="delete-btn"
                        on:click={() => confirmDeleteDialog(job.Command)}
                    >
                        <img
                            src={deleteIcon}
                            alt="Delete"
                            class="delete-icon"
                        />
                    </button>
                </li>
            {/each}
        </ul>
    {/if}
</div>

<style>
    .container {
        display: flex;
        flex-direction: column;
        max-width: 95%;
        margin: 2rem auto;
        background-color: #282828;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
        border-radius: 10px;
        color: #ddd;
        padding: 1rem;
    }

    h1, h2, p {
        color: #ddd;
        margin-bottom: 1em;
    }

    ul {
        list-style: none;
        padding: 0;
        border: 1px solid #333;
        border-radius: 4px;
        margin-bottom: 1em;
        background: #333;
    }

    li {
        padding: .8em;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid #444;
    }

    li:last-child {
        border-bottom: none;
    }

    button {
        padding: .8em 1em;
        border: none;
        border-radius: 4px;
        background: #1abc9c;
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 100%;
        margin-bottom: 1em;
    }

    button:hover {
        background-color: #16a085;
    }

    .delete-btn {
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
</style>

