<script lang="ts">
    import { onMount } from 'svelte';

    import {GetDistribution} from '../../wailsjs/go/main/App.js'

    import CpuInfo from '../components/CpuInfo.svelte'

    let currentView = 'system';

    onMount(() => {
      addEventListener('changeView', (event: CustomEvent) => {
        currentView = event.detail;
      });
    });

    let distribution: string = "";
    let info: string = "";
    let cpuInfo = {};

    function getDistribution() {
      GetDistribution().then(result => {
        distribution = result;
      });
    }


  </script>

  <h1>System</h1>

  <div>
    <p>You are using {distribution} based distribution</p>
    <p>Distribution: {distribution}</p>
    <button on:click={getDistribution}>Get Distribution</button>
    <div>
      <CpuInfo />
    </div>
  </div>

  <style>

  </style>