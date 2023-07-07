<script lang="ts">
  import { onMount } from "svelte";
  import { GenerateKeys, GenerateSelfSignedCertificate } from "../../wailsjs/go/backend/Backend";
  import { openDialog, closeDialog } from "../functions/functions";
  import CustomDialog from "../components/dialogs/CustomDialog.svelte";

  let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

  let keyType = '';
  let keyName = '';
  let outputPath = '';
  let overwrite = false;

  let privateKeyPath = '';

  async function generateKeys() {
      if (!keyType || !keyName || !outputPath) {
          dialog = openDialog(dialog, "Error", "All fields are required");
          return;
      }

      try {
          await GenerateKeys(keyType, keyName, outputPath, overwrite);
          dialog = openDialog(
              dialog,
              "Success",
              `Successfully created keys`
          );
          keyType = "";
          keyName = "";
          outputPath = "";
          overwrite = false;
      } catch (err) {
          dialog = openDialog(
              dialog,
              "Error",
              `${err}`
          );
      }
  }

  async function generateCertificate() {
        if (!privateKeyPath) {
            dialog = openDialog(dialog, "Error", "Private key path is required");
            return;
        }

        try {
            await GenerateSelfSignedCertificate(privateKeyPath);
            dialog = openDialog(
                dialog,
                "Success",
                `Successfully created self-signed certificate`
            );
            privateKeyPath = "";
        } catch (err) {
            dialog = openDialog(
                dialog,
                "Error",
                `${err}`
            );
        }
    }
</script>

<CustomDialog
    bind:show={dialog.showDialog}
    title={dialog.dialogTitle}
    message={dialog.dialogMessage}
    onClose={() => (dialog = closeDialog(dialog))}
    confirmButton={false}
/>

<form on:submit|preventDefault={generateKeys}>
    <label class="input-field">
        <span>Key Type:</span>
        <select bind:value={keyType}>
            <option value=''>--Please choose an option--</option>
            <option>rsa</option>
            <option>dsa</option>
            <option>ecdsa</option>
            <option>ed25519</option>
        </select>
    </label>

    <label class="input-field">
        <span>Key Name:</span>
        <input type="text" bind:value={keyName} />
    </label>

    <label class="input-field">
        <span>Output Path:</span>
        <input type="text" bind:value={outputPath} />
    </label>

    <label class="input-field checkbox-field">
        <span>Overwrite existing keys?</span>
        <input type="checkbox" bind:checked={overwrite} />
    </label>

    <button class="submit-button" type="submit">Generate Keys</button>
</form>

<hr>

<form on:submit|preventDefault={generateCertificate}>
  <label class="input-field">
      <span>Private Key Path:</span>
      <input type="text" bind:value={privateKeyPath} />
  </label>

  <button class="submit-button" type="submit">Generate Self-Signed Certificate</button>
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        max-width: 400px;
        margin: 0 auto;
        padding: 20px;
    }

    .input-field {
        display: flex;
        flex-direction: column;
        margin-bottom: 15px;
    }

    .input-field span {
        font-size: 14px;
        color: #555;
        margin-bottom: 5px;
        color: white;
    }

    .input-field input,
    .input-field select {
        padding: 10px;
        font-size: 16px;
        border: 1px solid #ccc;
        border-radius: 5px;
        color: #333;
        outline: none;
        color: white;
    }

    .checkbox-field {
        align-items: center;
    }

    .checkbox-field input {
        width: auto;
        margin-left: 5px;
    }

    .submit-button {
        padding: 10px 20px;
        font-size: 16px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .submit-button:hover {
        background-color: #0056b3;
    }
</style>
