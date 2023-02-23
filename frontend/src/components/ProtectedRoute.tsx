import type { ReactElement } from "react";
import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import type AuthProviderPropsType from "../context/AuthProviderPropsType";

const ProtectedRoute = ({ children }: AuthProviderPropsType): ReactElement => {
    const appContext = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        if (!appContext?.token) {
            navigate("/");
        }
    }, [appContext?.token]);

    return <> {children} </>;
};

export default ProtectedRoute;
