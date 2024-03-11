<script lang="ts">
    import Modal from "./Modal.svelte";
    import { showTaskModal, tasks } from "../../store";
    import { NewTask } from "../../../wailsjs/go/todo/Todo";

    const header: string = "Create New Task";

    const handler = async () => {
        if (title.length === 0) return false;
        const task = await NewTask(listid, title, starred, description, deadline, reward);
        tasks.update((t) => [...t, task]);
        listid = 0;
        title = "";
        starred = false;
        description = "";
        deadline = "";
        reward = 0;
        return true;
    };

    let listid: number = 0;
    let title: string = "";
    let starred: boolean = false;
    let description: string = "";
    let deadline: string = "";
    let reward: number = 0;
</script>

<Modal show={showTaskModal} {header} {handler}>
    <div class="">
        <input type="text" placeholder="Task Name" bind:value={title} />
        <input type="text" placeholder="Description" bind:value={description} />
        <input type="checkbox" placeholder="Star" bind:checked={starred} />
        <input type="number" placeholder="Reward" bind:value={reward} />
        <input type="date" placeholder="Deadline" bind:value={deadline} />
    </div>
</Modal>
