import axios from "axios"

export const login = async (email: string) => {
  try {
    return await axios.post("/user/login", { email }).then((response) => response.data.user)
  } catch (e) {
    return null
  }
}

export const createList = async (name: string) => {
  try {
    return await axios.post("/list/create", { name }).then((response) => response.data.todo_list)
  } catch (e) {
    return null
  }
}

export const joinList = async (listID: string) => {
  try {
    return await axios.post(`/list/join/${listID}`).then((response) => response.data.todo_list)
  } catch (e) {
    return null
  }
}

export const fetchLists = async () => {
  try {
    return await axios.get("/list/all").then((response) => response.data.todo_lists)
  } catch (e) {
    return []
  }
}

export const fetchList = async (listID: string) => {
  try {
    return await axios.get(`/list/${listID}`).then((response) => response.data.todo_list)
  } catch (e) {
    return null
  }
}

export type CreateTodoInput = { todo_list_id: string, description: string }
export const createTodo = async (input: CreateTodoInput) => {
  try {
    return await axios.post("/todo/create", input).then((response) => response.data.todo)
  } catch (e) {
    return null
  }
}

export type UpdateTodoInput = { todoID: string, status: string }
export const updateTodo = async ({ todoID, status }: UpdateTodoInput) => {
  try {
    return await axios.post(`/todo/update/${todoID}`, { status }).then((response) => response.data.todo)
  } catch (e) {
    return null
  }
}

export const fetchTodoEvents = async (todoID: string) => {
  try {
    return await axios.get(`/todo/events/${todoID}`).then((response) => response.data.todo_events)
  } catch (e) {
    return []
  }
}


