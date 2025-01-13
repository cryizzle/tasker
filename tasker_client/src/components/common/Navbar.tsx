import { AppBar, Box, Button, Toolbar, Typography } from "@mui/material"
import { useNavigate } from "react-router"
import { logout, selectActiveUser } from "../../app/userSlice";
import { useAppDispatch, useAppSelector } from "../../app/hooks";

const Navbar: React.FC = () => {

  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const user = useAppSelector(selectActiveUser);
  const handleLogout = () => {
    dispatch(logout())
    navigate("/")
  }

  return <AppBar position="static">
    <Toolbar>
      <Typography variant="h4" sx={{ flexGrow: 1 }}>
        Tasker
      </Typography>
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <Typography variant="body1" sx={{ mr: 2}}>
          <strong>{user?.email}</strong>
        </Typography>
        <Button color="inherit" onClick={handleLogout} size='medium'>
          Logout
        </Button>
      </Box>
    </Toolbar>
  </AppBar>
}

export default Navbar
