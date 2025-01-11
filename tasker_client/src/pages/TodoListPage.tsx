import type React from 'react'
import TodoList from '../components/TodoList'
import Navbar from '../components/common/Navbar'
import TodoListBreadcrumb from '../components/common/TodoListBreadcrumb'

const TodoListPage: React.FC = () => {
    return (
        <>
            <Navbar />
            <TodoListBreadcrumb />
            <TodoList />
        </>
    )
}

export default TodoListPage
