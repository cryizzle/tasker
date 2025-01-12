import { AlertColor } from "@mui/material"
import { createSlice, PayloadAction } from "@reduxjs/toolkit"

export type Notification = {
  message: string
  severity: AlertColor
}

export interface NotificationState {
  notification: Notification | null
}

const initialState: NotificationState = {
  notification: null,
}

export const notificationSlice = createSlice({
  name: "notification",
  initialState,
  reducers: {
    // Handle API error responses
    setError: (state, action: PayloadAction<string>) => {
      const message = action.payload ?? "An unexpected error occurred";
      state.notification = { message, severity: "error" };
    },

    // Handle specific success messages
    setSuccess: (state, action: PayloadAction<string>) => {
      state.notification = { message: action.payload, severity: "success" };
    },

    // Clear the notification
    clearNotification: (state) => {
      state.notification = null;
    }
  },
  selectors: {
    selectNotification: state => state.notification,
  },
});

export const {
  setError,
  setSuccess,
  clearNotification,
} = notificationSlice.actions

export const {
  selectNotification,
} = notificationSlice.selectors
