import type { ReactElement } from "react";
import { useEffect, useState } from "react";
import { AuthContext } from "./AuthContext";
import type AppContextDataType from "./AuthContextDataType";
import type AuthProviderPropsType from "./AuthProviderPropsType";

const AuthProvider = ({ children }: AuthProviderPropsType): ReactElement => {
    const [token, setToken] = useState(null);

    useEffect(() => {
        const localStorageToken = JSON.parse(
            localStorage.getItem("token") || "null"
        );

        if (localStorageToken) {
            setToken(localStorageToken);
        }
    }, []);

    useEffect(() => {
        localStorage.setItem("token", JSON.stringify(token));
    }, [token]);

    const handleLogin = async (): Promise<void> => {
        const request = await fetch("http://localhost:8080/login", {
            method: "POST",
            body: JSON.stringify({
                username: "schedulii-user"
            })
        });

        const data = await request.json();
        setToken(data.token);
    };

    const handleLogout = (): void => {
        setToken(null);
    };

    const value: AppContextDataType = {
        token,
        onLogin: handleLogin,
        onLogout: handleLogout
    };

    return (
        <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
    );
};

export default AuthProvider;
