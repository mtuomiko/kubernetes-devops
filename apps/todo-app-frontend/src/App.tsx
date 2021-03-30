import React, { useState, useEffect } from 'react'
import todoService from './services/todos'
import { NewTodo, Todo } from './types'

const App = () => {
  const [todos, setTodos] = useState<Todo[]>([])
  const [todoTitle, setTodoTitle] = useState<string>("")

  useEffect(() => {
    (async () => {
      const initialTodos = await todoService.getAll()
      setTodos(initialTodos)
    })()
  }, [])

  const addTodo = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    const newTodo: NewTodo = { title: todoTitle }
    try {
      const returnedTodo = await todoService.create(newTodo)
      setTodos(todos.concat(returnedTodo))
      setTodoTitle("")
    } catch (e) {
      console.log(e)
    }
  }

  return (
    <div>
      <h3>Todo app</h3>
      <img src={BACKEND_URL + '/image.jpg'} />
      <form onSubmit={addTodo}>
        <input
          value={todoTitle}
          onChange={(event) => {
            setTodoTitle(event.target.value)
          }}
        />
        <button type="submit">Add Todo</button>
      </form>
      <ul>
        {todos.map(t =>
          <li key={t.id}>{t.title}</li>
        )}
      </ul>
    </div>
  )
}

export default App