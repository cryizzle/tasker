import { AppBar, Button, Toolbar, Typography } from "@mui/material"
import { useNavigate } from "react-router"
import { logout } from "../../app/userSlice";
import { useAppDispatch } from "../../app/hooks";

const Navbar: React.FC = () => {

  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const handleLogout = () => {
    dispatch(logout())
    navigate("/")
  }

  return <AppBar position="static">
    <Toolbar>
      <Typography variant="h4" component="div" sx={{ flexGrow: 1 }}>
        Tasker
      </Typography>
      <Button color="inherit" onClick={handleLogout} size='medium'>
        Logout
      </Button>
    </Toolbar>
  </AppBar>
}

export default Navbar
