<script lang="ts">
  import { onMount } from 'svelte';
  import { GetLSCPU, GetSystemInfo } from '../../../wailsjs/go/backend/Backend';

  let cpuInfo = {
    architecture: '',
    cpus: '',
    modelName: '',
    threadPerCore: '',
    corePerSocket: '',
    socket: '',
    cpuModes: '',
  };

  let systemInfo = {
    osName: '',
    kernelVer: '',
    uptime: '',
  };

  onMount(async () => {
    const response = await GetLSCPU();
    cpuInfo = parseResponse(response);
    const response2 = await GetSystemInfo();
    systemInfo = parseSystemInfo(response2);
  });

  function parseResponse(response: string) {
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

  function parseSystemInfo(response: string) {
    const data = JSON.parse(response);
    return {
      osName: data.osName,
      kernelVer: data.kernelVer,
      uptime: data.uptime,
    };
  }

</script>

<template>
  <div class="container">
    <h2>System Information</h2>
    <div class="row">
      <div>OS Name:</div>
      <div>{systemInfo.osName}</div>
    </div>
    <div class="row">
      <div>Kernel Version:</div>
      <div>{systemInfo.kernelVer}</div>
    </div>
    <div class="row">
      <div>Uptime:</div>
      <div>{systemInfo.uptime}</div>
    </div>
    <h2>CPU Information</h2>
    <div class="row">
      <div>Architecture:</div>
      <div>{cpuInfo.architecture}</div>
    </div>
    <div class="row">
      <div>CPU(s):</div>
      <div>{cpuInfo.cpus}</div>
    </div>
    <div class="row">
      <div>Model name:</div>
      <div>{cpuInfo.modelName}</div>
    </div>
    <div class="row">
      <div>Threads per core:</div>
      <div>{cpuInfo.threadPerCore}</div>
    </div>
    <div class="row">
      <div>Cores per socket:</div>
      <div>{cpuInfo.corePerSocket}</div>
    </div>
    <div class="row">
      <div>Socket(s):</div>
      <div>{cpuInfo.socket}</div>
    </div>
    <div class="row">
      <div>CPU modes:</div>
      <div>{cpuInfo.cpuModes}</div>
    </div>
  </div>
</template>

<style>
  .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    color: whitesmoke;
  }

  h2 {
    font-size: 3rem;
    margin-bottom: 2rem;
    color: white;
  }

  .row {
    display: flex;
    justify-content: space-between;
    width: 40%;
    margin-bottom: 1rem;
    font-size: 1.2rem;
  }

  .row > div:first-child {
    font-weight: bold;
  }

  @media screen and (max-width: 768px) {
    .row {
      width: 80%;
    }
  }
</style>