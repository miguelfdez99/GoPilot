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

<div class="container">
    <h2>Clean System</h2>

    <input
        type="text"
        bind:value={tempFilesDirPath}
        placeholder="Enter a directory path for temporary files"
    />
    <input
        type="number"
        bind:value={expireDays}
        placeholder="Enter the number of days to keep temporary files"
    />
    <button on:click={startDeletingTempFiles}>Delete Temporary Files</button>

    <button on:click={startRemovingUnusedPackages}>Remove Unused Packages</button>

    <button on:click={startCleaningCachePackages}>Clean Cache Packages</button>

    <input
        type="text"
        bind:value={duplicatedFilesDirPath}
        placeholder="Enter a directory path to find duplicated files"
    />
    <button on:click={startDeletingDuplicatedFiles}>Delete Duplicated Files</button>

    <input
        type="text"
        bind:value={oldLogsDays}
        placeholder="Enter the number of days to keep old logs"
    />
    <button on:click={startRemovingOldLogs}>Remove Old Logs</button>
</div>

<style>
    .container {
        padding: 20px;
        background-color: #222;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
        border-radius: 10px;
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
        background-color: #007bff;
        color: white;
        border-radius: 5px;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }
    button:hover {
        background-color: #0056b3;
    }
</style>
