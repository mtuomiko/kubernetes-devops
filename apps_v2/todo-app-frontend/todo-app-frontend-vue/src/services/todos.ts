import type { NewTodo, Todo } from '@/models/todo.model'
import axios, { type AxiosResponse } from 'axios'

const baseUrl = import.meta.env.VITE_TODO_API_URL

const serviceUrl = baseUrl + '/todos'


const apiClient = axios.create({
  baseURL: serviceUrl,
});

interface TodosResponse {
  todos: Todo[];
}

interface NewTodoResponse {
  todo: Todo;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const apiRequest = async <T>(url: string, method: 'GET' | 'POST' | 'PUT' | 'DELETE', data?: any): Promise<T> => {
  const response: AxiosResponse<T> = await apiClient({
    method,
    url,
    data,
  });

  return response.data;
};

const getAll = async (): Promise<Todo[]> => {
  const response = await apiRequest<TodosResponse>('', 'GET')
  return response.todos
}

const create = async (newTodo: NewTodo): Promise<Todo> => {
  const response = await apiRequest<NewTodoResponse>('', 'POST', newTodo)
  return response.todo
}


export default { getAll, create }
