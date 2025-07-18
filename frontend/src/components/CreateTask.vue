<script setup lang="ts">
import { ref } from "vue";
import { createTask } from '@/api/task';

const emit = defineEmits(['task-created']);
const title = ref('');

async function addTask() {
    if (!title.value.trim()) return; // Ne rien faire si vide
    const task = await createTask(title.value);
    if (task === null) {
        console.error("Failed to create task");
    } else {
        emit('task-created', task);
        title.value = '';
    }
}
</script>

<template>
    <div class="w-full mx-auto mt-8 space-y-4">
        <h1 class="font-bold text-2xl">Add a new task</h1>
        <form @submit.prevent="addTask" class="flex gap-2">
            <input v-model="title" type="text" placeholder="Task title" class="flex-1 border rounded px-3 py-2" />
            <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
                Add
            </button>
        </form>
    </div>
</template>