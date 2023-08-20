<script lang="ts">
    import {
        DeleteUser,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from 'svelte';
    import { settings } from '../../stores';

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

    let username: string = "";
    let removeHomeDir: boolean = false;
    let forceDelete: boolean = false;

    function dismiss() {
        dispatch('dismiss');
    }

    async function deleteUser() {
        if (!username) {
            await OpenDialogError("Please enter a username");
            return;
        }

        try {
            await DeleteUser(username, removeHomeDir, forceDelete);
            await OpenDialogInfo(`Successfully deleted user ${username}`);
            username = "";
        } catch (err) {
            await OpenDialogError(`Failed to delete user ${username}: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Delete User</h2>
    <form on:submit|preventDefault={deleteUser} class="form-control">
        <label class="input-field">
            <span>Username:</span>
            <input type="text" bind:value={username} required />
        </label>
        <label class="input-field checkbox-field">
            <span>Remove home directory:</span>
            <input type="checkbox" bind:checked={removeHomeDir} />
        </label>
        <label class="input-field checkbox-field">
            <span>Force delete:</span>
            <input type="checkbox" bind:checked={forceDelete} />
        </label>
        <button type="submit" class="submit-button">Delete User</button>
        <button class="back-button" on:click={dismiss}>Back</button>
    </form>
</div>

<style>
    h2 {
      text-align: center;
      color: var(--main-color);
    }
    
    .container {
        position: relative;
        padding: 20px;
        max-width: 400px;
        margin: 0 auto;
    }

    .back-button {
        right: 0;
        padding: 10px 20px;
        border: none;
        background: var(--main-button-color);
        border-radius: 5px;
        font-size: 16px;
        color: #fff;
        transition: background-color 0.3s;
    }

    .form-control {
        padding: 30px;
        background-color: var(--main-bg-color2);
        border-radius: 5px;
    }

    .input-field {
        margin-bottom: 20px;
    }

    .input-field.checkbox-field {
        display: flex;
        align-items: center;
    }

    .input-field span, .input-field.checkbox-field span {
        font-size: 14px;
        color: var(--main-color);
    }

    input[type="text"], input[type="checkbox"] {
        border: none;
        padding: 10px;
        border-radius: 5px;
        font-size: 16px;
        background: var(--main-input-color);
        color: var(--main-color);
    }

    input[type="checkbox"] {
        width: auto;
        margin-left: 15px;
    }

    .submit-button {
        display: block;
        width: 100%;
        padding: 10px;
        border: none;
        background: var(--main-button-color);
        color: white;
        border-radius: 5px;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.3s;
    }

</style>
