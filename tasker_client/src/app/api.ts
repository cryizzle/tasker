import axios from "axios"
import { Todo, TodoEvent, TodoList, User } from "./types"

export type ErrorResponse = { error: string }

const login = async (email: string): Promise<User> => {
  return axios.post("/user/login", { email }).then((response) => response.data.user)
}

const createList = async (name: string): Promise<TodoList> => {
  return axios.post("/list/create", { name }).then((response) => response.data.todo_list)
}

const joinList = async (listID: string): Promise<TodoList> => {
  return axios.post(`/list/join/${listID}`).then((response) => response.data.todo_list)
}

const fetchLists = async (): Promise<TodoList[]> => {
  return axios.get("/list/all").then((response) => response.data.todo_lists)
}

const fetchList = async (listID: string): Promise<TodoList> => {
  return axios.get(`/list/${listID}`).then((response) => response.data.todo_list)
}

export type CreateTodoInput = { todo_list_id: string, description: string }
const createTodo = async (input: CreateTodoInput): Promise<Todo> => {
  return axios.post("/todo/create", input).then((response) => response.data.todo)
}

export type UpdateTodoInput = { todoID: string, status: string }
const updateTodo = async ({ todoID, status }: UpdateTodoInput): Promise<Todo> => {
  return axios.post(`/todo/update/${todoID}`, { status }).then((response) => response.data.todo)
}

const fetchTodoEvents = async (todoID: string): Promise<TodoEvent[]> => {
  return axios.get(`/todo/events/${todoID}`).then((response) => response.data.todo_events)
}

export default {
  login,
  createList,
  joinList,
  fetchLists,
  fetchList,
  createTodo,
  updateTodo,
  fetchTodoEvents,
}


