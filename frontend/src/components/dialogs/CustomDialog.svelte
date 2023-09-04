<script lang="ts">
    import { settings } from "../../stores";

    let fontSize: string;
    let color: string;
    let fontFamily: string;
    let backgroundColor: string;
    let backgroundColor2: string;
    let inputColor: string;
    let buttonColor: string;
    settings.subscribe(($settings) => {
        fontSize = $settings.fontSize;
        color = $settings.color;
        fontFamily = $settings.fontFamily;
        backgroundColor = $settings.backgroundColor;
        backgroundColor2 = $settings.backgroundColor2;
        inputColor = $settings.inputColor;
        buttonColor = $settings.buttonColor;
    });

    $: {
        document.documentElement.style.setProperty(
            "--main-font-size",
            fontSize
        );
        document.documentElement.style.setProperty("--main-color", color);
        document.documentElement.style.setProperty(
            "--main-font-family",
            fontFamily
        );
        document.documentElement.style.setProperty(
            "--main-bg-color",
            backgroundColor
        );
        document.documentElement.style.setProperty(
            "--main-bg-color2",
            backgroundColor2
        );
        document.documentElement.style.setProperty(
            "--main-input-color",
            inputColor
        );
        document.documentElement.style.setProperty(
            "--main-button-color",
            buttonColor
        );
    }

    export let show: boolean = false;
    export let title: string = "";
    export let message: string = "";
    export let onClose: () => void = () => {};

    function handleClick(event: MouseEvent) {
        if (event.target === event.currentTarget) {
            onClose();
        }
    }

    function handleKeyDown(event: KeyboardEvent) {
        if (event.key === "Escape") {
            onClose();
        }
    }
</script>

{#if show}
    <div class="modal" on:click={handleClick} on:keydown={handleKeyDown}>
        <div class="modal-content">
            <h1>{title}</h1>
            <p>{@html message}</p>
            <div class="modal-buttons">
                <button on:click={onClose}>Ok</button>
            </div>
        </div>
    </div>
{/if}

<style>
    .modal {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: rgba(0, 0, 0, 0.8);
        z-index: 9999;
    }
    .modal-content {
        background-color: var(--main-bg-color2);
        color: var(--main-color);
        padding: 30px;
        border-radius: 15px;
        max-width: 500px;
        max-height: 80vh;
        overflow-y: auto;
        width: 90%;
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
    }
    .modal-buttons {
        margin-top: 30px;
        display: flex;
        justify-content: space-between;
    }
    .modal-buttons button {
        padding: 10px 20px;
        border: none;
        border-radius: 5px;
        font-size: 16px;
        cursor: pointer;
        transition: background-color 0.2s;
        background: var(--main-button-color);
        color: var(--main-color);
    }

    .modal-content h1 {
        color: var(--main-color);
        margin-bottom: 15px;
    }
    .modal-content p {
        line-height: 1.6;
    }
</style>
