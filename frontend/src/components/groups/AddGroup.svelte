<script lang="ts">
    import {
        CreateGroup,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";
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

    let name: string = "";
    let gid: string = "";

    function dismiss() {
        dispatch("dismiss");
    }

    async function createGroup() {
        if (!name) {
            await OpenDialogError("Please enter a group name");
            return;
        }

        const group = {
            name,
            gid: gid ? parseInt(gid) : undefined,
        };

        try {
            await CreateGroup(group.name, group.gid);
            await OpenDialogInfo(`Successfully created group ${name}`);
            name = "";
            gid = "";
        } catch (err) {
            await OpenDialogError(`Failed to create group ${name}: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Add Group</h2>
    <form on:submit|preventDefault={createGroup} class="form-control">
        <label class="input-field">
            <span>Group name:</span>
            <input type="text" bind:value={name} required />
        </label>
        <label class="input-field">
            <span>Group ID (GID):</span>
            <input type="text" bind:value={gid} />
        </label>
        <button type="submit" class="submit-button">Create group</button>
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

    .input-field span {
      display: block;
      font-size: 14px;
      color: var(--main-color);
      margin-bottom: 5px;
    }

    input[type="text"] {
      width: 100%;
      border: none;
      padding: 10px;
      border-radius: 5px;
      font-size: 16px;
      background: var(--main-input-color);
      color: var(--main-color);
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
