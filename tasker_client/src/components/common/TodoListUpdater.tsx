import { useEffect } from "react"
import { useAppDispatch } from "../../app/hooks";
import { loadActiveListAsync } from "../../app/todoListsSlice";

const TodoListUpdater: React.FC<{
  todoListID: string;
  onUpdate: () => void;
}> = ({ todoListID, onUpdate }) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    if (!todoListID) {
      return;
    }
    const eventSource = new EventSource(`http://localhost:8000/list/updates/${todoListID}`, {
      withCredentials: true,
    });
    eventSource.onmessage = (_) => {
      dispatch(loadActiveListAsync(todoListID));
      onUpdate();
    }
    return () => {
      eventSource.close();
    };
  }, [todoListID, dispatch, onUpdate]);

  return null;
}

export default TodoListUpdater
