<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetMemoryUsage } from "../../wailsjs/go/backend/Backend.js";
    import {
        Chart,
        LineController,
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
    } from "chart.js";

    Chart.register(
        LineController,
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement
    );

    let chart: Chart;
    let intervalId: number;
    let canvas: HTMLCanvasElement;

    const MAX_HISTORY = 300; // Keep data for the last 60 seconds

    onMount(async () => {
        const ctx = canvas.getContext("2d");
        chart = new Chart(ctx, {
            type: "line",
            data: {
                labels: Array(MAX_HISTORY).fill(""),
                datasets: [
                    {
                        label: `RAM Usage`,
                        data: Array(MAX_HISTORY).fill(null),
                        borderColor: "blue",
                        borderWidth: 1,
                        fill: false,
                        tension: 0.1,
                        pointRadius: 0,
                    },
                    {
                        label: `Swap Usage`,
                        data: Array(MAX_HISTORY).fill(null),
                        borderColor: "red",
                        borderWidth: 1,
                        fill: false,
                        tension: 0.1,
                        pointRadius: 0,
                    },
                ],
            },
            options: {
                animation: false,
                scales: {
                    y: {
                        beginAtZero: true,
                        min: 0,
                        max: 100,
                        ticks: {
                            stepSize: 20,
                        },
                    },
                },
                plugins: {
                    tooltip: {
                        enabled: true,
                    },
                },
            },
        });

        intervalId = setInterval(updateChart, 100);
    });

    onDestroy(() => {
        clearInterval(intervalId);
    });

    async function updateChart() {
        const newMemoryUsage = await GetMemoryUsage();

        chart.data.datasets[0].data.shift();
        chart.data.datasets[0].data.push(newMemoryUsage.ram);
        const roundedRamPercentage = newMemoryUsage.ram.toFixed(1);
        chart.data.datasets[0].label = `RAM Usage (${roundedRamPercentage}%)`;

        chart.data.datasets[1].data.shift();
        chart.data.datasets[1].data.push(newMemoryUsage.swap);
        const roundedSwapPercentage = newMemoryUsage.swap.toFixed(1);
        chart.data.datasets[1].label = `Swap Usage (${roundedSwapPercentage}%)`;

        chart.update();
    }
</script>

<canvas bind:this={canvas} />

<style>
    canvas {
        width: 100%;
        height: 100%;
    }
</style>
