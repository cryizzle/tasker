import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Typography from '@mui/material/Typography';
import Tooltip from '@mui/material/Tooltip';
import { useAppDispatch, useAppSelector } from '../app/hooks';
import { FormatTodoEvent, Todo, TodoStatus } from '../app/types';
import { createTodoAsync, loadActiveListAsync, loadActiveTodoEventsAsync, selectActiveList, selectActiveTodoEvents, updateTodoAsync } from '../app/todoListsSlice';
import { Box, Collapse, Container, IconButton, Paper, TextField } from '@mui/material';
import moment from 'moment';
import { Add, Visibility } from '@mui/icons-material';
import TodoStatusSelect from './common/TodoStatusSelect';
import TodoListUpdater from './common/TodoListUpdater';

const TodoList: React.FC = () => {
  const { todoListID } = useParams<{ todoListID: string }>();
  const [openRowTodoID, setOpenRowTodoID] = useState<string | null>(null);

  const dispatch = useAppDispatch();
  const activeList = useAppSelector(selectActiveList);

  useEffect(() => {
    if (todoListID) {
      dispatch(loadActiveListAsync(todoListID));
    }
  }, [todoListID]);

  const handleAddTodo = async (description: string) => {
    if (activeList) {
      await dispatch(createTodoAsync({ todo_list_id: activeList.id, description }))
    }
  };

  const handleStatusChange = async (todoID: string, status: TodoStatus) => {
    setOpenRowTodoID(null);
    await dispatch(updateTodoAsync({ todoID, status }));
  };

  const handleRowClick = async (openRowTodoID: string) => {
    await dispatch(loadActiveTodoEventsAsync(openRowTodoID));
    setOpenRowTodoID(prevOpenRowId => (prevOpenRowId === openRowTodoID ? null : openRowTodoID));
  };

  const getListOwner = () => {
    return activeList?.members.find((member) => member.type === "OWNER")?.email;
  }

  if (!activeList) {
    return <Typography variant="h6">Loading...</Typography>;
  }

  return (
    <>
      <Container
        maxWidth="lg"
      >
        <Typography variant="h4" component="h1" gutterBottom textAlign="center">
          {activeList.name}
        </Typography>
        <Typography variant="body1" gutterBottom textAlign="right">
          List Owner: {getListOwner()}
        </Typography>
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Description</TableCell>
                <TableCell width="20%">Updated At</TableCell>
                <TableCell width="20%">Status</TableCell>
                <TableCell width="10%">Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {activeList.todos?.map((todo) => (
                <TodoItem
                  key={todo.id}
                  todo={todo}
                  handleStatusChange={handleStatusChange}
                  isOpen={openRowTodoID === todo.id}
                  handleRowOpen={() => handleRowClick(todo.id)}
                />
              ))}
              <NewTodoItem onAdd={handleAddTodo} />
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
      <TodoListUpdater todoListID={todoListID} onUpdate={() => setOpenRowTodoID(null)} />
    </>
  );
};

const TodoItem: React.FC<{
  todo: Todo;
  handleStatusChange: (todoID: string, status: TodoStatus) => void;
  handleRowOpen: () => void;
  isOpen: boolean;
}> = ({ todo, handleStatusChange, handleRowOpen, isOpen }) => {

  return <>
    <TableRow>
      <TableCell>{todo.description}</TableCell>
      <TableCell>
        <Tooltip title={moment(todo.updated_at).toLocaleString()} placement='top'>
          <span>{moment(todo.updated_at).fromNow()}</span>
        </Tooltip>
      </TableCell>
      <TableCell>
        <TodoStatusSelect value={todo.status} handleChange={(e) => handleStatusChange(todo.id, e.target.value as TodoStatus)} />
      </TableCell>
      <TableCell>
        <Tooltip title="View Todo Details" placement='top'>
          <IconButton color="primary" onClick={handleRowOpen}>
            <Visibility />
          </IconButton>
        </Tooltip>
      </TableCell>
    </TableRow>
    <TableRow>
      <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={4}>
        <Collapse in={isOpen} timeout="auto" unmountOnExit>
          <TodoDetails todo={todo} />
        </Collapse>
      </TableCell>
    </TableRow>
  </>

}

const TodoDetails: React.FC<{ todo: Todo }> = ({ todo }) => {
  const activeTodoEvents = useAppSelector(selectActiveTodoEvents);

  return <Box p={2}>
    Created At: {moment(todo.created_at).toLocaleString()}
    < br />
    Owned by {todo.email}
    < br />
    <Typography variant="h6" gutterBottom component="div">
      History
    </Typography>
    <Table>
      <TableHead>
        <TableRow>
          <TableCell width="20%">Date</TableCell>
          <TableCell>Event</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {activeTodoEvents.map((event) => (
          <TableRow key={event.id}>
            <TableCell>
              <Tooltip title={moment(event.created_at).toLocaleString()} placement='top'>
                <span>{moment(event.created_at).fromNow()}</span>
              </Tooltip>
            </TableCell>
            <TableCell>{FormatTodoEvent(event)}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </Box>
}

const NewTodoItem: React.FC<{ onAdd: (description: string) => void }> = ({ onAdd }) => {
  const [description, setDescription] = useState('');

  const handleAddClick = () => {
    if (description.trim()) {
      onAdd(description);
      setDescription('');
    }
  };

  return (
    <TableRow>
      <TableCell colSpan={2}>
        <TextField
          fullWidth
          label="New Todo Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          variant="outlined"
          size="medium"
        />
      </TableCell>
      <TableCell>
        <TodoStatusSelect value={TodoStatus.TODO} disabled />
      </TableCell>
      <TableCell>
        <Tooltip title="Create New Todo">
          <IconButton onClick={handleAddClick} disabled={!description.trim()} color="secondary">
            <Add />
          </IconButton>
        </Tooltip>
      </TableCell>
    </TableRow>
  );
};

export default TodoList;
