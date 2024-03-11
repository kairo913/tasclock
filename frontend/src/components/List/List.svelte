<script lang="ts">
    import BxTrash from "svelte-boxicons/BxTrash.svelte";
    import { RemoveList } from "../../../wailsjs/go/todo/Todo";
    import type { todo } from "../../../wailsjs/go/models";
    import { lists } from "../../store";

    export let list: todo.List;
    export let selection: number[];

    const deleteList = async () => {
        if (list.id === 0) return;
        await RemoveList(list.id);
        lists.update((t) =>
            t.filter((item) => {
                return item.id !== list.id;
            }),
        );
    };
</script>

<li class="flex flex-row justify-between space-x-2 p-2">
    <input type="checkbox" value={list.id} class="rounded border-gray-300 bg-gray-100" bind:group={selection} />
    <div class="grow text-left">{list.title}</div>
    {#if list.id != 0}
        <button class="rounded-full hover:bg-slate-300" on:click={deleteList}>
            <BxTrash class="outline-none" />
        </button>
    {/if}
</li>
