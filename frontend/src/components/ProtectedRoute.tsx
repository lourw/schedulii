import React from 'react'
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from '../context/AuthContext';
import Props from '../types/Props';

const ProtectedRoute = ({ children }: Props) => {
    const token = useAuth();
    const location = useLocation();

    if (!token) {
        return <Navigate to="/" replace state={{ from: location }}/>
    }    

    return <> {children} </>;
}

export default ProtectedRoute
