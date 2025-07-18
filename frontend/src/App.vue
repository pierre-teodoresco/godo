<script setup lang="ts">
import ListTask from './components/ListTask.vue'
</script>

<template>
    <div class="flex flex-col items-center mt-8 w-1/2 mx-auto">
        <div class="flex items-center gap-3 mb-6">
            <img src="@/assets/godo.png" alt="Godo logo" class="w-30 h-30" />
            <h1 class="text-6xl font-bold text-blue-700 tracking-tight">Godo</h1>
        </div>
        <Suspense>
            <ListTask />
        </Suspense>
    </div>
</template>

<style scoped>
.logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
}

.logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
    filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
