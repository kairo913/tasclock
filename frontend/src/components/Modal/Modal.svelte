<script lang="ts">
    import type { Writable } from "svelte/store";

    export let show: Writable<boolean>;
    export let header: string = "";

    export let handler: () => Promise<boolean>;

    const modalHandler = () => {
        if (!handler()) return;
        dialog.close();
    };

    let dialog: HTMLDialogElement;

    $: if (dialog && $show) dialog.showModal();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
    bind:this={dialog}
    on:close={() => show.set(false)}
    on:click|self={() => dialog.close()}
    class="open:animate-zoom open:backdrop:animate-fade max-w-min overflow-hidden rounded-lg border-0 p-0 backdrop:bg-black/30"
>
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div on:click|stopPropagation class="px-4 py-2">
        <form class="flex flex-col space-y-2">
            <h2>{header}</h2>
            <hr />
            <slot />
            <hr />
            <div class="flex flex-row justify-end space-x-4">
                <button on:click={() => dialog.close()}>Cancel</button>
                <!-- svelte-ignore a11y-autofocus -->
                <button autofocus type="submit" class="block text-blue-600" on:click={modalHandler}>Create</button>
            </div>
        </form>
    </div>
</dialog>
