import { Breadcrumbs, Link, Typography } from "@mui/material";
import { useAppSelector } from "../../app/hooks";
import { selectActiveList } from "../../app/todoListsSlice";
import { useNavigate } from "react-router";


const TodoListBreadcrumb: React.FC = () => {

  const activeList = useAppSelector(selectActiveList);
  const navigate = useNavigate();

  return <Breadcrumbs sx={{ mb: 3, mt: 3, textAlign: "left"}}>
    <Link underline="hover" color="inherit" onClick={() => navigate("/")} href="#">
      Home
    </Link>
    <Typography color="textPrimary">{activeList?.name}</Typography>
  </Breadcrumbs>
};

export default TodoListBreadcrumb
