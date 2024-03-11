<script lang="ts">
    import type { todo } from "../../../wailsjs/go/models";
    import { RemoveTask, UpdateTask } from "../../../wailsjs/go/todo/Todo";
    import BxStar from "svelte-boxicons/BxStar.svelte";
    import BxCircle from "svelte-boxicons/BxCircle.svelte";
    import BxCheckCircle from "svelte-boxicons/BxCheckCircle.svelte";
    import BxsStar from "svelte-boxicons/BxsStar.svelte";
    import BxDetail from "svelte-boxicons/BxDetail.svelte";
    import BxTrash from "svelte-boxicons/BxTrash.svelte";
    import BxCalendar from "svelte-boxicons/BxCalendar.svelte";
    import BxYen from "svelte-boxicons/BxYen.svelte";
    import BxTimer from "svelte-boxicons/BxTimer.svelte";
    import BxListUl from "svelte-boxicons/BxListUl.svelte";
    import { tasks } from "../../store";

    export let task: todo.Task;

    const deleteTask = async () => {
        await RemoveTask(task.id);
        tasks.update((t) =>
            t.filter((item) => {
                return item.id !== task.id;
            }),
        );
    };

    $: UpdateTask(task);
</script>

<div class="min-w-80 max-w-80 space-y-2 rounded-md bg-slate-200 p-3">
    <div class="flex justify-between">
        <button class="rounded-full px-1 hover:bg-slate-300" on:click={() => (task.is_done = !task.is_done)}>
            {#if task.is_done}
                <BxCheckCircle class="shrink-0 outline-none" />
            {:else}
                <BxCircle class="shrink-0 outline-none" />
            {/if}
        </button>
        <div class="py-1">
            {task.title}
        </div>
        <button class="rounded-full px-1 hover:bg-slate-300" on:click={() => (task.starred = !task.starred)}>
            {#if task.starred}
                <BxsStar class="shrink-0 outline-none" />
            {:else}
                <BxStar class="shrink-0 outline-none" />
            {/if}
        </button>
    </div>
    {#if task.description != ""}
        <div class="flex justify-start px-1">
            <BxDetail class="mr-2 shrink-0 outline-none" />
            <p class="flex-auto break-all text-left">{task.description}</p>
        </div>
    {/if}
    {#if task.deadline != ""}
        <div class="flex justify-start px-1">
            <BxCalendar class="mr-2 shrink-0 outline-none" />
            <p class="flex-auto break-all text-left">{task.deadline}</p>
        </div>
    {/if}
    {#if task.reward != 0}
        <div class="flex justify-start px-1">
            <BxYen class="mr-2 shrink-0 outline-none" />
            <p class="flex-auto break-all text-left">{task.reward}</p>
        </div>
    {/if}
    <div class="flex justify-start px-1">
        <BxTimer class="mr-2 shrink-0 outline-none" />
        {task.elapsed}
    </div>
    {#if task.list_id != 0}
        <div class="flex justify-start px-1">
            <BxListUl class="mr-2 shrink-0 outline-none" />
            {task.list_id}
        </div>
    {/if}
    <div class="flex justify-end px-1">
        <button class="rounded-full hover:bg-slate-300" on:click={deleteTask}>
            <BxTrash class="outline-none" />
        </button>
    </div>
</div>
