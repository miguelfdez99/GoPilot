<script lang="ts">
  import AddUser from "../components/users/AddUser.svelte";
  import DelUser from "../components/users/DelUser.svelte";
  import ModifyUser from "../components/users/ModifyUser.svelte";
  import AddGroup from "../components/groups/AddGroup.svelte";
  import DelGroup from "../components/groups/DelGroup.svelte";
  import ModifyGroup from "../components/groups/ModifyGroup.svelte";

  let viewState = 'default';

  function setViewState(newViewState: string): void {
    viewState = newViewState;
  }
</script>

{#if viewState === 'default'}
  <main>
    <div class="section" id="users">
      <h1>Users</h1>
      <button class="btn" on:click={() => setViewState('addUser')}>Add User</button>
      <button class="btn" on:click={() => setViewState('delUser')}>Delete User</button>
      <button class="btn" on:click={() => setViewState('modUser')}>Modify User</button>
    </div>

    <div class="section" id="groups">
      <h1>Groups</h1>
      <button class="btn" on:click={() => setViewState('addGroup')}>Create Group</button>
      <button class="btn" on:click={() => setViewState('delGroup')}>Delete Group</button>
      <button class="btn" on:click={() => setViewState('modGroup')}>Modify Group</button>
    </div>
  </main>
{:else if viewState === 'addUser'}
  <AddUser on:dismiss={() => setViewState('default')} />
{:else if viewState === 'delUser'}
  <DelUser on:dismiss={() => setViewState('default')} />
{:else if viewState === 'modUser'}
  <ModifyUser on:dismiss={() => setViewState('default')} />
{:else if viewState === 'addGroup'}
  <AddGroup on:dismiss={() => setViewState('default')} />
{:else if viewState === 'delGroup'}
  <DelGroup on:dismiss={() => setViewState('default')} />
{:else if viewState === 'modGroup'}
  <ModifyGroup on:dismiss={() => setViewState('default')} />
{/if}


<style>
  main {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    padding-top: 100px;
    padding-bottom: 30px;
  }

  .section {
    background-color: #f5f5f5;
    padding: 2rem;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    margin-right: 2rem;
  }

  h1 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
    color: #333;
  }

  .btn {
    background-color: #333;
    color: #fff;
    border: none;
    border-radius: 4px;
    padding: 0.8rem 1.2rem;
    margin-right: 1rem;
    cursor: pointer;
    font-size: 1rem;
    transition: background-color 0.3s ease-in-out;
  }

  .btn:hover {
    background-color: #555;
  }

  @media only screen and (max-width: 768px) {
    .section {
      padding: 1rem;
    }

    h1 {
      font-size: 1.5rem;
    }

    .btn {
      font-size: 0.9rem;
      padding: 0.6rem 1rem;
      margin-right: 0.5rem;
    }
  }
</style>
