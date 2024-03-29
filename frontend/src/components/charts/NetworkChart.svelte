<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetNetworkUsage } from "../../../wailsjs/go/backend/Backend";
    import {
        Chart,
        LineController,
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Tooltip,
        Title,
        Legend,
    } from "chart.js";
    import { settings } from '../../stores';

    let colorH: string;
    settings.subscribe(($settings) => {
        colorH = $settings.color;
    });

    $: {
    document.documentElement.style.setProperty('--main-color', colorH);
  }

    Chart.register(
        LineController,
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Tooltip,
        Title,
        Legend
    );

    let chart: Chart;
    let intervalId: number;
    let canvas: HTMLCanvasElement;

    const MAX_HISTORY = 300;

    onMount(async () => {
        const ctx = canvas.getContext("2d");
        chart = new Chart(ctx, {
            type: "line",
            data: {
                labels: Array(MAX_HISTORY).fill(""),
                datasets: [
                    {
                        label: `MB Sent/s`,
                        data: Array(MAX_HISTORY).fill(null),
                        borderColor: "blue",
                        borderWidth: 1,
                        fill: false,
                        tension: 0.1,
                        pointRadius: 0,
                    },
                    {
                        label: `MB Received/s`,
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
                responsive: true,
                animation: false,
                plugins: {
                    tooltip: {
                        enabled: true,
                    },
                    title: {
                        display: true,
                        text: "Network Usage",
                        color: colorH
                    },
                    legend: {
                        position: "top",
                        labels: {
                            color: colorH,
                        }
                    },
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        ticks: {
                            stepSize: 20,
                        },
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
        const newNetworkUsage = await GetNetworkUsage();

        chart.data.datasets[0].data.shift();
        chart.data.datasets[0].data.push(newNetworkUsage.bytes_sent);
        chart.data.datasets[0].label = `MB Sent/s (${newNetworkUsage.bytes_sent.toFixed(
            2
        )})`;

        chart.data.datasets[1].data.shift();
        chart.data.datasets[1].data.push(newNetworkUsage.bytes_received);
        chart.data.datasets[1].label = `MB Received/s (${newNetworkUsage.bytes_received.toFixed(
            2
        )})`;

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
