import { writable } from "svelte/store";

import { todo } from "../wailsjs/go/models"

export const showTaskModal = writable(false);
export const showListModal = writable(false);
export const displayMode = writable(0);
export const lists = writable([] as todo.List[]);
export const tasks = writable([] as todo.Task[]);