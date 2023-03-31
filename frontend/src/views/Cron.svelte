<!-- CronJobs.svelte -->

<script lang="ts">
    import { onMount } from 'svelte';
    import {ListCronJobs, RemoveAllCronJobs} from '../../wailsjs/go/main/App.js'

    let jobs = [];
    let currentView = 'cron';

    function handleViewChange(event: CustomEvent<string>) {
    currentView = event.detail;
  }

    async function getCronJobs() {
        jobs = await ListCronJobs();
    }

    function removeAllCronJobs() {
        RemoveAllCronJobs();
        jobs = [];
    }

    onMount(() => {
      addEventListener('changeView', (event: CustomEvent) => {
        currentView = event.detail;
      });
    });

    onMount(async () => {
        await getCronJobs();
    });


</script>

<div class="container">
    <h1>Cron Jobs</h1>

    <button on:click={removeAllCronJobs}>Remove all cron jobs</button>

    {#if jobs.length < 1}
        <p>No jobs found</p>
    {:else}
        <ul>
            {#each jobs as job}
                <li>{job.Schedule} - {job.Command}</li>
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
    }

    /* Component styles */
    .container {
        max-width: 800px;
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
    }

    li {
        padding: 10px 20px;
        margin-bottom: 10px;
        margin-top: 10px;
        background-color: #a70a0a;
        border-radius: 5px;
    }

    li:hover {
        background-color: #d8cf4b;
    }

    button {
        padding: 10px 20px;
        background-color: #a70a0a;
        border: none;
        color: white;
        font-weight: bold;
        cursor: pointer;
        margin-bottom: 20px;
    }
</style>