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
    import deleteIcon from "../assets/images/delete.png";
    import { checkCommand } from "../functions/functions";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
    import infoIcon from "../assets/images/info.png";
    import { settings } from '../stores';

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    let showInfoButton: boolean;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
        showInfoButton = $settings.showInfoButton;
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

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
            <p>
                This page allows you to <strong>manage cron jobs</strong>.
                <br />
                <br />
                You can <em>add a cron job</em> by clicking the "Add Cron Job" button.
                <br />
                <br />
                You can <em>delete a cron job</em> by clicking the "Delete" button next to the cron job.
                <br />
                <br />
                You can <em>delete all cron jobs</em> by clicking the "Remove all cron jobs" button.
                <br />
                <br />
                Cron jobs use a specific syntax for scheduling tasks. The syntax is made up of <strong>five fields</strong>:
                <br />
                <br />
                - <strong>Minute</strong> (0 - 59)
                <br />
                - <strong>Hour</strong> (0 - 23, where 0 is midnight)
                <br />
                - <strong>Day of the month</strong> (1 - 31)
                <br />
                - <strong>Month</strong> (1 - 12)
                <br />
                - <strong>Day of the week</strong> (0 - 7, where both 0 and 7 are Sunday)
                <br />
                <br />
                Each field is separated by a space and can contain a <em>single value</em>, a <em>range of values</em>, or an <em>asterisk</em> to indicate "any value".
                <br />
                <br />
                Here are some examples:
                <br />
                <br />
                - "<strong>5 0 * * *</strong>" runs a job at 0:05 every day
                <br />
                - "<strong>15 14 1 * *</strong>" runs a job at 14:15 on the first day of every month
                <br />
                - "<strong>* * * * 1</strong>" runs a job every minute on Mondays
                <br />
                <br />
                Please refer to the <em>cron syntax guide</em> for more detailed information.
                (https://www.baeldung.com/cron-expressions)
            `
        );
    }

    function onDialogClose() {
        dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    }
</script>

<CustomDialog
    bind:show={dialog.showDialog}
    title="Info"
    message={dialog.dialogMessage}
    onClose={onDialogClose}
/>

<div class="container">
    {#if showInfoButton}
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    {/if}
    <h1>Cron Jobs</h1>

    {#if viewState === "default"}
        <button on:click={() => setViewState("addCronJob")}>Add Cron Job</button
        >
        <button on:click={removeAllCronJobs}>Remove all cron jobs</button>

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
                            on:click={() => deleteCron(job.Command)}
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
    .container {
        display: flex;
        flex-direction: column;
        max-width: 75%;
        margin: 2rem auto;
        background-color: var(--main-bg-color2);
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
        border-radius: 10px;
        color: #ddd;
        padding: 1rem;
    }

    .info-button {
        position: absolute;
        top: 10px;
        right: 10px;
        border: none;
        background: var(--main-button-color);
        height: 40px;
        width: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 50%;
    }

    .info-icon {
        max-width: none;
    }

    h1,
    h2,
    p {
        font-size: var(--main-font-size);
        color: var(--main-color);
        margin-bottom: 1em;
        font-family: var(--main-font-family);
    }

    ul {
        list-style: none;
        padding: 0;
        border: 1px solid #333;
        border-radius: 4px;
        margin-bottom: 1em;
        background-color: var(--main-bg-color);
    }

    li {
        padding: 0.8em;
        display: flex;
        justify-content: space-between;
        align-items: center;
        border-bottom: 1px solid #444;
    }

    li:last-child {
        border-bottom: none;
    }

    button {
        padding: 0.8em 1em;
        border: none;
        border-radius: 4px;
        background: var(--main-button-color);
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
        width: 100%;
        margin-bottom: 1em;
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
