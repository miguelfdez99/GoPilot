<script lang="ts">
    import { onMount } from 'svelte';
    import {ListPackages} from '../../wailsjs/go/main/App.js'

    let currentView = 'packages';
    let adminText: string

    onMount(() => {
      addEventListener('changeView', (event: CustomEvent) => {
        currentView = event.detail;
      });
    });

  let packages = [];

  function listPackages() {
    ListPackages().then(result => {
      packages = result;
    });
  }
  </script>

  <main>
    <div class="input-box" id="input">
      <h1>Packages</h1>
      <div>
        {#if packages.length > 0}
          <ul>
            {#each packages as pkg}
              <li>{pkg}</li>
            {/each}
          </ul>
        {:else}
          <p>No packages found</p>
        {/if}
        <button on:click={listPackages}>List Packages</button>
      </div>
    </div>
  </main>