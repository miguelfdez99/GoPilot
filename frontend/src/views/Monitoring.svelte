<script lang="ts">
    import { onMount } from 'svelte';

    import {GetDistribution} from '../../wailsjs/go/backend/Backend';

    import CpuChart from '../components/CpuChart.svelte'
    import DiskChart from '../components/DiskChart.svelte'


    onMount(() => {
      getDistribution();
    });

    let distribution: string = "";

    function getDistribution() {
      GetDistribution().then(result => {
        distribution = result;
      });
    }


  </script>
  <div>
    <div>
      <CpuChart />
    </div>
    <div>
      <DiskChart />
    </div>
  </div>
