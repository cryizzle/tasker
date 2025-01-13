import { useEffect, useState } from "react"
import { useAppDispatch } from "../../app/hooks";
import { loadActiveListAsync } from "../../app/todoListsSlice";
import { BASE_URL } from "../../app/api";
import moment from "moment";
import { Typography } from "@mui/material";

const TodoListUpdater: React.FC<{
  todoListID: string;
  onUpdate: () => void;
}> = ({ todoListID, onUpdate }) => {
  const dispatch = useAppDispatch();
  const [timestamp, setTimestamp] = useState<moment.Moment>(moment());
  useEffect(() => {
    if (!todoListID) {
      return;
    }
    const eventSource = new EventSource(`${BASE_URL}/list/updates/${todoListID}`, {
      withCredentials: true,
    });
    eventSource.onmessage = (_) => {
      dispatch(loadActiveListAsync(todoListID));
      setTimestamp(moment());
      onUpdate();
    }
    return () => {
      eventSource.close();
    };
  }, [todoListID, dispatch, onUpdate]);

  return <Typography variant="body1" textAlign='right' mt={1}>Last Refreshed: {timestamp.toLocaleString()}</Typography>;
}

export default TodoListUpdater
