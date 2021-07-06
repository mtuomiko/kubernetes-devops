export interface Todo {
  id: string;
  title: string;
  done: boolean;
}

export type NewTodo = Omit<Todo, "id">