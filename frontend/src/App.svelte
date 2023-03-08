<script lang="ts">
  import {Greet} from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }
</script>

<main>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={greet}>Greet</button>
  </div>
</main>

<style>


</style>
