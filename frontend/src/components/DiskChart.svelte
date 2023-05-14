<script lang="ts">
    import { onMount } from "svelte";
    import { GetDiskUsage } from "../../wailsjs/go/backend/Backend.js";
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
    let canvas: HTMLCanvasElement;
    const COLORS: string[] = ["#4BC0C0", "#FF6384"];

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
                        backgroundColor: COLORS,
                    },
                ],
            },
            options: {
                animation: false,
                responsive: true,
                plugins: {
                    title: {
                        display: true,
                        text: "Disk Usage (GB)",
                    },
                    tooltip: {
                        callbacks: {
                            label: function (context) {
                                const label = context.label || "";
                                const value = context.parsed;
                                const total = context.dataset.data.reduce(
                                    (a, b) => a + b
                                );
                                const percentage = ((value / total) * 100).toFixed(2); // Calculate the percentage
                                return `${label}: ${value.toFixed(2)} GB (${percentage}%)`;
                            },
                        },
                    },
                },
            },
        });
    });
</script>

<canvas bind:this={canvas} />
