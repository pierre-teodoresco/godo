<script setup lang="ts">
import { ref, onMounted } from "vue";
import CreateTask from "./CreateTask.vue";
import { fetchTasks, updateTask, deleteTask } from '@/api/task';
import type { Task } from '@/types/task';
import { formatDate } from "@/lib/utils";

const tasks = ref<Task[]>([]);

onMounted(async () => {
    tasks.value = await fetchTasks();
})

function addTaskToList(newTask: Task) {
    tasks.value.push(newTask);
}

async function editTask(task: Task) {
    if (await updateTask(task.id, task.title, task.completed) === null) {
        console.error("Failed to update task");
    }
}

async function toggleCompleted(task: Task) {
    task.completed = !task.completed
    if (await updateTask(task.id, task.title, task.completed) === null) {
        console.error("Failed to complete task");
    }
}

async function removeTask(task: Task) {
    if (!await deleteTask(task.id)) {
        console.error("Failed to delete task");
    } else {
        tasks.value = tasks.value.filter(t => t.id !== task.id);
    }
}
</script>

<template>
    <CreateTask @task-created="addTaskToList" />
    <div class="w-full mx-auto mt-8 space-y-4">
        <div v-for="task in tasks" :key="task.id"
            class="bg-white shadow rounded p-4 flex items-center justify-between cursor-pointer"
            @click="toggleCompleted(task)">
            <div class="flex items-center gap-3">
                <input type="checkbox" v-model="task.completed" @change="editTask(task)" @click.stop
                    class="accent-blue-500 cursor-pointer" />
                <div>
                    <div class="font-semibold text-lg" :class="{ 'line-through text-gray-400': task.completed }">
                        {{ task.title }}
                    </div>
                    <div class="text-xs text-gray-500">{{ formatDate(task.created_at, "en-EN") }}</div>
                </div>
            </div>
            <div class="flex gap-2">
                <button class="cursor-pointer bg-transparent hover:bg-red-100 rounded transition-colors px-2 py-1"
                    @click.stop="removeTask(task)">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-6">
                        <path fill-rule="evenodd"
                            d="M16.5 4.478v.227a48.816 48.816 0 0 1 3.878.512.75.75 0 1 1-.256 1.478l-.209-.035-1.005 13.07a3 3 0 0 1-2.991 2.77H8.084a3 3 0 0 1-2.991-2.77L4.087 6.66l-.209.035a.75.75 0 0 1-.256-1.478A48.567 48.567 0 0 1 7.5 4.705v-.227c0-1.564 1.213-2.9 2.816-2.951a52.662 52.662 0 0 1 3.369 0c1.603.051 2.815 1.387 2.815 2.951Zm-6.136-1.452a51.196 51.196 0 0 1 3.273 0C14.39 3.05 15 3.684 15 4.478v.113a49.488 49.488 0 0 0-6 0v-.113c0-.794.609-1.428 1.364-1.452Zm-.355 5.945a.75.75 0 1 0-1.5.058l.347 9a.75.75 0 1 0 1.499-.058l-.346-9Zm5.48.058a.75.75 0 1 0-1.498-.058l-.347 9a.75.75 0 0 0 1.5.058l.345-9Z"
                            clip-rule="evenodd" />
                    </svg>

                </button>
            </div>
        </div>
    </div>
</template>