export type TodoList = {
  id: string
  name: string
  token: string
  members: Membership[]
  todos: Todo[]
}

export enum TodoStatus {
  TODO = "TODO",
  ONGOING = "ONGOING",
  DONE = "DONE",
}

export const GetPossibleNextTodoStatus = (status: TodoStatus): TodoStatus[] => {
  switch (status) {
    case TodoStatus.TODO:
      return [TodoStatus.ONGOING]
    case TodoStatus.ONGOING:
      return [TodoStatus.TODO, TodoStatus.DONE]
    case TodoStatus.DONE:
      return [TodoStatus.ONGOING]
  }
}

export type Todo = {
  id: string
  todo_list_id: string
  description: string
  status: TodoStatus
  created_at: string
  updated_at: string
  created_by: string
  email: string
}

export enum TodoEventType {
  TODO_CREATED = "TODO_CREATED",
  TODO_STATUS_CHANGED = "TODO_STATUS_CHANGED",
}

export type TodoEvent = {
  id: string
  todoID: string
  old_value: string
  new_value: string
  created_at: string
  email: string
  event_type: TodoEventType
}

export type User = {
  id: string
  email: string
}

export enum MembershipType {
  OWNER = "OWNER",
  MEMBER = "MEMBER",
}

export type Membership = {
  user_id: string
  email: string
  todo_list_id: string
  type: MembershipType
}

export const FormatTodoEvent = (event: TodoEvent) => {
  switch (event.event_type) {
    case TodoEventType.TODO_CREATED:
      return `Todo created by ${event.email}`
    case TodoEventType.TODO_STATUS_CHANGED:
      return `Status changed from ${event.old_value} to ${event.new_value} by ${event.email}`
  }
}
