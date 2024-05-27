<script lang="ts">
    import type { List } from "./type.ts";

    export let list: List;
</script>

<div class="container">
    <p class="text-lg">{list.Title}</p>
    <slot />
</div>

<style lang="postcss">
    .container {
        @apply h-fit bg-gray-200 rounded-md p-2 flex flex-col justify-center items-center gap-2;
    }

    :global(body.dark) .container {
        @apply bg-gray-700;
    }
</style>