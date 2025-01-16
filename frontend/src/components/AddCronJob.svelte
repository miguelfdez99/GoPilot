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

<div class="add-cron-container">
    <form class="form" on:submit|preventDefault={onSubmit}>
        <h2 class="form-title">Add a New Cron Job</h2>

        <!-- Quick Schedule Dropdown -->
        <div class="input-container">
            <span class="input-indicator">Quick Schedule:</span>
            <select class="form-select" bind:value={quickScheduleOption}>
                <option value="">None</option>
                <option value="@reboot">Startup</option>
                <option value="@hourly">Hourly</option>
                <option value="@daily">Daily</option>
                <option value="@weekly">Weekly</option>
                <option value="@monthly">Monthly</option>
                <option value="@yearly">Yearly</option>
            </select>
        </div>

        <!-- Manual Schedule Inputs -->
        {#if !quickScheduleOption}
            <div class="input-container">
                <span class="input-indicator">Minutes:</span>
                <input class="form-input" type="text" bind:value={minute} />
            </div>
            <div class="input-container">
                <span class="input-indicator">Hour:</span>
                <input class="form-input" type="text" bind:value={hour} />
            </div>
            <div class="input-container">
                <span class="input-indicator">Day:</span>
                <input class="form-input" type="text" bind:value={day} />
            </div>
            <div class="input-container">
                <span class="input-indicator">Month:</span>
                <input class="form-input" type="text" bind:value={month} />
            </div>
            <div class="input-container">
                <span class="input-indicator">Day of Week:</span>
                <input class="form-input" type="text" bind:value={dayOfWeek} />
            </div>
        {/if}

        <!-- Command Input -->
        <div class="input-container">
            <span class="input-indicator">Command:</span>
            <input class="form-input" type="text" bind:value={command} />
        </div>

        <!-- Composed Job Preview -->
        <div class="input-container">
            <span class="input-indicator">Job Preview:</span>
            <input class="form-input" type="text" bind:value={composedJob} readonly />
        </div>

        <!-- Form Notes -->
        <p class="form-note">
            Enter cron job schedule in the format: <code>minute hour day month day-of-week</code>
        </p>
        <p class="form-note">
            For example: <code>30 * * * *</code> for every hour at 30 minutes past the hour.
        </p>
        <p class="form-note">
            Or: <code>0 0 * * *</code> for midnight every day.
        </p>
        <p class="form-note">
            Refer to cron syntax for more information on how to specify cron job schedules.
        </p>
        <p class="form-note">
            Or select a quick schedule from the dropdown above.
        </p>

        <!-- Form Buttons -->
        <div class="form-buttons">
            <button class="form-button" type="submit">
                <span class="material-icons">add</span>
                Add Job
            </button>
            <button class="back-button" on:click={dismiss}>
                <span class="material-icons">arrow_back</span>
                Back
            </button>
        </div>
    </form>
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

    .add-cron-container {
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        font-family: sans-serif;
    }

    .form-title {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .input-container {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-md);
    }

    .input-indicator {
        flex-basis: 30%;
        color: var(--color-text-primary);
        font-weight: bold;
    }

    .form-input,
    .form-select {
        flex: 1;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-family: monospace;
    }

    .form-note {
        color: var(--color-text-secondary);
        font-size: 0.9rem;
        margin-bottom: var(--spacing-sm);
    }

    .form-buttons {
        display: flex;
        gap: var(--spacing-md);
        margin-top: var(--spacing-lg);
    }

    .form-button,
    .back-button {
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

    .form-button:hover,
    .back-button:hover {
        background-color: var(--color-bg-tertiary);
    }

    .material-icons {
        font-size: 1.2rem;
        color: var(--color-accent-blue);
    }
</style>
