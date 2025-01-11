import { BrowserRouter, Route, Routes } from "react-router"
import { ROUTES } from "./resources/route-constants"
import HomePage from "./pages/HomePage"
import { useAppSelector } from "./app/hooks"
import { selectActiveUser } from "./app/userSlice"
import LoginPage from "./pages/LoginPage"
import TodoListPage from "./pages/TodoListPage"

const AppRouter = () => {

  const activeUser = useAppSelector(selectActiveUser)
  return (
      <BrowserRouter>
        <Routes>
          <Route index element={activeUser == null ? <LoginPage /> : <HomePage/>} />
          <Route path={ROUTES.HOME_ROUTE} element={<HomePage />} />
          <Route path={ROUTES.LOGIN_ROUTE} element={<LoginPage />} />
          <Route path={ROUTES.TODO_LIST_ROUTE} element={<TodoListPage />} />
        </Routes>
      </BrowserRouter>
  )
}

export default AppRouter
