import type { ReactElement } from "react";
import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../context/AuthContext";
import type AuthProviderProps from "../types/AuthProviderProps";

const ProtectedRoute = ({ children }: AuthProviderProps): ReactElement => {
    const token = useAuth();
    const location = useLocation();

    if (!token) {
        return <Navigate to="/" replace state={{ from: location }} />;
    }

    return <> {children} </>;
};

export default ProtectedRoute;
