import { AxiosError } from "axios"
import api, { CreateTodoInput, ErrorResponse, UpdateTodoInput } from "./api"
import { createAppSlice } from "./createAppSlice"
import { setError, setSuccess } from "./notificationSlice"
import { TodoEvent, TodoList } from "./types"

export interface TodoListState {
  lists: TodoList[]
  activeList: TodoList | null
  activeTodoEvents: TodoEvent[]
}

const initialState: TodoListState = {
  lists: [],
  activeList: null,
  activeTodoEvents: [],
}

// If you are not using async thunks you can use the standalone `createSlice`.
export const todoListSlice = createAppSlice({
  name: "todoList",
  initialState,
  reducers: create => ({
    loadTodoListsAsync: create.asyncThunk(
      async (_, { dispatch }) => {
        try {
          return await api.fetchLists()
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
      {
        fulfilled: (state, action) => {
          state.lists = action.payload
        },
        rejected: (state) => {
          state.lists = []
        },
      },
    ),
    createTodoListAsync: create.asyncThunk(
      async (name: string, { dispatch }) => {
        try {
          const response = await api.createList(name)
          dispatch(loadTodoListsAsync())
          dispatch(setSuccess("Successfully created todo list"))
          return response
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      }
    ),
    joinTodoListAsync: create.asyncThunk(
      async (listID: string, { dispatch }) => {
        try {
          const response = await api.joinList(listID)
          dispatch(loadTodoListsAsync())
          dispatch(setSuccess("Successfully joined todo list"))
          return response
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
    ),
    resetTodoLists: create.reducer((state) => {
      state.lists = []
    }),
    loadActiveListAsync: create.asyncThunk(
      async (listID: string, { dispatch }) => {
        try {
          return await api.fetchList(listID)
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
      {
        fulfilled: (state, action) => {
          state.activeList = action.payload
        },
        rejected: (state) => {
          state.activeList = null
        },
      },
    ),
    createTodoAsync: create.asyncThunk(
      async (createTodoInput: CreateTodoInput, { dispatch }) => {
        try {
          const createdTodo = await api.createTodo(createTodoInput)
          dispatch(loadActiveListAsync(createdTodo.todo_list_id))
          dispatch(setSuccess("Successfully created todo"))
          return createdTodo
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
    ),
    updateTodoAsync: create.asyncThunk(
      async (updateTodoInput: UpdateTodoInput, { dispatch }) => {
        try {
          const updatedTodo = await api.updateTodo(updateTodoInput)
          dispatch(loadActiveListAsync(updatedTodo.todo_list_id))
          dispatch(setSuccess("Successfully updated todo"))
          return updatedTodo
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
    ),
    resetActiveList: create.reducer((state) => {
      state.activeList = null
    }),
    loadActiveTodoEventsAsync: create.asyncThunk(
      async (todoID: string, { dispatch }) => {
        try {
          return await api.fetchTodoEvents(todoID)
        } catch (e) {
          const message = (e as AxiosError<ErrorResponse>).response?.data.error
          dispatch(setError(message))
        }
      },
      {
        fulfilled: (state, action) => {
          state.activeTodoEvents = action.payload
        },
      },
    ),
  }),
  selectors: {
    selectLists: todoListState => todoListState.lists,
    selectActiveList: todoListState => todoListState.activeList,
    selectActiveTodoEvents: todoListState => todoListState.activeTodoEvents,
  },
})

export const {
  loadTodoListsAsync,
  createTodoListAsync,
  joinTodoListAsync,
  loadActiveListAsync,
  createTodoAsync,
  updateTodoAsync,
  loadActiveTodoEventsAsync,
  resetTodoLists,
  resetActiveList,
} = todoListSlice.actions

export const {
  selectLists,
  selectActiveList,
  selectActiveTodoEvents
} = todoListSlice.selectors
