<script lang="ts">
    import IconPlus from "@tabler/icons-svelte/IconPlus.svelte";
    import IconCircleCheck from "@tabler/icons-svelte/IconCircleCheck.svelte";
    import IconStar from "@tabler/icons-svelte/IconStar.svelte";
    import IconChevronDown from "@tabler/icons-svelte/IconChevronDown.svelte";
    import IconChevronUp from "@tabler/icons-svelte/IconChevronUp.svelte";

    import type { List } from "./type.ts";
    export let lists: List[];

    let expanded = true;
</script>

<div class="sidebar">
    <button class="full">
        <IconPlus />
        Create
    </button>
    <button>
        <IconCircleCheck stroke={1.5} />
        All Tasks
    </button>
    <button>
        <IconStar stroke={1.5} />
        Starred
    </button>
    <div>
        Lists
        <button on:click={() => expanded = !expanded}>
            {#if expanded}
                <IconChevronDown />
            {:else}
                <IconChevronUp />
            {/if}
        </button>
    </div>
    {#if expanded}
        {#each lists as list}
            <button>
                <IconCircleCheck stroke={1.5} />
                {list.Title}
            </button>
        {/each}
    {/if}
</div>

<style lang="postcss">
    .sidebar {
        @apply min-w-[250px] max-w-[250px] flex flex-col justify-start items-center gap-2 p-2;
    }

    .sidebar button {
        @apply w-full flex flex-row p-1.5 gap-2 rounded-full;
    }

    .sidebar button:hover {
        @apply bg-gray-300;
    }

    .sidebar .full {
        @apply p-2 border-2 bg-gray-50;
    }

    .sidebar .full:hover {
        @apply bg-gray-200;
    }

    :global(body.dark) .sidebar button:hover {
        @apply bg-gray-700;
    }

    :global(body.dark) .sidebar .full {
        @apply bg-gray-800;
    }

    :global(body.dark) .sidebar .full:hover {
        @apply bg-gray-600;
    }

    .sidebar div {
        @apply w-full text-lg flex flex-row justify-between items-center;
    }

    .sidebar div button {
        @apply w-fit grow-0 p-1.5 rounded-full;
    }
</style>