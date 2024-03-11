<script lang="ts">
    import BxPlus from "svelte-boxicons/BxPlus.svelte";
    import BxTask from "svelte-boxicons/BxTask.svelte";
    import BxStar from "svelte-boxicons/BxStar.svelte";
    import BxChevronUp from "svelte-boxicons/BxChevronUp.svelte";
    import BxChevronDown from "svelte-boxicons/BxChevronDown.svelte";

    import { showListModal, showTaskModal, displayMode } from "../../store";
    import { slide } from "svelte/transition";
    import List from "../List/List.svelte";
    import AllList from "../List/AllList.svelte";
    import { todo } from "../../../wailsjs/go/models";

    let expanded = true;
    let selection = [0];
</script>

<button
    class="flex rounded-full bg-white p-3 shadow-md hover:text-sky-300 hover:shadow-blue-500/50"
    on:click={() => showTaskModal.set(true)}
>
    <BxPlus class="mr-2 outline-none" />
    Create
</button>

<div class="flex flex-col space-y-2 p-2">
    <button class="flex rounded-full p-3 hover:bg-slate-200" on:click={() => displayMode.set(0)}>
        <BxTask class="mr-2 outline-none" />
        All Tasks
    </button>
    <button class="flex rounded-full p-3 hover:bg-slate-200" on:click={() => displayMode.set(1)}>
        <BxStar class="mr-2 outline-none" />
        Starred
    </button>
</div>

<div class="flex flex-col space-y-2 overflow-y-auto">
    <button
        class="flex flex-row justify-between rounded-full p-2 hover:bg-slate-200"
        on:click={() => (expanded = !expanded)}
    >
        Lists
        {#if expanded}
            <BxChevronUp class="outline-none" />
        {:else}
            <BxChevronDown class="outline-none" />
        {/if}
    </button>
    {#if expanded}
        <ul transition:slide class="flex flex-col overflow-y-auto">
            <List {selection} list={new todo.List({ id: 0, title: "No List" })} />
            <AllList {selection} />
        </ul>
    {/if}
</div>

<div class="flex flex-col space-y-4 p-2">
    <button class="flex flex-row rounded-full p-3 hover:bg-slate-200" on:click={() => showListModal.set(true)}>
        <BxPlus class="mr-2 outline-none" />
        Create new list
    </button>
</div>
