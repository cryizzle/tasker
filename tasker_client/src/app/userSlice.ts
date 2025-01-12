import api from "./api"
import { createAppSlice } from "./createAppSlice"
import { User } from "./types"

export interface UserState {
  activeUser: User | null
}

const initialState: UserState = {
  activeUser: null,
}

// If you are not using async thunks you can use the standalone `createSlice`.
export const userSlice = createAppSlice({
  name: "user",
  initialState,
  reducers: create => ({
    loginAsync: create.asyncThunk(
      async (email: string) => await api.login(email),
      {
        fulfilled: (state, action) => {
          state.activeUser = action.payload
        },
        rejected: state => {
          state.activeUser = null
        },
      },
    ),
    logout: create.reducer((state) => {
      state.activeUser = null
    }),
  }),

  selectors: {
    selectActiveUser: state => state.activeUser,
  },
})

export const {
  loginAsync,
  logout,
} = userSlice.actions

export const {
  selectActiveUser,
} = userSlice.selectors
