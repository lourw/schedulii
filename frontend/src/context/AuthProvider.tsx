import React from 'react'
import Props from "../types/Props"
import { AppContextInterface, AuthContext } from './AuthContext';

const AuthProvider = ({ children }: Props) => {
    const [token, setToken] = React.useState(null);

    const handleLogin = async () => {
        const request = await fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify({
                "username": "schedulii-user"
            }),
        });

        const data = await request.json();
        setToken(data.token);
    };

    const handleLogout = () => {
        setToken(null);
    };

    const value: AppContextInterface = {
        token, 
        onLogin: handleLogin,
        onLogout: handleLogout
    };

    return (
        <AuthContext.Provider value={value}>
            { children }
        </AuthContext.Provider>
    )
}


export default AuthProvider
