<script lang="ts">
    import { onMount } from "svelte";
    import {
        GenerateKeys,
        GenerateSelfSignedCertificate,
        OpenFile,
        OpenDir,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
    import { openDialog } from "../functions/functions";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";
    import openIcon from "../assets/images/open.png";
    import infoIcon from "../assets/images/info.png";
    import { settings } from "../stores";

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
        document.documentElement.style.setProperty(
            "--main-font-size",
            fontSize
        );
        document.documentElement.style.setProperty("--main-color", color);
        document.documentElement.style.setProperty(
            "--main-font-family",
            fontFamily
        );
        document.documentElement.style.setProperty(
            "--main-bg-color",
            backgroundColor
        );
        document.documentElement.style.setProperty(
            "--main-bg-color2",
            backgroundColor2
        );
        document.documentElement.style.setProperty(
            "--main-input-color",
            inputColor
        );
        document.documentElement.style.setProperty(
            "--main-button-color",
            buttonColor
        );
    }

    let keyType = "";
    let keyName = "";
    let outputPath = "";
    let overwrite = false;
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    let privateKeyPath = "";

    async function generateKeys() {
        if (!keyType || !keyName || !outputPath) {
            await OpenDialogError("All fields are required");
            return;
        }

        try {
            await GenerateKeys(keyType, keyName, outputPath, overwrite);
            await OpenDialogInfo("Successfully created keys");
            keyType = "";
            keyName = "";
            outputPath = "";
            overwrite = false;
        } catch (err) {
            await OpenDialogError(`Error: ${err}`);
        }
    }

    async function generateCertificate() {
        if (!privateKeyPath) {
            await OpenDialogError("Private key path is required");
            return;
        }

        try {
            await GenerateSelfSignedCertificate(privateKeyPath);
            await OpenDialogInfo(
                "Successfully created self-signed certificate"
            );
            privateKeyPath = "";
        } catch (err) {
            await OpenDialogError(`Error: ${err}`);
        }
    }

    const selectFile = async () => {
        const file = await OpenFile();
        privateKeyPath = file;
    };

    const selectDir = async () => {
        const dir = await OpenDir();
        outputPath = dir;
    };

    function openInfo() {
        dialog = openDialog(
        dialog,
        "Info",
        `
        <div style="color: ${color};">
            <p style="color: ${color};>This component generates cryptographic keys and self-signed certificates. 
            It supports key types RSA, DSA, ECDSA, and Ed25519. 'Generate Keys' lets users define a key type, name, and output path. Existing keys at the same path can be overwritten if desired.</p>
            
            <p style="color: ${color};> 
                - <strong>RSA:</strong> Public key cryptography based on factoring problem. <a href="https://simple.wikipedia.org/wiki/RSA_algorithm" style="color: inherit;">RSA Algorithm</a><br>
                - <strong>DSA:</strong> A digital signature standard using modular exponentiation. <a href="https://en.wikipedia.org/wiki/Digital_Signature_Algorithm" style="color: inherit;">Digital Signature Algorithm</a><br>
                - <strong>ECDSA:</strong> A DSA variant utilizing elliptic curve cryptography. <a href="https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm" style="color: inherit;">Elliptic Curve DSA</a><br>
                - <strong>Ed25519:</strong> A public-key signature system known for its speed and small signature size. <a href="https://ed25519.cr.yp.to" style="color: inherit;">Ed25519</a>
            </p>
            
            <p style="color: ${color};>'Generate Self-Signed Certificate' requires a private key path to create a self-signed certificate.</p>
            
            <p style="color: ${color};>This component utilizes 'ssh-keygen' to generate keys. <a href="https://linux.die.net/man/1/ssh-keygen" style="color: inherit;">ssh-keygen</a></p>
        </div>
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

<div class="main-container">
    {#if showInfoButton}
        <button
            type="button"
            class="info-button"
            title="Info"
            on:click={openInfo}
        >
            <img src={infoIcon} alt="Open Info" class="info-icon" />
        </button>
    {/if}
    <h2>Certificates & Keys</h2>
    <div class="form-container">
        <h2>Generate Keys</h2>
        <form on:submit|preventDefault={generateKeys}>
            <label class="input-field">
                <span>Key Type:</span>
                <select bind:value={keyType}>
                    <option value="">--Please choose an option--</option>
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
                <div class="input-group">
                    <input type="text" bind:value={outputPath} />
                    <button
                        type="button"
                        class="open-btn"
                        title="Select Output Directory"
                        on:click={selectDir}
                    >
                        <img src={openIcon} alt="Open Dir" class="open-icon" />
                    </button>
                </div>
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

    <hr />

    <div class="form-container">
        <h2>Generate Self-Signed Certificate</h2>
        <form on:submit|preventDefault={generateCertificate}>
            <label class="input-field">
                <span>Private Key Path:</span>
                <div class="input-group">
                    <input type="text" bind:value={privateKeyPath} />
                    <button
                        type="button"
                        class="open-btn"
                        title="Select Private Key"
                        on:click={selectFile}
                    >
                        <img src={openIcon} alt="Open File" class="open-icon" />
                    </button>
                </div>
            </label>

            <button class="submit-button" type="submit"
                >Generate Self-Signed Certificate</button
            >
        </form>
    </div>
</div>

<style>
    .main-container {
        background-color: var(--main-bg-color);
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

    .form-container {
        width: 90%;
        max-width: 600px;
        margin: 3rem auto;
        padding: 2rem;
        background-color: var(--main-bg-color2);
        box-shadow: 0 0 15px rgba(0, 0, 0, 0.3);
        border-radius: 5px;
    }

    .input-field,
    .checkbox-field {
        margin-bottom: 1.5rem;
    }

    .input-field > span,
    .checkbox-field > span {
        font-weight: 500;
        color: var(--main-color);
        margin-bottom: 0.5rem;
    }

    .input-field > input,
    .input-field > select {
        padding: 0.7em;
        border: 0;
        border-radius: 4px;
        background: var(--main-input-color);
        color: var(--main-color);
    }

    .checkbox-field {
        display: flex;
        align-items: center;
        color: var(--main-color);
    }

    .submit-button {
        padding: 0.8em 1em;
        border: none;
        border-radius: 4px;
        background: var(--main-button-color);
        color: #fff;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    button.open-btn {
        background: none;
        width: 2.5rem;
        height: 2.5rem;
        padding: 0;
        margin: 0;
        margin-left: 10px;
        margin-bottom: 15px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-color: #333;
    }

    .input-group {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 1em;
    }

    .input-group input {
        margin-right: 0.5em;
        background: var(--main-input-color);
        color: var(--main-color);
    }

    hr {
        border: 0;
        height: 1px;
        background: #444;
    }

    h2 {
        color: var(--main-color);
    }
</style>
