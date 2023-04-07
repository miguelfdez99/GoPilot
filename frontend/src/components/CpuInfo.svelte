<script lang="ts">
    import { onMount } from 'svelte';
    import { GetLSCPU } from "../../wailsjs/go/main/App.js";

  let systemInfo = {
    architecture: '',
    cpus: '',
    modelName: '',
    threadPerCore: '',
    corePerSocket: '',
    socket: '',
    cpuModes: '',
  };

  onMount(async () => {
    const response = await GetLSCPU();
    systemInfo = parseResponse(response);
  });

  function parseResponse(response) {
    const data = JSON.parse(response);
    return {
      architecture: data.architecture,
      cpus: data.cpus,
      modelName: data.modelName,
      threadPerCore: data.threadPerCore,
      corePerSocket: data.corePerSocket,
      socket: data.socket,
      cpuModes: data.cpuModes,
    };
  }
</script>


    <template>
  <div class="container">
    <h2>System Information</h2>
    <div class="row">
      <div class="col-6">Architecture:</div>
      <div class="col-6">{systemInfo.architecture}</div>
    </div>
    <div class="row">
      <div class="col-6">CPU(s):</div>
      <div class="col-6">{systemInfo.cpus}</div>
    </div>
    <div class="row">
      <div class="col-6">Model name:</div>
      <div class="col-6">{systemInfo.modelName}</div>
    </div>
    <div class="row">
      <div class="col-6">Threads per core:</div>
      <div class="col-6">{systemInfo.threadPerCore}</div>
    </div>
    <div class="row">
      <div class="col-6">Cores per socket:</div>
      <div class="col-6">{systemInfo.corePerSocket}</div>
    </div>
    <div class="row">
      <div class="col-6">Socket(s):</div>
      <div class="col-6">{systemInfo.socket}</div>
    </div>
    <div class="row">
      <div class="col-6">CPU modes:</div>
      <div class="col-6">{systemInfo.cpuModes}</div>
    </div>
  </div>
</template>

<style>
  .container {
    padding: 20px;
    background-color: #65d81d;
    color: black;
  }

  .row {
    display: flex;
    margin-bottom: 10px;
  }

  .col-6 {
    flex: 1;
  }

  .col-6:last-child {
    text-align: right;
  }

  h2 {
    margin-top: 0;
  }
</style>
