<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetCPUUsage } from "../../wailsjs/go/backend/Backend.js";
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

    let chart;
    let cpuUsages = [];
    let intervalId;
    let canvas;

    const MAX_HISTORY = 60; // Keep data for the last 60 seconds
    const COLORS = [
        "red",
        "blue",
        "green",
        "yellow",
        "purple",
        "cyan",
        "magenta",
        "orange",
        "teal",
        "pink",
        "lime",
        "deepskyblue",
        "violet",
        "gold",
        "darkgreen",
        "salmon",
    ];

    onMount(async () => {
        cpuUsages = await GetCPUUsage();

        const ctx = canvas.getContext("2d");
        chart = new Chart(ctx, {
            type: "line",
            data: {
                labels: ["", ...Array(MAX_HISTORY - 1).fill("")],
                datasets: cpuUsages.map((usage, i) => ({
                    label: `CPU ${i + 1} (${usage}%)`,
                    data: [null, ...Array(MAX_HISTORY - 1).fill(usage)],
                    borderColor: COLORS[i % COLORS.length],
                    borderWidth: 1,
                    fill: false,
                    tension: 0.1,
                })),
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true,
                    },
                },
            },
        });

        intervalId = setInterval(updateChart, 1000);
    });

    onDestroy(() => {
        clearInterval(intervalId);
    });

    async function updateChart() {
        const newCpuUsages = await GetCPUUsage();

        newCpuUsages.forEach((usage, i) => {
            chart.data.datasets[i].data.push(usage);

            const roundedPercentage = usage.toFixed(1);
            chart.data.datasets[i].label = `CPU ${
                i + 1
            } (${roundedPercentage}%)`;

            if (chart.data.datasets[i].data.length > MAX_HISTORY) {
                chart.data.datasets[i].data.shift();
            }
        });

        chart.update();
    }
</script>

<canvas bind:this={canvas} />
