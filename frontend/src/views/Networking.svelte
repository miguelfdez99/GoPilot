<script lang="ts">
  import { onMount } from "svelte";
  import { GetNetworkInterfaces, SetInterfaceStatus } from "../../wailsjs/go/backend/Backend";

  let interfaces = [];

    async function fetchInterfaces() {
    try {
      const result = await GetNetworkInterfaces();
      interfaces = [...result];
    } catch (err) {
      console.error(err);
    }
  }


  async function changeStatus(name: string, currentStatus: string) {
    const newStatus = currentStatus === 'up' ? 'down' : 'up';
    try {
      await SetInterfaceStatus(name, newStatus);
      fetchInterfaces();
    } catch (err) {
      console.error(err);
    }
  }

  onMount(fetchInterfaces);
</script>

<div class="pico-container">
  <h2 class="pico-title">Network interfaces</h2>
  <table class="pico-table">
    <thead>
      <tr>
        <th>Interface Name</th>
        <th>Status</th>
        <th>IP</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
      {#each interfaces as interf (interf.Name)}
        <tr>
          <td>{interf.Name}</td>
          <td>{interf.Status}</td>
          <td>{interf.IP}</td>
          <td>
            <button on:click={() => changeStatus(interf.Name, interf.Status)}
              disabled={interf.Status.toLowerCase() === 'unknown'}>
              {interf.Status.toLowerCase() === 'up' ? 'Bring down' : 
               interf.Status.toLowerCase() === 'down' ? 'Bring up' : 'Unkown' }
            </button>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style>
  .pico-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }

  td, th {
    color: white;
  }

  h2 {
    color: white;
  }
</style>
