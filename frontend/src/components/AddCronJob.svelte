<script lang="ts">
    import {
        AddCronJob,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let minute: string = "";
    let hour: string = "";
    let day: string = "";
    let month: string = "";
    let dayOfWeek: string = "";
    let command: string = "";
    let composedJob: string = "";
    let quickScheduleOption: string = "";

    $: {
        if (quickScheduleOption) {
            composedJob = `${quickScheduleOption} ${command}`;
        } else {
            composedJob = `${minute} ${hour} ${day} ${month} ${dayOfWeek} ${command}`;
        }
    }

    function dismiss() {
        dispatch("dismiss");
    }

    async function onSubmit() {
        if (minute && hour && day && month && dayOfWeek && command) {
            const schedule = `${minute} ${hour} ${day} ${month} ${dayOfWeek}`;
            await AddCronJob(schedule, command);
            minute = "";
            hour = "";
            day = "";
            month = "";
            dayOfWeek = "";
            command = "";
            dispatch("cronJobAdded", null);
        } else if (quickScheduleOption && command) {
            await AddCronJob(quickScheduleOption, command);
            quickScheduleOption = "";
            command = "";
            dispatch("cronJobAdded", null);
        } else {
            await OpenDialogError(
                "Please fill out all fields or select a quick schedule option"
            );
        }
    }
</script>

<div class="container">
    <form class="form" on:submit|preventDefault={onSubmit}>
        <h2 class="form-title">Add a New Cron Job</h2>

        <div class="input-container">
            <span class="input-indicator">Quick Schedule:</span>
            <span class="input-value">
                <select class="form-select" bind:value={quickScheduleOption}>
                    <option value="">None</option>
                    <option value="@reboot">Startup</option>
                    <option value="@hourly">Hourly</option>
                    <option value="@daily">Daily</option>
                    <option value="@weekly">Weekly</option>
                    <option value="@monthly">Monthly</option>
                    <option value="@yearly">Yearly</option>
                </select>
            </span>
        </div>

        {#if !quickScheduleOption}
            <label class="form-label">
                <div class="input-container">
                    <span class="input-indicator">Minutes:</span>
                    <span class="input-value">
                        <input
                            class="form-input"
                            type="text"
                            bind:value={minute}
                        />
                    </span>
                </div>
                <div class="input-container">
                    <span class="input-indicator">Hour:</span>
                    <span class="input-value">
                        <input
                            class="form-input"
                            type="text"
                            bind:value={hour}
                        />
                    </span>
                </div>
                <div class="input-container">
                    <span class="input-indicator">Day:</span>
                    <span class="input-value">
                        <input
                            class="form-input"
                            type="text"
                            bind:value={day}
                        />
                    </span>
                </div>
                <div class="input-container">
                    <span class="input-indicator">Month:</span>
                    <span class="input-value">
                        <input
                            class="form-input"
                            type="text"
                            bind:value={month}
                        />
                    </span>
                </div>
                <div class="input-container">
                    <span class="input-indicator">Day of Week:</span>
                    <span class="input-value">
                        <input
                            class="form-input"
                            type="text"
                            bind:value={dayOfWeek}
                        />
                    </span>
                </div>
            </label>
        {/if}
        <div class="input-container">
            <span class="input-indicator">Command:</span>
            <span class="input-value">
                <input class="form-input" type="text" bind:value={command} />
            </span>
        </div>

        <label class="form-label">
            Job:
            <span class="input-value">
                <input
                    class="form-input"
                    type="text"
                    bind:value={composedJob}
                    readonly
                />
            </span>
        </label>

        <p class="form-note">
            Enter cron job schedule in the format: <code
                >minute hour day month day-of-week</code
            >
        </p>
        <p class="form-note">
            For example: <code>30 * * * *</code> for every hour at 30 minutes past
            the hour
        </p>
        <p class="form-note">
            Or: <code>0 0 * * *</code> for midnight every day
        </p>
        <p class="form-note">
            Refer to cron syntax for more information on how to specify cron job
            schedules.
        </p>
        <p class="form-note">
            Or select a quick schedule from the dropdown above.
        </p>
        <button class="form-button" type="submit">Add Job</button>
        <button class="back-button" on:click={dismiss}>Back</button>
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

    .form-select {
        color: white;
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
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
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
        background-color: #1abc9c;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
        transition: background-color 0.2s;
    }

    .form-button:hover {
        background-color: #16a085;
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

    .back-button {
        right: 0;
        padding: 10px 20px;
        border: none;
        background-color: #333;
        border-radius: 5px;
        font-size: 16px;
        color: #fff;
        transition: background-color 0.3s;
    }

    .back-button:hover {
        background-color: #555;
    }
</style>
