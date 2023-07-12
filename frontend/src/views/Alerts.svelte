<script lang="ts">
    import { MonitorFile, MonitorDir } from "../../wailsjs/go/backend/Backend";

    let filepath = '';
    let dirpath = '';
  
    const startMonitoringFile = () => {
        MonitorFile(filepath);
    }
  
    const startMonitoringDir = () => {
        MonitorDir(dirpath);
    }
  
  </script>
  
  <div>
    <h1>File System Monitor</h1>
    <div>
      <h2>File Monitor</h2>
      <input type="text" bind:value={filepath} placeholder="Enter a file path" />
      <button on:click={startMonitoringFile}>Start Monitoring File</button>
    </div>
    <div>
      <h2>Directory Monitor</h2>
      <input type="text" bind:value={dirpath} placeholder="Enter a directory path" />
      <button on:click={startMonitoringDir}>Start Monitoring Directory</button>
    </div>
  </div>
  