<script lang="ts">
    import Task from "./Task.svelte";
    import { tasks } from "../../store";
</script>

{#each $tasks as task}
    <Task {task} />
{/each}
