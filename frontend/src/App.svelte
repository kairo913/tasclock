<script lang="ts">
    import ListModal from "./components/Modal/ListModal.svelte";
    import TaskModal from "./components/Modal/TaskModal.svelte";
    import Sidebar from "./components/Sidebar/Sidebar.svelte";
    import AllTask from "./components/Task/AllTask.svelte";
    import { tasks, lists } from "./store";
    import { Tasks, Lists } from "../wailsjs/go/todo/Todo";
    import { onMount } from "svelte";

    const fetchTasks = async () => {
        tasks.set([]);
        const items = await Tasks(100);
        if (items == null || items.length === 0) {
            return;
        }
        items.forEach((item) => {
            tasks.update((l) => [...l, item]);
        });
    };

    const fetchLists = async () => {
        lists.set([]);
        const items = await Lists(100);
        if (items == null || items.length === 0) {
            return;
        }
        items.forEach((item) => {
            lists.update((l) => [...l, item]);
        });
    };

    onMount(() => {
        fetchTasks();
        fetchLists();
    });
</script>

<main class="flex max-h-full min-h-full flex-row space-x-8 bg-slate-50 p-4 text-slate-900">
    <div class="flex w-64 flex-col content-center justify-center space-y-4">
        <Sidebar />
    </div>
    <div class="space-y-4 overflow-auto">
        <AllTask />
    </div>
</main>

<TaskModal />
<ListModal />
