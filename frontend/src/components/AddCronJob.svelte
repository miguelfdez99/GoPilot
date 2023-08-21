<script lang="ts">
    import {
        AddCronJob,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";
    import { settings } from '../stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
    });

    $: {
    document.documentElement.style.setProperty('--main-font-size', fontSize);
    document.documentElement.style.setProperty('--main-color', color);
    document.documentElement.style.setProperty('--main-font-family', fontFamily);
    document.documentElement.style.setProperty('--main-bg-color', backgroundColor);
    document.documentElement.style.setProperty('--main-bg-color2', backgroundColor2);
    document.documentElement.style.setProperty('--main-input-color', inputColor);
    document.documentElement.style.setProperty('--main-button-color', buttonColor);
  }

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
        font-size: var(--main-font-size);
        color: var(--main-color);
        margin-bottom: 1em;
        font-family: var(--main-font-family);
    }

    .form {
        padding: 20px;
        background-color: var(--main-bg-color2);
        color: var(--main-color);
        border-radius: 4px;
    }

    .form-select {
        background-color: var(--main-input-color);
        color: var(--main-color);
    }

    .form-title {
        font-size: 24px;
        margin-bottom: 10px;
    }

    .form-label {
        display: block;
        margin-bottom: 10px;
        color: var(--main-color);
    }

    .form-input {
        width: 100%;
        padding: 8px;
        border: none;
        border-radius: 4px;
        background: var(--main-input-color);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        color: var(--main-color);
        flex-grow: 1;
    }

    .form-note {
        font-size: 14px;
        margin-bottom: 10px;
        color: var(--main-color);
    }

    .form-button {
        display: inline-block;
        padding: 8px 16px;
        background: var(--main-button-color);
        color: var(--main-color);
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
        transition: background-color 0.2s;
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
        background: var(--main-button-color);
        border-radius: 5px;
        font-size: 16px;
        color: var(--main-color);
        transition: background-color 0.3s;
    }

</style>
