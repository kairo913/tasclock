<script lang="ts">
    import Modal from "./Modal.svelte";
    import { lists, showListModal } from "../../store";
    import { NewList } from "../../../wailsjs/go/todo/Todo";

    const header: string = "Create New List";

    const handler = async () => {
        if (title.length === 0) return false;
        const list = await NewList(title);
        lists.update((l) => [...l, list]);
        title = "";
        return true;
    };

    let title: string;
</script>

<Modal show={showListModal} {header} {handler}>
    <div>
        <input type="text" placeholder="List Name" bind:value={title} />
    </div>
</Modal>
