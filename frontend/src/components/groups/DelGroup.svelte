<script lang="ts">
    import {
        DeleteGroup,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../../wailsjs/go/backend/Backend";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let name: string = "";

    function dismiss() {
        dispatch("dismiss");
    }

    async function deleteGroup() {
        if (!name) {
            await OpenDialogError("Please enter a group name");
            return;
        }

        try {
            await DeleteGroup(name);
            await OpenDialogInfo(`Successfully deleted group: ${name}`);
            name = "";
        } catch (err) {
            await OpenDialogError(`Failed to delete group: ${err}`);
        }
    }
</script>

<div class="container">
    <h2>Delete Group</h2>
    <form on:submit|preventDefault={deleteGroup} class="form-control">
        <label class="input-field">
            <span>Group name:</span>
            <input type="text" bind:value={name} required />
        </label>
        <button type="submit" class="submit-button">Delete group</button>
        <button class="back-button" on:click={dismiss}>Back</button>
    </form>
</div>

<style>
    h2 {
      text-align: center;
      color: #fff;
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
      background-color: #333;
      border-radius: 5px;
      font-size: 16px;
      color: #fff;
      transition: background-color 0.3s;
    }

    .back-button:hover {
      background-color: #555;
    }

    .form-control {
      padding: 30px;
      background-color: #222;
      border-radius: 5px;
    }

    .input-field {
      margin-bottom: 20px;
    }

    .input-field span {
      display: block;
      font-size: 14px;
      color: #ccc;
      margin-bottom: 5px;
    }

    input[type="text"] {
      width: 100%;
      border: none;
      padding: 10px;
      border-radius: 5px;
      font-size: 16px;
      background-color: #333;
      color: #fff;
    }

    .submit-button {
      display: block;
      width: 100%;
      padding: 10px;
      border: none;
      background: #c0392b;
      color: white;
      border-radius: 5px;
      font-size: 16px;
      cursor: pointer;
      transition: background-color 0.3s;
    }

    .submit-button:hover {
        background-color: #e74c3c;
    }
</style>
