import { createList, createTodo, CreateTodoInput, fetchList, fetchLists, fetchTodoEvents, joinList, updateTodo, UpdateTodoInput } from "./api"
import { createAppSlice } from "./createAppSlice"
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
      async () => await fetchLists(),
      {
        fulfilled: (state, action) => {
          state.lists = action.payload
        },
        rejected: state => {
          state.lists = []
        },
      },
    ),
    createTodoListAsync: create.asyncThunk(
      async (name: string) => await createList(name),
      {
        fulfilled: (state, action) => {
          // state.lists = [action.payload, ...state.lists]
        },
      },
    ),
    joinTodoListAsync: create.asyncThunk(
      async (listID: string) => await joinList(listID),
      {
        fulfilled: (state, action) => {
          // state.lists = [action.payload, ...state.lists]
        },
      },
    ),
    resetTodoLists: create.reducer((state) => {
      state.lists = []
    }),
    loadActiveListAsync: create.asyncThunk(
      async (listID: string) => await fetchList(listID),
      {
        fulfilled: (state, action) => {
          state.activeList = action.payload
        },
        rejected: state => {
          state.activeList = null
        },
      },
    ),
    createTodoAsync: create.asyncThunk(
      async (createTodoInput: CreateTodoInput) => await createTodo(createTodoInput),
      {
        fulfilled: (state, action) => {
          // state.activeTodo = action.payload
        },
        rejected: state => {
          // state.activeTodo = null
        },
      },
    ),
    updateTodoAsync: create.asyncThunk(
      async (updateTodoInput: UpdateTodoInput) => await updateTodo(updateTodoInput),
      {
        fulfilled: (state, action) => {
          // state.activeTodo = action.payload
        },
        rejected: state => {
          // state.activeTodo = null
        },
      },
    ),
    resetActiveList: create.reducer((state) => {
      state.activeList = null
    }),
    loadActiveTodoEventsAsync: create.asyncThunk(
      async (todoID: string) => await fetchTodoEvents(todoID),
      {
        fulfilled: (state, action) => {
          state.activeTodoEvents = action.payload
        },
        rejected: state => {
          // state.activeTodo = null
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
