<script lang="ts">
    export let show: boolean = false;
    export let title: string = "";
    export let message: string = "";
    export let onConfirm: () => void = () => {};
    export let onClose: () => void = () => {};
    export let confirmButton: boolean = true;
</script>

{#if show}
    <div class="modal">
        <div class="modal-content">
            <h1>{title}</h1>
            <p>{message}</p>
            <div class="modal-buttons">
                <button on:click={onClose}>Close</button>
                {#if confirmButton}
                    <button on:click={onConfirm}>Confirm</button>
                {/if}
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
        background-color: rgba(0,0,0,0.5);

    }
    .modal-content {
        background-color: white;
        padding: 20px;
        border-radius: 10px;
        max-width: 90%;
    }
    .modal-buttons {
        margin-top: 20px;
        display: flex;
        justify-content: flex-end;
    }
    .modal-buttons button {
        margin-left: 10px;
    }

    .modal-content h1,
    .modal-content p {
        color: black;
    }

</style>