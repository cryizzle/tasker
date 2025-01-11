import type React from 'react'
import UserLogin from '../components/UserLogin'

const LoginPage: React.FC = () => {
    return (
        <>
            <h1 style={{ fontSize: '4em' }}>Tasker</h1>
            <UserLogin />
        </>
    )
}

export default LoginPage
