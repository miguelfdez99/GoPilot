<script lang="ts">
  import CpuChart from "../components/charts/CpuChart.svelte";
  import DiskChart from "../components/charts/DiskChart.svelte";
  import MemChart from "../components/charts/MemChart.svelte";
  import NetworkChart from "../components/charts/NetworkChart.svelte";
  import dropdownIcon from "../assets/images/dropdown.png";
  import { settings } from '../stores';

    let colorH: string;
    settings.subscribe(($settings) => {
        colorH = $settings.color;
    });

    $: {
    document.documentElement.style.setProperty('--main-color', colorH);
  }

  let showCpuChart: boolean = true;
  let showDiskChart: boolean = true;
  let showMemChart: boolean = true;
  let showNetworkChart: boolean = true;

  function toggleCpuChart() {
    showCpuChart = !showCpuChart;
  }

  function toggleDiskChart() {
    showDiskChart = !showDiskChart;
  }

  function toggleMemChart() {
    showMemChart = !showMemChart;
  }

  function toggleNetworkChart() {
    showNetworkChart = !showNetworkChart;
  }
</script>

<div>
  <div>
    <div class="chart-header">
      <button class="dropdown-btn" on:click={toggleCpuChart}>
        <img src={dropdownIcon} alt="Dropdown" class="dropdown-icon" />
      </button>
      <h2>CPU</h2>
    </div>
    {#if showCpuChart}
      <CpuChart />
    {/if}
  </div>

  <div>
    <div class="chart-header">
      <button class="dropdown-btn" on:click={toggleDiskChart}>
        <img src={dropdownIcon} alt="Dropdown" class="dropdown-icon" />
      </button>
      <h2>Disk Space</h2>
    </div>
    {#if showDiskChart}
      <DiskChart />
    {/if}
  </div>

  <div>
    <div class="chart-header">
      <button class="dropdown-btn" on:click={toggleMemChart}>
        <img src={dropdownIcon} alt="Dropdown" class="dropdown-icon" />
      </button>
      <h2>RAM</h2>
    </div>
    {#if showMemChart}
      <MemChart />
    {/if}
  </div>

  <div>
    <div class="chart-header">
      <button class="dropdown-btn" on:click={toggleNetworkChart}>
        <img src={dropdownIcon} alt="Dropdown" class="dropdown-icon" />
      </button>
      <h2>Network</h2>
    </div>
    {#if showNetworkChart}
      <NetworkChart />
    {/if}
  </div>
</div>

<style>
  .chart-header {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .chart-header h2 {
    margin: 0;
    line-height: 24px;
    padding: 10px;
    color: var(--main-color);;
  }

  .dropdown-btn {
    border: none;
    background: none;
    cursor: pointer;
    padding: 0;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 24px;
    height: 24px;
  }

  .dropdown-icon {
    width: 100%;
    height: 100%;
  }
</style>
