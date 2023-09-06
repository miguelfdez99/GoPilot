<script lang="ts">
    import { onMount } from "svelte";
    import { GetDiskUsage } from "../../../wailsjs/go/backend/Backend";
    import {
        Chart,
        DoughnutController,
        ArcElement,
        Title,
        Tooltip,
        CategoryScale,
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
                            diskUsage.Used / 1024 / 1024 / 1024,
                            (diskUsage.Total - diskUsage.Used) / 1024 / 1024 / 1024,
                        ],
                        backgroundColor: COLORS,
                    },
                ],
            },
            options: {
                animation: false,
                responsive: false,
                plugins: {
                    legend: {
                        display: true,
                        labels: {
                            color: colorH,
                        }
                    },
                    title: {
                        display: true,
                        text: "Disk Usage (GB)",
                        color: colorH
                    },
                    tooltip: {
                        callbacks: {
                            label: function (context) {
                                const label = context.label || "";
                                const value = context.parsed;
                                const total = context.dataset.data.reduce(
                                    (a, b) => a + b
                                );
                                const percentage = ((value / total) * 100).toFixed(2);
                                return `${label}: ${value.toFixed(2)} GB (${percentage}%)`;
                            },
                        },
                    },
                },
            },
        });
    });
</script>

<div class="chart-container">
    <canvas bind:this={canvas} width="400" height="400" />
</div>

<style>
    .chart-container {
      display: flex;
      justify-content: center;
      align-items: center;
    }
  </style>
