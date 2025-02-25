import { defineStore } from 'pinia'
import type { NewTodo, Todo } from '@/models/todo.model'
import todoService from '../services/todos'

export type TodoState = {
  todos: Todo[]
}

export const useTodos = defineStore('todoStore', {
  state: () => ({
    todos: []
  } as TodoState),
  actions: {
    async getTodos() {
      try {
        const allTodos = await todoService.getAll()
        this.todos = allTodos
      } catch (error) {
        console.error('Todos get failed: ', error)
      }
    },
    async addTodo(newTodo: NewTodo) {
      try {
        const createdTodo = await todoService.create(newTodo)
        this.todos.push(createdTodo)
      } catch (error) {
        console.error('Todo add failed: ', error)
      }
    },
  },
})
