<script lang="ts">
    import "../app.css";
    import IconMenu2 from "@tabler/icons-svelte/IconMenu2.svelte";
    import IconSun from "@tabler/icons-svelte/IconSun.svelte";
    import IconMoon from "@tabler/icons-svelte/IconMoon.svelte";
    import Sidebar from "$lib/Sidebar.svelte";

    import ListPanel from "$lib/ListPanel.svelte";
    import TaskPanel from "$lib/TaskPanel.svelte";
    import type { List, Task } from "$lib/type.ts";

    let darkMode = false;
    let fetched = true;

    let tasks: Task[] = [
        {
            Id: 1,
            ListId: 1,
            Title: "Task 1",
            Description: "Description 1",
            Completed: false,
            CreatedAt: "2021-10-01T00:00:00Z",
            UpdatedAt: "2021-10-01T00:00:00Z",
        },
        {
            Id: 2,
            ListId: 2,
            Title: "Task 2",
            Description: "Description 2",
            Completed: true,
            CreatedAt: "2021-10-01T00:00:00Z",
            UpdatedAt: "2021-10-01T00:00:00Z",
        },
        {
            Id: 3,
            ListId: 1,
            Title: "Task 3",
            Description: "Description 3",
            Completed: false,
            CreatedAt: "2021-10-01T00:00:00Z",
            UpdatedAt: "2021-10-01T00:00:00Z",
        }
    ];
    let lists: List[] = [
        {
            Id: 1,
            Title: "List 1",
            Description: "Description 1",
            CreatedAt: "2021-10-01T00:00:00Z",
            UpdatedAt: "2021-10-01T00:00:00Z",
        },
        {
            Id: 2,
            Title: "List 2",
            Description: "Description 2",
            CreatedAt: "2021-11-01T00:00:00Z",
            UpdatedAt: "2021-11-01T00:00:00Z",
        },
    ];

    function toggleDarkMode() {
        if (darkMode) {
            document.body.classList.remove("dark");
            darkMode = false;
        } else {
            document.body.classList.add("dark");
            darkMode = true;
        }
    }
</script>

<div class="header">
    <button class="absolute left-6">
        <IconMenu2 />
    </button>
    <p class="text-3xl font-bold">Tasclock</p>
    <button class="absolute right-6" on:click={toggleDarkMode}>
        {#if darkMode}
            <IconMoon />
        {:else}
            <IconSun />
        {/if}
    </button>
</div>
<main>
    <Sidebar bind:lists={lists} />
    {#if fetched}
        {#each lists as list}
            <ListPanel {list}>
                {#each tasks.filter((task) => task.ListId == list.Id) as task}
                    <TaskPanel {task} />
                {/each}
            </ListPanel>
        {/each}
    {:else}
        <p>Loading...</p>
    {/if}
</main>

<style lang="postcss">
    main {
        @apply flex flex-row justify-center gap-2 py-2;
    }

    .header {
        @apply bg-gray-200 rounded-md p-1 flex flex-row justify-center items-center;
    }

    :global(body.dark) .header {
        @apply bg-gray-700;
    }
</style>
