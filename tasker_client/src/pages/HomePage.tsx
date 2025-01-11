import type React from 'react'
import { useAppSelector } from '../app/hooks'
import { selectActiveUser } from '../app/userSlice'
import AllLists from '../components/AllLists'
import Navbar from '../components/common/Navbar'

const HomePage: React.FC = () => {
    const activeUser = useAppSelector(selectActiveUser)
    return (
        <>
            <Navbar />
            <h1>Welcome {activeUser.email}</h1>
            <AllLists />
        </>
    )
}

export default HomePage
