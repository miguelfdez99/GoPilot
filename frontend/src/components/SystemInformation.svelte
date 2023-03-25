<script lang="ts">
    import { onMount } from 'svelte';
    import {GetDistribution} from '../../wailsjs/go/main/App.js'

    let currentView = 'system';

    onMount(() => {
      addEventListener('changeView', (event: CustomEvent) => {
        currentView = event.detail;
      });
      getDistribution();
    });

    let distribution: string = "";

    function getDistribution() {
      GetDistribution().then(result => {
        distribution = result;
      });
    }

  </script>

  <h1>System</h1>

  <div>
    <p>Distribution: {distribution}</p>
    <button on:click={getDistribution}>Get Distribution</button>
  </div>