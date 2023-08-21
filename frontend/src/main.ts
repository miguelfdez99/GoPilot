import './style.css'
import App from './App.svelte'
import { settings } from './stores';

const app = new App({
  target: document.getElementById('app')
})

settings.subscribe($settings => {
  document.documentElement.style.setProperty('--background-color', $settings.backgroundColor);
});

export default app
