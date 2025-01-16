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
    <!-- Make the div focusable and interactive -->
    <div
        class="modal"
        role="dialog"
        aria-modal="true"
        tabindex="0"
        on:click={handleClick}
        on:keydown={handleKeyDown}
    >
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
    :global(:root) {
        /* Tokyo Night theme colors */
        --color-bg-primary: #1a1b26;
        --color-bg-secondary: #16161e;
        --color-bg-tertiary: #1f2335;
        --color-text-primary: #a9b1d6;
        --color-text-secondary: #787c99;
        --color-accent-blue: #7aa2f7;
        --color-accent-purple: #9d7cd8;
        --color-accent-cyan: #7dcfff;
        --color-accent-green: #9ece6a;
        --color-accent-orange: #ff9e64;
        --color-accent-red: #f7768e;

        /* Layout */
        --spacing-sm: 0.5rem;
        --spacing-md: 1rem;
        --spacing-lg: 2rem;
    }

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
        background-color: var(--color-bg-secondary);
        color: var(--color-text-primary);
        padding: var(--spacing-lg);
        border-radius: 0.5rem;
        max-width: 500px;
        max-height: 80vh;
        overflow-y: auto;
        width: 90%;
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.5);
    }

    .modal-content h1 {
        color: var(--color-accent-blue);
        margin-bottom: var(--spacing-md);
        font-size: 1.5rem;
    }

    .modal-content p {
        color: var(--color-text-primary);
        line-height: 1.6;
        font-size: 1rem;
    }

    .modal-buttons {
        margin-top: var(--spacing-lg);
        display: flex;
        justify-content: flex-end;
    }

    .modal-buttons button {
        padding: var(--spacing-sm) var(--spacing-md);
        border: none;
        border-radius: 0.5rem;
        background-color: var(--color-accent-blue);
        color: var(--color-bg-primary);
        cursor: pointer;
        transition: background-color 0.2s ease-in-out;
    }

    .modal-buttons button:hover {
        background-color: var(--color-accent-cyan);
    }
</style>
