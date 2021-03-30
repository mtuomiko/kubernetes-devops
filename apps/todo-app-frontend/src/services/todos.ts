import axios from 'axios'
import { Todo, NewTodo } from '../types'

const baseHost = BACKEND_URL
const baseUrl = `${baseHost}/todos`

const getAll = async (): Promise<Todo[]> => {
  const response = await axios.get(baseUrl)
  return response.data
}

const create = async (newTodo: NewTodo) => {
  const response = await axios.post(baseUrl, newTodo)
  return response.data
}

export default { getAll, create }