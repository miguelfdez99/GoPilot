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
    <form class="form" on:submit|preventDefault="{onSubmit}">
        <h2 class="form-title">Add a New Cron Job</h2>
        <label class="form-label">
            Schedule:
            <input class="form-input" type="text" bind:value="{schedule}" />
        </label>
        <label class="form-label">
            Command:
            <input class="form-input" type="text" bind:value="{command}" />
        </label>
        <p class="form-note">Enter cron job schedule in the format: <code>minute hour day month day-of-week</code></p>
        <p class="form-note">For example: <code>30 * * * *</code> for every hour at 30 minutes past the hour</p>
        <p class="form-note">Or: <code>0 0 * * *</code> for midnight every day</p>
        <p class="form-note">Refer to cron syntax for more information on how to specify cron job schedules.</p>
        <button class="form-button" type="submit">Add Job</button>
    </form>
</div>

<style>
.container {
    max-width: 400px;
    margin: 0 auto;
}

h2 {
    color: white;
}

.form {
    padding: 20px;
    background-color: #333;
    color: #fff;
    border-radius: 4px;
}

.form-title {
    font-size: 24px;
    margin-bottom: 10px;
}

.form-label {
    display: block;
    margin-bottom: 10px;
    color: #fff;
}

.form-input {
    width: 100%;
    padding: 8px;
    border: none;
    border-radius: 4px;
    background-color: #555;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    color: #fff;
}

.form-note {
    font-size: 14px;
    margin-bottom: 10px;
    color: #fff;
}

.form-button {
    display: inline-block;
    padding: 8px 16px;
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
    transition: background-color 0.2s;
}

.form-button:hover {
    background-color: #0056b3;
}
</style>
