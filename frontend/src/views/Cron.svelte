<script lang="ts">
    import { onMount } from "svelte";
    import {
        ListCronJobs,
        RemoveAllCronJobs,
        RemoveCronJob,
        AddCronJob as AddCronJobAPI // Rename to avoid conflict with component name
    } from "../../wailsjs/go/main/App.js";
    import AddCronJob from "../components/AddCronJob.svelte";
    import deleteIcon from "../assets/images/delete.png";

    let jobs = [];
    let showAddCronJob = false;

    async function getCronJobs() {
        jobs = await ListCronJobs();
    }

    function removeAllCronJobs() {
        RemoveAllCronJobs();
        jobs = [];
    }

    onMount(() => {
        getCronJobs();
    });

    function toggleAddCronJob() {
        showAddCronJob = !showAddCronJob;
    }

    function deleteCron(name) {
        RemoveCronJob(name)
            .then(() => {
                alert(`Successfully deleted ${name}`);
                jobs = jobs.filter((job) => job.Command !== name);
            })
            .catch((err) => {
                alert(`Failed to delete ${name}: ${err}`);
            });
    }

    function confirmDelete(name) {
        if (confirm(`Are you sure you want to delete ${name}?`)) {
            deleteCron(name);
        }
    }

    $: addButtonText = showAddCronJob ? 'Hide Add Cron Job' : 'Add Cron Job';

</script>

<div class="container">
    <h1>Cron Jobs</h1>

    <button on:click={toggleAddCronJob}>{addButtonText}</button>

    {#if showAddCronJob}
        <AddCronJob />
    {/if}

    <button on:click={removeAllCronJobs}>Remove all cron jobs</button>

    <h2>Crontab List</h2>
    {#if jobs.length < 1}
        <p>No jobs found</p>
    {:else}
        <ul>
            {#each jobs as job}
                <li>{job.Schedule} {job.Command}
                    <button class="delete-btn" on:click={() => confirmDelete(job.Command)}>
                        <img src={deleteIcon} alt="Delete" class="delete-icon" />
                    </button>
                </li>
            {/each}
        </ul>
    {/if}

</div>


<style>
    /* Reset styles */
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        color: white;
    }

    /* Component styles */
    .container {
        margin: 0 auto;
        padding: 20px;
    }

    h1 {
        font-size: 32px;
        margin-bottom: 20px;
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
  }

    li:hover {
        background-color: #d8cf4b;
    }

    button {
        display: inline-block;
        padding: 8px 16px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
        transition: background-color 0.2s;
        margin-bottom: 20px;
        margin-top: 10px;
    }

    button:hover {
        background-color: #0056b3;
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
</style>
