import React, { useState, useEffect } from 'react'
import todoService from './services/todos'
import { NewTodo, Todo } from './types'

const App = () => {
  const [todos, setTodos] = useState<Todo[]>([])
  const [todoTitle, setTodoTitle] = useState<string>("")

  useEffect(() => {
    (async () => {
      const initialTodos = await todoService.getAll()
      const collator = new Intl.Collator(undefined, { numeric: true })
      initialTodos.sort((a, b) => collator.compare(a.id, b.id))
      setTodos(initialTodos)
    })()
  }, [])

  const addTodo = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    const newTodo: NewTodo = { title: todoTitle, done: false }
    try {
      const returnedTodo = await todoService.create(newTodo)
      setTodos(todos.concat(returnedTodo))
      setTodoTitle("")
    } catch (e) {
      console.log(e?.response?.data?.message)
    }
  }

  const toggleDone = async (index: number) => {
    const todoToUpdate = { ...todos[index] }
    todoToUpdate.done = !todoToUpdate.done
    try {
      const returnedTodo: Todo = await todoService.update(todoToUpdate)
      const newTodos = todos.map(t => t.id === returnedTodo.id ? returnedTodo : t)
      setTodos(newTodos)
    } catch (e) {
      console.log(e?.response?.data?.message)
    }
  }

  return (
    <div>
      <h3>Todo app (GitOps update, go!)</h3>
      <img src={BACKEND_URL + '/image.jpg'} />
      <form onSubmit={addTodo}>
        <input
          value={todoTitle}
          onChange={(event) => {
            setTodoTitle(event.target.value.slice(0, event.target.maxLength))
          }}
          maxLength={140}
        />
        <button type="submit">Add Todo</button>
      </form>
      <ul>
        {todos.map((t, i) =>
          <li key={t.id}>
            {t.title}
            <button onClick={() => toggleDone(i)}>
              {t.done ? <span>✓</span> : <span>❌</span>}
            </button>
          </li>
        )}
      </ul>
    </div>
  )
}

export default App