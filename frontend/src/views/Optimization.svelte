<script lang="ts">
    import {
        DeleteTempFiles,
        RemoveUnusedPackages,
        CleanCachePackages,
        DuplicatedFiles,
        RemoveOldLogs,
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

    let tempFilesDirPath: string = "";
    let expireDays: number;
    let duplicatedFilesDirPath: string = "";
    let oldLogsDays: string = "";
    let dialog = { showDialog: false, dialogTitle: "", dialogMessage: "" };

    const startDeletingTempFiles = () => {
        if (!tempFilesDirPath || !expireDays) {
            OpenDialogError(
                "Both directory path and number of days are required."
            );
            return;
        }
        try {
            DeleteTempFiles(tempFilesDirPath, expireDays);
            OpenDialogInfo("Temp files removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing temp files: ${err.message}`);
        }
    };

    const startRemovingUnusedPackages = () => {
        try {
            RemoveUnusedPackages();
            OpenDialogInfo("Unused packages removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing unused packages: ${err.message}`);
        }
    };

    const startCleaningCachePackages = () => {
        try {
            CleanCachePackages();
            OpenDialogInfo("Cache packages cleaned successfully");
        } catch (err) {
            OpenDialogError(`Error cleaning cache packages: ${err.message}`);
        }
    };

    const startDeletingDuplicatedFiles = () => {
        if (!duplicatedFilesDirPath) {
            OpenDialogError("Directory path is required.");
            return;
        }
        try {
            DuplicatedFiles(duplicatedFilesDirPath);
            OpenDialogInfo("Duplicated files removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing duplicated files: ${err.message}`);
        }
    };

    const startRemovingOldLogs = () => {
        if (!oldLogsDays) {
            OpenDialogError("Number of days is required.");
            return;
        }
        try {
            RemoveOldLogs(oldLogsDays);
            OpenDialogInfo("Old logs removed successfully");
        } catch (err) {
            OpenDialogError(`Error removing old logs: ${err.message}`);
        }
    };

    const selectDir = async () => {
        const dir = await OpenDir();
        if (dir) {
            tempFilesDirPath = dir;
        }
    };

    const selectPath = async () => {
        const dir = await OpenDir();
        if (dir) {
            duplicatedFilesDirPath = dir;
        }
    };

    function openInfo() {
        dialog = openDialog(
            dialog,
            "Info",
            `
            <p style="color: ${color}; font-size: ${fontSize};>
                This section allows you to clean the system and packages. You can do the following:
            <p style="color: ${color}; font-size: ${fontSize};>
                <b>Remove Unused Packages:</b> Remove unused packages from the system.
            </p>
            <p style="color: ${color}; font-size: ${fontSize};>
                <b>Clean Cache Packages:</b> Clean cache packages from the system.
            </p>
            <p style="color: ${color}; font-size: ${fontSize};>
                <b>Delete Temporary Files:</b> Delete temporary files from the system.
            </p>
            <p style="color: ${color}; font-size: ${fontSize};>
                <b>Delete Duplicated Files:</b> Delete duplicated files from the system.
            </p>
            <p style="color: ${color}; font-size: ${fontSize};>
                <b>Remove Old Logs:</b> Remove old logs from the system.
            </p>
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

<div class="clean-system-container">
    {#if showInfoButton}
    <button type="button" class="info-button" title="Info" on:click={openInfo}>
        <img src={infoIcon} alt="Open Info" class="info-icon" />
    </button>
    {/if}
    <div class="clean-system-section">
        <h2>Clean System</h2>

        <div class="input-group">
            <div class="flex-input">
                <input
                    type="text"
                    bind:value={tempFilesDirPath}
                    placeholder="Enter a directory path for temporary files"
                />
                <button
                    class="open-btn"
                    title="Select Path"
                    on:click={selectDir}
                >
                    <img src={openIcon} alt="Open Dir" class="open-icon" />
                </button>
            </div>
            <input
                type="number"
                bind:value={expireDays}
                placeholder="Enter the number of days to keep temporary files"
            />
            <button on:click={startDeletingTempFiles}
                >Delete Temporary Files</button
            >
        </div>

        <div class="input-group">
            <div class="flex-input">
                <input
                    type="text"
                    bind:value={duplicatedFilesDirPath}
                    placeholder="Enter a directory path to find duplicated files"
                />
                <button
                    class="open-btn"
                    title="Select Path"
                    on:click={selectPath}
                >
                    <img src={openIcon} alt="Open Dir" class="open-icon" />
                </button>
            </div>
            <button on:click={startDeletingDuplicatedFiles}
                >Delete Duplicated Files</button
            >
        </div>

        <div class="input-group">
            <input
                type="text"
                bind:value={oldLogsDays}
                placeholder="Enter the number of days to keep old logs"
            />
            <button on:click={startRemovingOldLogs}>Remove Old Logs</button>
        </div>
    </div>

    <div class="clean-system-section">
        <h2>Package Cleaning</h2>

        <button on:click={startRemovingUnusedPackages}
            >Remove Unused Packages</button
        >
        <button on:click={startCleaningCachePackages}
            >Clean Cache Packages</button
        >
    </div>
</div>

<style>
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
    .clean-system-container {
        display: flex;
        justify-content: space-between;
        max-width: 95%;
        margin: 2rem auto;
    }

    .clean-system-section {
        flex: 1;
        padding: 1rem;
        background-color: var(--main-bg-color2);
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.2);
        border-radius: 10px;
        color: #ddd;
        margin: 1rem;
    }

    .flex-input {
        display: flex;
        align-items: center;
    }

    .input-group {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        margin-bottom: 2rem;
    }

    h2 {
        color: var(--main-color);
        font-family: var(--main-font-family);
    }
    input {
        margin-bottom: 10px;
        padding: 10px;
        width: 100%;
        box-sizing: border-box;
        border: 1px solid var(--main-input-color);;
        background: var(--main-input-color);
        color: var(--main-color);
        font-size: var(--main-font-size);
        font-family: var(--main-font-family);
        border-radius: 5px;
    }
    button {
        margin-bottom: 10px;
        padding: 10px 20px;
        border: none;
        background: var(--main-button-color);
        color: white;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s ease;
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
    }

    button.open-btn:hover {
        background-color: #131313;
    }
</style>
