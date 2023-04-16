<script lang="ts">
    import { onMount } from 'svelte';

    import {GetDistribution} from '../../wailsjs/go/main/App.js'

    import CpuInfo from '../components/CpuInfo.svelte'


    onMount(() => {
      getDistribution();
    });

    let distribution: string = "";
    let info: string = "";

    function getDistribution() {
      GetDistribution().then(result => {
        distribution = result;
      });
    }


  </script>
  <div>
    <div>
      <CpuInfo />
    </div>
  </div>

  <style>

  </style>