<!-- CronJobForm.svelte -->
<script>
    import { AddCronJob } from '../../wailsjs/go/main/App.js';
    let schedule = '';
    let command = '';

    async function onSubmit() {
        if (schedule && command) {
            await AddCronJob(schedule, command);
            schedule = '';
            command = '';
        }
    }
</script>

<div class="container">
    <form on:submit|preventDefault="{onSubmit}">
        <h2>Add a New Cron Job</h2>
        <label>
            Schedule:
            <input type="text" bind:value="{schedule}" />
        </label>
        <label>
            Command:
            <input type="text" bind:value="{command}" />
        </label>
        <button type="submit">Add Job</button>
    </form>
</div>

<style>
    /* Reset styles */
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
        margin-bottom: 20px;
    }

    /* Component styles */
    .container {
        max-width: 800px;
        margin: 0 auto;
        padding: 20px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }

    h2 {
        font-size: 24px;
        margin-bottom: 20px;
    }

    form {
        display: flex;
        flex-direction: column;
    }

    label {
        display: flex;
        flex-direction: column;
        margin-bottom: 20px;
    }

    input {
        padding: 10px;
        font-size: 16px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }

    button[type="submit"] {
        padding: 10px;
        font-size: 16px;
        border: none;
        border-radius: 5px;
        background-color: #a70a0a;
        color: white;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    button[type="submit"]:hover {
        background-color: #721c24;
    }
</style>
