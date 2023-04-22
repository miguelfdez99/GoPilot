<script>
    import { AddCronJob } from '../../wailsjs/go/main/App.js';
    let minute = '';
    let hour = '';
    let day = '';
    let month = '';
    let dayOfWeek = '';
    let command = '';
    let composedJob = '';

    $: {
    composedJob = `${minute} ${hour} ${day} ${month} ${dayOfWeek} ${command}`;
}

    async function onSubmit() {
        if (minute && hour && day && month && dayOfWeek && command) {
            const schedule = `${minute} ${hour} ${day} ${month} ${dayOfWeek}`;
            await AddCronJob(schedule, command);
            minute = '';
            hour = '';
            day = '';
            month = '';
            dayOfWeek = '';
            command = '';
        }
    }

</script>

<div class="container">
    <form class="form" on:submit|preventDefault="{onSubmit}">
        <h2 class="form-title">Add a New Cron Job</h2>
        <label class="form-label">
            <div class="input-container">
                <span class="input-indicator">Minutes:</span>
                <span class="input-value">
                    <input class="form-input" type="text" bind:value="{minute}" />
                </span>
            </div>
            <div class="input-container">
                <span class="input-indicator">Hour:</span>
                <span class="input-value">
                    <input class="form-input" type="text" bind:value="{hour}" />
                </span>
            </div>
            <div class="input-container">
                <span class="input-indicator">Day:</span>
                <span class="input-value">
                    <input class="form-input" type="text" bind:value="{day}" />
                </span>
            </div>
            <div class="input-container">
                <span class="input-indicator">Month:</span>
                <span class="input-value">
                    <input class="form-input" type="text" bind:value="{month}" />
                </span>
            </div>
            <div class="input-container">
                <span class="input-indicator">Day of Week:</span>
                <span class="input-value">
                    <input class="form-input" type="text" bind:value="{dayOfWeek}" />
                </span>
            </div>
        </label>
        <div class="input-container">
            <span class="input-indicator">Command:</span>
            <span class="input-value">
            <input class="form-input" type="text" bind:value="{command}" />
            </span>
        </div>

        <label class="form-label">
            Job:
            <span class="input-value">
                <input class="form-input" type="text" bind:value="{composedJob}" readonly />
            </span>
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
    flex-grow: 1;
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

.input-container {
        display: flex;
        align-items: center;
        margin-bottom: 10px;
    }

    .input-indicator {
        flex-basis: 50%;
        font-weight: bold;
    }
</style>
