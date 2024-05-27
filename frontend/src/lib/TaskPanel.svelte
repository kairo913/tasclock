<script lang="ts">
    import type { Task } from "./type.ts";
    import IconCircle from "@tabler/icons-svelte/IconCircle.svelte";
    import IconCheck from "@tabler/icons-svelte/IconCheck.svelte";

    export let task: Task;
</script>

<div class="container">
    <div class="relative">
        <p class="text-lg">{task.Title}</p>
        <button class="absolute right-0 top-0" on:click={() => (task.Completed = !task.Completed)}>
            {#if task.Completed}
                <IconCheck stroke={1.5} />
            {:else}
                <IconCircle stroke={1.5} />
            {/if}
        </button>
    </div>
    <p>{task.Description}</p>
</div>

<style lang="postcss">
    .container {
        @apply bg-gray-300 rounded-md p-2 w-full text-left;
    }

    :global(body.dark) .container {
        @apply bg-gray-700;
    }
</style>
