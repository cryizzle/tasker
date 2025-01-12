import axios from "axios"

export type ErrorResponse = { error: string }

const login = async (email: string) => {
  return await axios.post("/user/login", { email }).then((response) => response.data.user)
}

const createList = async (name: string) => {
  return await axios.post("/list/create", { name }).then((response) => response.data.todo_list)
}

const joinList = async (listID: string) => {
  return await axios.post(`/list/join/${listID}`).then((response) => response.data.todo_list)
}

const fetchLists = async () => {
  return await axios.get("/list/all").then((response) => response.data.todo_lists)
}

const fetchList = async (listID: string) => {
  return await axios.get(`/list/${listID}`).then((response) => response.data.todo_list)
}

export type CreateTodoInput = { todo_list_id: string, description: string }
const createTodo = async (input: CreateTodoInput) => {
  return await axios.post("/todo/create", input).then((response) => response.data.todo)
}

export type UpdateTodoInput = { todoID: string, status: string }
const updateTodo = async ({ todoID, status }: UpdateTodoInput) => {
  return await axios.post(`/todo/update/${todoID}`, { status }).then((response) => response.data.todo)
}

const fetchTodoEvents = async (todoID: string) => {
  return await axios.get(`/todo/events/${todoID}`).then((response) => response.data.todo_events)
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


