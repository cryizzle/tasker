import { Alert, Snackbar } from "@mui/material"
import { useAppDispatch, useAppSelector } from "../../app/hooks"
import { clearNotification, selectNotification } from "../../app/notificationSlice"

const NOTIFICATION_DURATION = 5000

const Notification: React.FC = () => {
  const notification = useAppSelector(selectNotification)
  const dispatch = useAppDispatch()

  const handleClose = () => {
    dispatch(clearNotification())
  }

  return (
    <Snackbar
      open={notification !== null}
      autoHideDuration={NOTIFICATION_DURATION}
      onClose={handleClose}
      anchorOrigin={{ vertical: "top", horizontal: "center" }}
    >
      <Alert onClose={handleClose} severity={notification?.severity}>
        {notification?.message ?? ""}
      </Alert>
    </Snackbar>
  )
}

export default Notification
