import React, { useEffect, useState } from 'react';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import { useAppDispatch, useAppSelector } from '../app/hooks';
import { createTodoListAsync, joinTodoListAsync, loadTodoListsAsync, resetActiveList, selectLists } from '../app/todoListsSlice';
import { useNavigate } from 'react-router';
import { Box, Button, Container, IconButton, ListItem, Paper, TextField, Tooltip } from '@mui/material';
import { Check, ContentCopy, NoteAdd } from '@mui/icons-material';
import { TodoList } from '../app/types';

const AllLists: React.FC = () => {
  const dispatch = useAppDispatch();
  const todoLists = useAppSelector(selectLists);

  useEffect(() => {
    dispatch(loadTodoListsAsync());
  }, [dispatch]);

  const handleCreateList = async (listName: string) => {
    await dispatch(createTodoListAsync(listName));
  };

  const handleJoinList = async (listToken: string) => {
    await dispatch(joinTodoListAsync(listToken));
  };

  return (
    <Container
      maxWidth="md"
    >
      <JoinList handleJoinList={handleJoinList} />
      <Box sx={{ m: 3, textAlign: 'center' }}>
        <h2>My Lists</h2>
      </Box>
      <List component={Box} bgcolor={'transparent'}>
        {todoLists.map((list) => (
          <SingleList list={list} key={list.id} />
        ))}
        <NewList handleCreateList={handleCreateList} />
      </List>
    </Container>
  );
};

const SingleList: React.FC<{
  list: TodoList
}> = ({ list }) => {

  const [copied, setCopied] = useState(false);
  const navigate = useNavigate();

  const handleListClick = (id: string) => {
    navigate(`/list/${id}`);
  };

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(list.token);
      setCopied(true);
      setTimeout(() => setCopied(false), 1500); // Reset after 1.5 seconds
    } catch (err) {
      console.error("Failed to copy text", err);
    }
  };

  return <ListItem component={Paper} elevation={2} sx={{ mb: 2, p: 2 }}>
    <Box sx={{ flexGrow: 1, cursor: "pointer" }} onClick={() => handleListClick(list.id)}>
      <Typography variant="h6" gutterBottom>
        {list.name}
      </Typography>
    </Box>
    <Tooltip title={copied ? "Copied!" : "Copy share token"} arrow placement="top">
      <IconButton color="primary" onClick={handleCopy} size='small'>
        {copied ? <Check /> : <ContentCopy />}
      </IconButton>
    </Tooltip>
  </ListItem>;
}

const NewList: React.FC<{
  handleCreateList: (listName: string) => void
}> = ({
  handleCreateList,
}) => {

    const [listName, setListName] = useState('');

    const onCreateList = async () => {
      await handleCreateList(listName.trim());
      setListName(''); // Clear input after dispatch
    }

    return <ListItem component={Paper} elevation={2} sx={{ mb: 2, p: 2 }}>
      <TextField
        label="List Name"
        variant="outlined"
        value={listName}
        onChange={(e) => setListName(e.target.value)}
        sx={{ flexGrow: 1 }}
      />
      <Button
        variant="contained"
        color="primary"
        onClick={onCreateList}
        startIcon={<NoteAdd />}
        sx={{ ml: 2 }}
        disabled={!listName.trim()}
      >
        Create List
      </Button>
    </ListItem>

  }

const JoinList: React.FC<{
  handleJoinList: (listToken: string) => void
}> = ({ handleJoinList }) => {
  const [listToken, setListToken] = useState('');

  const onJoinList = async () => {
    if (!listToken.trim()) {
      return;
    }
    await handleJoinList(listToken);
    setListToken(''); // Clear input after dispatch
  };

  return (
    <ListItem component={Paper} elevation={2} sx={{ mb: 2, p: 2 }}>
      <TextField
        label="List Token"
        variant="outlined"
        value={listToken}
        onChange={(e) => setListToken(e.target.value)}
        sx={{ flexGrow: 1 }}
      />
      <Button
        variant="contained"
        color="secondary"
        onClick={onJoinList}
        startIcon={<NoteAdd />}
        sx={{ ml: 2 }}
        disabled={!listToken.trim()}
      >
        Join List
      </Button>
    </ListItem>
  );
};


export default AllLists;
