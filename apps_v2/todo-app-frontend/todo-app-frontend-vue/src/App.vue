<script setup lang="ts">
// import imageUrl from '../public/image.jpg'
import { useTodos } from './stores/todo';
import { nanoid } from 'nanoid'
import { onMounted, ref } from 'vue';
import type { NewTodo, Todo } from './models/todo.model';

// haxy, trying to use the image during dev but allow the final runtime go http server to serve the actual image
const imageUrl = "image.jpg"

const todosStore = useTodos()

const newTodoText = ref('')

const addTodo = () => {
  if (!newTodoText.value) {
    return
  }
  const newTodo: NewTodo = { name: newTodoText.value }
  todosStore.addTodo(newTodo)
  newTodoText.value = ''
}

onMounted(() => {
  todosStore.getTodos()
})

</script>

<template>
  <header>
    <div>
      <h2>Todo app</h2>
    </div>
    <div>
      <img alt="Image" class="logo" :src="imageUrl" width="400" height="400" />
    </div>
    <label>
      New todo:
      <input v-model="newTodoText" type="text" @keypress.enter="addTodo">
    </label>
    <button :disabled="!newTodoText" @click="addTodo">Add</button>
    <div><label>Todos:</label></div>
    <ul>
      <li v-for="todo in todosStore.todos" :key="todo.id">
        {{ todo.name }}
      </li>
    </ul>

  </header>

</template>
