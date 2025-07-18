import type { Task } from '@/types/task';

const API_URL = import.meta.env.VITE_API_URL;

export async function fetchTasks(): Promise<Task[]> {
    const url = `${API_URL}/tasks`;
    try {
        const response: Response = await fetch(url);
        if (!response.ok) {
            throw new Error(`${response.status} error while trying to fetch API`);
        }

        const tasks = await response.json() as Task[];
        return tasks;
    } catch (error) {
        console.error('Failed to fetch tasks:', error);
        return [];
    }
}

export async function createTask(title: string): Promise<Task | null> {
    const url = `${API_URL}/task`;
    try {
        const response: Response = await fetch(url, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title: title })
        });
        if (!response.ok) {
            throw new Error(`${response.status} error while trying to fetch API`);
        }

        const task = await response.json() as Task;
        return task;
    } catch (error) {
        console.error('Failed to fetch tasks:', error);
        return null;
    }
}

export async function updateTask(id: string, title: string, completed: boolean): Promise<Task | null> {
    const url = `${API_URL}/task`;
    try {
        const response: Response = await fetch(url, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id, title, completed })
        });
        if (!response.ok) {
            throw new Error(`${response.status} error while trying to fetch API`);
        }

        const task = await response.json() as Task;
        return task;
    } catch (error) {
        console.error('Failed to fetch tasks:', error);
        return null;
    }
}

export async function deleteTask(id: string): Promise<boolean> {
    const url = `${API_URL}/task`;
    try {
        const response: Response = await fetch(url, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id })
        });
        if (!response.ok) {
            throw new Error(`${response.status} error while trying to fetch API`);
        }

        return true;
    } catch (error) {
        console.error('Failed to fetch tasks:', error);
        return false;
    }
}