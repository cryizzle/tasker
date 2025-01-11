import React, { useEffect, useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { loginAsync } from '../app/userSlice';
import { useAppDispatch } from '../app/hooks';
import { Container } from '@mui/material';
import { resetTodoLists } from '../app/todoListsSlice';

const UserLogin: React.FC = () => {
  const [email, setEmail] = useState('a@b.com');
  const [error, setError] = useState('');
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(resetTodoLists());
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();

    // Simple email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setError('Please enter a valid email address.');
      return;
    }

    setError(''); // Clear error if validation passes
    dispatch(loginAsync(email));
  };

  return (
    <Container
      maxWidth="sm"
      component="form"
      onSubmit={handleSubmit}
      sx={{
        textAlign: 'center',
      }}
    >
      <TextField
        label="Email Address"
        variant="outlined"
        fullWidth
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        error={!!error}
        helperText={error}
        sx={{ mb: 2 }}
      />
      <Button type="submit" variant="contained" color="primary" fullWidth>
        Login
      </Button>
    </Container>
  );
};

export default UserLogin;
