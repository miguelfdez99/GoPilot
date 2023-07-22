<script lang="ts">
    import { DeleteTempFiles, RemoveUnusedPackages, CleanCachePackages, DuplicatedFiles, RemoveOldLogs } from "../../wailsjs/go/backend/Backend";

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
</script>

<div class="clean-system-container">

    <div class="clean-system-section">
        <h2>Clean System</h2>

        <div class="input-group">
            <input type="text" bind:value={tempFilesDirPath} placeholder="Enter a directory path for temporary files" />
            <input type="number" bind:value={expireDays} placeholder="Enter the number of days to keep temporary files" />
            <button on:click={startDeletingTempFiles}>Delete Temporary Files</button>
        </div>

        <div class="input-group">
            <input type="text" bind:value={duplicatedFilesDirPath} placeholder="Enter a directory path to find duplicated files" />
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
</style>
