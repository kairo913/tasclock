<script lang="ts">
    import List from "./List.svelte";
    import { lists } from "../../store";

    export let selection: number[];
</script>

{#each $lists as list}
    <List {selection} {list} />
{/each}
