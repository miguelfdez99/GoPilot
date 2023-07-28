<script lang="ts">
    import { DeleteTempFiles, RemoveUnusedPackages, CleanCachePackages, DuplicatedFiles, RemoveOldLogs, OpenDir } from "../../wailsjs/go/backend/Backend";
    import openIcon from "../assets/images/open.png";

    let tempFilesDirPath: string = "";
    let expireDays: number;
    let duplicatedFilesDirPath: string = "";
    let oldLogsDays: string = "";

    const startDeletingTempFiles = () => {
        DeleteTempFiles(tempFilesDirPath, expireDays);
    };

    const startRemovingUnusedPackages = () => {
        RemoveUnusedPackages();
    };

    const startCleaningCachePackages = () => {
        CleanCachePackages();
    };

    const startDeletingDuplicatedFiles = () => {
        DuplicatedFiles(duplicatedFilesDirPath);
    };

    const startRemovingOldLogs = () => {
        RemoveOldLogs(oldLogsDays);
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

</script>

<div class="clean-system-container">

    <div class="clean-system-section">
        <h2>Clean System</h2>

        <div class="input-group">
            <div class="flex-input">
                <input type="text" bind:value={tempFilesDirPath} placeholder="Enter a directory path for temporary files" />
                <button class="open-btn" title="Select Path" on:click={selectDir}>
                    <img src={openIcon} alt="Open Dir" class="open-icon" />
                </button>
            </div>
            <input type="number" bind:value={expireDays} placeholder="Enter the number of days to keep temporary files" />
            <button on:click={startDeletingTempFiles}>Delete Temporary Files</button>
        </div>
        
        
        <div class="input-group">
            <div class="flex-input">
            <input type="text" bind:value={duplicatedFilesDirPath} placeholder="Enter a directory path to find duplicated files" />
            <button class="open-btn" title="Select Path" on:click={selectPath}>
                <img src={openIcon} alt="Open Dir" class="open-icon" />
            </button>
            </div>
            <button on:click={startDeletingDuplicatedFiles}>Delete Duplicated Files</button>
        </div>
        

        <div class="input-group">
            <input type="text" bind:value={oldLogsDays} placeholder="Enter the number of days to keep old logs" />
            <button on:click={startRemovingOldLogs}>Remove Old Logs</button>
        </div>
    </div>

    <div class="clean-system-section">
        <h2>Package Cleaning</h2>

        <button on:click={startRemovingUnusedPackages}>Remove Unused Packages</button>
        <button on:click={startCleaningCachePackages}>Clean Cache Packages</button>
    </div>

</div>

<style>
    .clean-system-container {
        display: flex;
        justify-content: space-between;
        max-width: 95%;
        margin: 2rem auto;
    }

    .clean-system-section {
        flex: 1;
        padding: 1rem;
        background-color: #282828;
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
        color: #fff;
    }
    input {
        margin-bottom: 10px;
        padding: 10px;
        width: 100%;
        box-sizing: border-box;
        border: 1px solid #444;
        background-color: #333;
        color: #fff;
        border-radius: 5px;
    }
    button {
        margin-bottom: 10px;
        padding: 10px 20px;
        border: none;
        background: #1abc9c;
        color: white;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }
    button:hover {
        background-color: #16a085;
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
