<script>
    import Home from './Home.svelte';
    import Users from './Users.svelte';

    let currentView = Home;

    function changeView(view) {
      if (view === 'home') {
        currentView = Home;
      } else if (view === 'users') {
        currentView = Users;
      }
    }
  </script>

  <nav>
    <ul>
      <li><a on:click={() => changeView('home')}>Home</a></li>
      <li><a on:click={() => changeView('users')}>Users</a></li>
    </ul>
  </nav>

  <main>
    <svelte:component this={currentView} />
  </main>

  <style>
    /* Styles for the nav and main sections go here */
  </style>
