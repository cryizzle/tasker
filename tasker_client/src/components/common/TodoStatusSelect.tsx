import { MenuItem, Select, SelectChangeEvent } from "@mui/material";
import { GetPossibleNextTodoStatus, TodoStatus } from "../../app/types";

type TodoStatusSelectProps = {
  value: TodoStatus;
  handleChange?: (e: SelectChangeEvent<TodoStatus>) => void
  disabled?: boolean;
};

const TodoStatusSelect: React.FC<TodoStatusSelectProps> = ({ value, handleChange, disabled }: TodoStatusSelectProps) => {
  const possibleOptions = GetPossibleNextTodoStatus(value);

  return <Select
    value={value}
    onChange={handleChange}
    disabled={disabled}
    fullWidth
  >
    {Object.values(TodoStatus).map((status) => (
      <MenuItem key={status} value={status} disabled={!possibleOptions.includes(status) && status !== value}>
        {status}
      </MenuItem>
    ))}
  </Select>;
}

export default TodoStatusSelect;
