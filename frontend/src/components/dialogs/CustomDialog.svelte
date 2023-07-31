<script lang="ts">
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
        background-color: rgba(0,0,0,0.8); 
        z-index: 9999;
    }
    .modal-content {
        background-color: #2f2f2f; 
        color: #f5f5f5;
        padding: 30px;
        border-radius: 15px;
        max-width: 500px;
        max-height: 80vh;
        overflow-y: auto;
        width: 90%;
        box-shadow: 0 5px 15px rgba(0,0,0,0.5); 
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
        background-color: #4b4b4b; 
        color: #f5f5f5;
    }
    .modal-buttons button:hover {
        background-color: #666;
    }
    .modal-content h1 {
        color: #f5f5f5;
        margin-bottom: 15px;
    }
    .modal-content p {
        color: #bdbdbd;
        line-height: 1.6;
    }
</style>
