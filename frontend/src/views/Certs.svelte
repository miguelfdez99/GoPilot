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
  
  <div class="form-container">
      <h2>Generate Keys</h2>
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
          <select bind:value={overwrite}>
              <option value={true}>Yes</option>
              <option value={false}>No</option>
          </select>
      </label>
  
      <button class="submit-button" type="submit">Generate Keys</button>
  </form>
  </div>
  
  <hr>
  
  <div class="form-container">
      <h2>Generate Self-Signed Certificate</h2>
  <form on:submit|preventDefault={generateCertificate}>
    <label class="input-field">
        <span>Private Key Path:</span>
        <input type="text" bind:value={privateKeyPath} />
    </label>
  
    <button class="submit-button" type="submit">Generate Self-Signed Certificate</button>
  </form>
  </div>
  
  <style>
      .form-container {
          width: 90%;
          max-width: 600px;
          margin: 3rem auto;
          padding: 2rem;
          background: #282828;
          box-shadow: 0 0 15px rgba(0,0,0,0.3);
          border-radius: 5px;
      }
  
      .input-field, .checkbox-field {
          margin-bottom: 1.5rem;
      }
  
      .input-field > span, .checkbox-field > span {
          font-weight: 500;
          color: #ddd;
          margin-bottom: .5rem;
      }
  
      .input-field > input, .input-field > select {
          padding: .7em;
          border: 0;
          border-radius: 4px;
          background: #333;
          color: #fff;
      }
  
      .checkbox-field {
          display: flex;
          align-items: center;
          color: #ddd;
      }
  
      .submit-button {
          padding: .8em 1em;
          border: none;
          border-radius: 4px;
          background: #1abc9c;
          color: #fff;
          cursor: pointer;
          transition: background-color 0.3s;
      }
  
      .submit-button:hover {
          background-color: #16a085;
      }
  
      hr {
          border: 0;
          height: 1px;
          background: #444;
      }
  
      h2 {
          color: white;
      }
  </style>