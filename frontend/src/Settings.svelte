<script lang="ts">
    import { settings, lightTheme, darkTheme, highContrastTheme, nordTheme, oceanTheme, sunnyTheme } from './stores';
    import { onDestroy } from 'svelte';
  
    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    let showInfoButton: boolean;
    let currentSettings = {};
    let theme = 'dark';

    const unsubscribe = settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor; 
        buttonColor = $settings.buttonColor;
        showInfoButton = $settings.showInfoButton;
        currentSettings = { ...$settings };
    });

    const applyTheme = (selectedTheme: string) => {
        theme = selectedTheme;
        switch(selectedTheme) {
            case 'dark':
                settings.set(darkTheme);
                break;
            case 'light':
                settings.set(lightTheme);
                break;
            case 'high-contrast':
                settings.set(highContrastTheme);
                break;
            case 'nord':
                settings.set(nordTheme);
                break;
            case 'ocean':
                settings.set(oceanTheme);
                break;
            case 'sunny':
                settings.set(sunnyTheme);
                break;
            default:
                settings.set(darkTheme);
        }
    };

    let fontSizeNumber = parseInt(fontSize.replace('px', ''));

    $: currentSettings = { ...currentSettings };
    $: fontSize = `${fontSizeNumber}px`;

    $: {
    document.documentElement.style.setProperty('--main-font-size', fontSize);
    document.documentElement.style.setProperty('--main-color', color);
    document.documentElement.style.setProperty('--main-font-family', fontFamily);
    document.documentElement.style.setProperty('--main-bg-color', backgroundColor);
    document.documentElement.style.setProperty('--main-bg-color2', backgroundColor2);
    document.documentElement.style.setProperty('--main-input-color', inputColor);
    document.documentElement.style.setProperty('--main-button-color', buttonColor);
  }

    onDestroy(unsubscribe);
</script>

<div class="container">
    <h1>Settings</h1>
    <div class="settings-grid">
        <div class="setting-group theme-selector">
        <label for="theme">Theme</label>
        <select bind:value={theme} on:change={e => applyTheme(e.currentTarget.value)}>
            <option value="dark">Dark</option>
            <option value="light">Light</option>
            <option value="high-contrast">High Contrast</option>
            <option value="nord">Nord</option>
            <option value="ocean">Ocean</option>
            <option value="sunny">Sunny</option>
        </select>
    </div>
        <div class="setting-group">
            <label for="fontSize">Font Size</label>
            <input type="number" min="0" bind:value={fontSizeNumber} on:blur={() => settings.update(n => ({...n, fontSize}))}>
        </div>
        <div class="setting-group">
            <label for="fontFamily">Font Family</label>
            <select bind:value={fontFamily} on:blur={() => settings.update(n => ({...n, fontFamily}))}>
                <option value="Arial">Arial</option>
                <option value="Roboto">Roboto</option>
                <option value="Times New Roman">Times New Roman</option>
                <option value="Courier New">Courier New</option>
                <option value="Verdana">Verdana</option>
                <option value="Georgia">Georgia</option>
                <option value="Palatino">Palatino</option>
                <option value="Garamond">Garamond</option>
            </select>
        </div>
        <div class="setting-group">
            <label for="color">Color</label>
            <input type="color" bind:value={color} on:blur={() => settings.update(n => ({...n, color}))}>
        </div>
        <div class="setting-group">
            <label for="backgroundColor">Background Color</label>
            <input type="color" bind:value={backgroundColor} on:blur={() => settings.update(n => ({...n, backgroundColor}))}>
        </div>
        <div class="setting-group">
            <label for="backgroundColor2">Secondary Background Color</label>
            <input type="color" bind:value={backgroundColor2} on:blur={() => settings.update(n => ({...n, backgroundColor2}))}>
        </div>
        <div class="setting-group">
            <label for="inputColor">Input Color</label>
            <input type="color" bind:value={inputColor} on:blur={() => settings.update(n => ({...n, inputColor}))}>
        </div>
        <div class="setting-group">
            <label for="buttonColor">Button Color</label>
            <input type="color" bind:value={buttonColor} on:blur={() => settings.update(n => ({...n, buttonColor}))}>
        </div>
        <div class="setting-group custom-checkbox">
            <label for="showInfoButton">Show Info Button</label>
            <input id="showInfoButton" type="checkbox" class="hidden-checkbox" bind:checked={showInfoButton} on:change={() => settings.update(n => ({...n, showInfoButton}))}>
            <span class="styled-checkbox"></span>
        </div>                        
    </div>
</div>

<style>
    .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    font-family: 'Roboto', sans-serif;
    background-color: var(--main-bg-color2);
    padding: 2rem;
    border-radius: 15px;
    color: var(--main-color);
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.25);
    width: 80%;
    margin: auto;
    max-width: 1000px;
}

h1 {
    font-size: 2.5rem;
    margin-bottom: 2rem;
    text-transform: capitalize;
    letter-spacing: 1.5px;
    margin-top: 0;
    color: var(--main-color);
}

.settings-grid {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 2rem;
    align-items: center; 
}

.setting-group {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
}

label {
    font-size: 1.1rem;
    color: var(--main-color);
    font-size: var(--main-font-size);
    font-family: var(--main-font-family);
    letter-spacing: 0.8px;
    text-transform: capitalize;
    font-weight: 500;
}

input, select {
    width: 100%;
    padding: 0.7rem;
    border: none;
    border-radius: 5px;
    background-color: var(--main-bg-color);
    color: var(--main-color);
    font-size: var(--main-font-size);
    font-family: var(--main-font-family);
    font-size: 1rem;
    transition: background 0.3s;
}

input[type="color"] {
    padding: 0.5rem;
}

.theme-selector {
    grid-column: 1 / 3;
    margin-top: 1rem;
}

</style>

