<script lang="ts">
    import { onMount, afterUpdate } from "svelte";
    import { GetDiskUsage, GetFolderSizes } from "../../wailsjs/go/backend/Backend.js";
    import {
        Chart,
        DoughnutController,
        ArcElement,
        Title,
        Tooltip,
        CategoryScale,
    } from "chart.js";

    Chart.register(
        DoughnutController,
        ArcElement,
        Title,
        Tooltip,
        CategoryScale
    );

    let chart;
    let canvas;
    let folderSizes = [];

    onMount(async () => {
        const diskUsage = await GetDiskUsage();

        const ctx = canvas.getContext("2d");
        chart = new Chart(ctx, {
            type: "doughnut",
            data: {
                labels: ["Used", "Free"],
                datasets: [
                    {
                        data: [
                            diskUsage.Used / 1024 / 1024 / 1024, // Convert from bytes to GB
                            (diskUsage.Total - diskUsage.Used) / 1024 / 1024 / 1024, // Convert from bytes to GB
                        ],
                        backgroundColor: ["#4BC0C0", "#FF6384"],
                    },
                ],
            },
            options: {
                onClick: (event, elements) => {
                   
                    if (elements.length > 0) {
                        const chartElement = elements[0];
                        const label = chart.data.labels[chartElement.index];
                        if (label === "Used") {
                            fetchFolderSizes();
                        }
                    }
                },
            },
        });
    });

    async function fetchFolderSizes() {
        folderSizes = await GetFolderSizes("/");
        console.log("Folder sizes fetched", folderSizes);
        // Update the UI as needed
    }
</script>

<canvas bind:this={canvas} />

<div>
    {#each folderSizes as folderSize}
        <div>{folderSize.Path}: {Math.round(folderSize.Size / 1024 / 1024)} MB</div>
    {/each}
</div>
