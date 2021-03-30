export interface Todo {
  id: string;
  title: string;
}

export type NewTodo = Omit<Todo, "id">