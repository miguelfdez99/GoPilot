<script lang="ts">
    import { onMount } from "svelte";
    import {
        GenerateKeys,
        GenerateSelfSignedCertificate,
        OpenDir,
        OpenDialogInfo,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";
    import CustomDialog from "../components/dialogs/CustomDialog.svelte";

    let keyType: string = "";
    let keyName: string = "";
    let outputPath: string = "";
    let certName: string = "";
    let overwrite: boolean = false;
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    async function generateKeys() {
        if (!keyType || !keyName || !outputPath) {
            await OpenDialogError("All fields are required");
            return;
        }

        try {
            await GenerateKeys(keyType, keyName, outputPath, overwrite);
            await OpenDialogInfo("Keys generated successfully");
            resetKeyForm();
        } catch (err) {
            await OpenDialogError(`Error generating keys: ${err}`);
        }
    }

    async function generateCertificate() {
        if (!outputPath || !certName) {
            await OpenDialogError("All fields are required");
            return;
        }

        try {
            const fullPath = `${outputPath}/${certName}`;
            await GenerateSelfSignedCertificate(fullPath);
            await OpenDialogInfo(
                "Self-signed certificate generated successfully",
            );
            resetCertForm();
        } catch (err) {
            await OpenDialogError(`Error generating certificate: ${err}`);
        }
    }

    const selectDir = async () => {
        const dir = await OpenDir();
        if (dir) {
            outputPath = dir;
        }
    };

    function resetKeyForm() {
        keyType = "";
        keyName = "";
        outputPath = "";
        overwrite = false;
    }

    function resetCertForm() {
        outputPath = "";
        certName = "";
    }

    function onDialogClose() {
        dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };
    }
</script>

<CustomDialog
    bind:show={dialog.showDialog}
    title={dialog.dialogTitle}
    message={dialog.dialogMessage}
    onClose={onDialogClose}
/>

<div class="main-container">
    <h2>Certificates & Keys</h2>

    <div class="form-container">
        <h3>Generate Keys</h3>
        <form on:submit|preventDefault={generateKeys}>
            <label class="input-field">
                <span>Key Type <span class="required">*</span>:</span>
                <select bind:value={keyType}>
                    <option value="">-- Select Key Type --</option>
                    <option value="rsa">RSA</option>
                    <option value="dsa">DSA</option>
                    <option value="ecdsa">ECDSA</option>
                    <option value="ed25519">Ed25519</option>
                </select>
            </label>

            <label class="input-field">
                <span>Key Name <span class="required">*</span>:</span>
                <input
                    type="text"
                    bind:value={keyName}
                    placeholder="Enter key name"
                />
            </label>

            <label class="input-field">
                <span>Output Path <span class="required">*</span>:</span>
                <div class="input-group">
                    <input
                        type="text"
                        bind:value={outputPath}
                        placeholder="Select output directory"
                        readonly
                    />
                    <button
                        type="button"
                        class="open-btn"
                        on:click={selectDir}
                        title="Select Output Directory"
                    >
                        <span class="material-icons">folder_open</span>
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
        <h3>Generate Self-Signed Certificate</h3>
        <form on:submit|preventDefault={generateCertificate}>
            <label class="input-field">
                <span
                    >Certificate Destination Path <span class="required">*</span
                    >:</span
                >
                <div class="input-group">
                    <input
                        type="text"
                        bind:value={outputPath}
                        placeholder="Select certificate directory"
                        readonly
                    />
                    <button
                        type="button"
                        class="open-btn"
                        on:click={selectDir}
                        title="Select Certificate Directory"
                    >
                        <span class="material-icons">folder_open</span>
                    </button>
                </div>
            </label>

            <label class="input-field">
                <span>Cert Name <span class="required">*</span>:</span>
                <input
                    type="text"
                    bind:value={certName}
                    placeholder="Enter certificate name"
                />
            </label>

            <button class="submit-button" type="submit">
                Generate Self-Signed Certificate
            </button>
        </form>
    </div>
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

    button {
        background-color: var(--color-accent-blue);
        color: var(--color-bg-primary);
        border: none;
        padding: var(--spacing-sm) var(--spacing-md);
        border-radius: 0.5rem;
        cursor: pointer;
        font-size: 1rem;
        transition: background-color 0.2s ease-in-out;
    }

    button:hover {
        background-color: var(--color-accent-cyan);
    }

    .main-container {
        padding: var(--spacing-lg);
        background-color: var(--color-bg-primary);
        color: var(--color-text-primary);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        max-width: 800px;
        margin: 0 auto;
    }

    h2,
    h3 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
    }

    h2 {
        font-size: 1.5rem;
        text-align: center;
    }

    h3 {
        font-size: 1.25rem;
    }

    .form-container {
        margin-bottom: var(--spacing-lg);
        padding: var(--spacing-lg);
        background-color: var(--color-bg-secondary);
        border-radius: 0.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .input-field {
        margin-bottom: var(--spacing-md);
    }

    .input-field > span {
        display: block;
        margin-bottom: var(--spacing-sm);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    .input-field > input,
    .input-field > select {
        width: 100%;
        padding: var(--spacing-sm);
        border: 1px solid var(--color-bg-tertiary);
        border-radius: 0.5rem;
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        font-size: 1rem;
    }

    .input-group {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
    }

    .open-btn {
        background: none;
        border: none;
        padding: var(--spacing-sm);
        color: var(--color-text-primary);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .open-btn:hover {
        color: var(--color-accent-blue);
    }

    .material-icons {
        font-size: 1.5rem;
    }

    .required {
        color: var(--color-accent-red);
        font-size: 0.9rem;
        margin-left: 0.25rem;
    }

    .submit-button {
        width: 100%;
        margin-top: var(--spacing-md);
    }

    hr {
        border: 0;
        height: 1px;
        background-color: var(--color-bg-tertiary);
        margin: var(--spacing-lg) 0;
    }
</style>
