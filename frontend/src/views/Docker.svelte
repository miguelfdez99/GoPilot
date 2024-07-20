<script lang="ts">
    import { onMount } from "svelte";
    import { writable } from "svelte/store";
    import {
        ListImages,
        RemoveImage,
        ListContainers,
        GetContainerLogs,
        OpenDialogError,
    } from "../../wailsjs/go/backend/Backend";

    let images = writable([]);
    let containers = writable([]);

    onMount(() => {
        listImages();
        listContainers();
    });

    async function listImages() {
        try {
            const result = await ListImages();
            images.set(
                result.map((imageString) => {
                    const [repository, imageId, tag, created, size] =
                        imageString.split(/\s+/);
                    const createdDate = new Date(
                        parseInt(created) * 1000,
                    ).toLocaleString();
                    const sizeMB =
                        (parseInt(size, 10) / 1024 / 1024).toFixed(2) + "MB";
                    return {
                        repository,
                        tag,
                        imageId,
                        created: createdDate,
                        size: sizeMB,
                    };
                }),
            );
        } catch (e) {
            OpenDialogError(`Failed to list images: ${e.message}`);
        }
    }

    async function listContainers() {
        try {
            const result = await ListContainers();
            containers.set(
                result.map((containerString) => {
                    const [
                        containerId,
                        image,
                        name,
                        command,
                        ports,
                        status,
                        created,
                    ] = containerString.split(/\s+/);
                    const createdDate = new Date(
                        parseInt(created) * 1000,
                    ).toLocaleString();
                    return {
                        containerId,
                        image,
                        name,
                        command,
                        ports,
                        status,
                        created: createdDate,
                    };
                }),
            );
        } catch (e) {
            OpenDialogError(`Failed to list containers: ${e.message}`);
        }
    }

    async function removeImage(imageId: string) {
        try {
            await RemoveImage(imageId);
            listImages();
        } catch (e) {
            OpenDialogError(`Failed to remove image: ${e.message}`);
        }
    }

    async function getContainerLogs(containerId) {
        try {
            const logs = await GetContainerLogs(containerId);
            const logsDiv = document.getElementById("logs");
            logsDiv.innerHTML = logs.replace(/\n/g, "<br>");
        } catch (e) {
            OpenDialogError(
                `Failed to get container logs: ${containerId} - ${e.message}`,
            );
        }
    }

    let showImages: boolean = false;
    let showContainers: boolean = false;
    let showLogs: boolean = false;

    function ShowImages(): boolean {
        showImages = !showImages;
        return showImages;
    }

    function ShowContainers(): boolean {
        showContainers = !showContainers;
        return showContainers;
    }

    function ShowLogs(): boolean {
        showLogs = !showLogs;
        return showLogs;
    }
</script>

<div class="container">
    <h1>Docker Images</h1>
    <button on:click={ShowImages}>Show Images list</button>
    {#if showImages}
        <table class="images-table">
            <thead>
                <tr>
                    <th>Repository</th>
                    <th>Tag</th>
                    <th>Image ID</th>
                    <th>Created</th>
                    <th>Size</th>
                    <th>Remove</th>
                </tr>
            </thead>
            <tbody>
                {#each $images as image}
                    <tr>
                        <td>{image.repository}</td>
                        <td>{image.tag}</td>
                        <td>{image.imageId}</td>
                        <td>{image.created}</td>
                        <td>{image.size}</td>
                        <td
                            ><button on:click={() => removeImage(image.imageId)}
                                >Remove</button
                            ></td
                        >
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}

    <button on:click={ShowContainers}>Show Containers</button>
    {#if showContainers}
        <table class="images-table">
            <thead>
                <tr>
                    <th>Container ID</th>
                    <th>Image</th>
                    <th>Name</th>
                    <th>Command</th>
                    <th>Ports</th>
                    <th>Status</th>
                    <th>Created</th>
                    <th>Logs</th>
                </tr>
            </thead>
            <tbody>
                {#each $containers as container}
                    <tr>
                        <td>{container.containerId}</td>
                        <td>{container.image}</td>
                        <td>{container.name}</td>
                        <td>{container.command}</td>
                        <td>{container.ports}</td>
                        <td>{container.status}</td>
                        <td>{container.created}</td>
                        <td
                            ><button
                                on:click={() =>
                                    getContainerLogs(container.containerId)}
                                >Logs</button
                            ></td
                        >
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}

    <button on:click={ShowLogs}>Show Logs</button>
    {#if showLogs}
        <div id="logs"></div>
        <div>Logs are shown here</div>
    {/if}
</div>

<style>
    .container {
        display: flex;
        flex-direction: column;
        max-width: 90%;
        margin: 2rem auto;
        background-color: var(--main-bg-color);
        padding: 1rem;
        color: var(--main-color);
    }

    .images-table {
        width: 100%;
        border-collapse: collapse;
    }

    th,
    td {
        border: 1px solid #ddd;
        padding: 8px;
        text-align: left;
    }

    th {
        background-color: #333;
        color: white;
    }

    button {
        padding: 0.8em 1em;
        border: none;
        border-radius: 4px;
        background-color: #5a6268;
        color: white;
        cursor: pointer;
        margin-bottom: 1em;
    }
</style>
