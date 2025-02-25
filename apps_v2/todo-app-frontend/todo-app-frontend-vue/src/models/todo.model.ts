export interface Todo {
  id: string;
  name: string;
}

export type NewTodo = Omit<Todo, "id">
