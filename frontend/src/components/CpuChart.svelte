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

    let chart: Chart;
    let intervalId: number;
    let canvas: HTMLCanvasElement;

    const MAX_HISTORY = 300;
    let COLORS: string[] = [];

    function generateColor(i: number, max_cpu: number): string {
        const hue = i * 360 / max_cpu;
        return `hsl(${hue}, 100%, 50%)`;
    }

    onMount(async () => {
        const ctx = canvas.getContext("2d");
        const newCpuUsages = await GetCPUUsage();
        const MAX_CPU = newCpuUsages.length;
        COLORS = Array.from({length: MAX_CPU}, (_, i) => generateColor(i, MAX_CPU));
        chart = new Chart(ctx, {
            type: "line",
            data: {
                labels: Array(MAX_HISTORY).fill(""),
                datasets: COLORS.map((color, i) => ({
                    label: `CPU ${i + 1}`,
                    data: Array(MAX_HISTORY).fill(null),
                    borderColor: color,
                    borderWidth: 1,
                    fill: false,
                    tension: 0.1,
                    pointRadius: 0,
                })),
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
                    legend: {
                        display: true,
                    },
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
        const newCpuUsages = await GetCPUUsage();

        newCpuUsages.forEach((usage, i) => {
            chart.data.datasets[i].data.shift();
            chart.data.datasets[i].data.push(usage);
            const roundedPercentage = usage.toFixed(1);
            chart.data.datasets[i].label = `CPU ${
                i + 1
            } (${roundedPercentage}%)`;
        });

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
