import {createContext, useContext} from 'react'

export interface AppContextInterface {
    token: string | null,
    onLogin: () => {};
    onLogout: () => void;
}

export const AuthContext = createContext<AppContextInterface | null>(null);

export const useAuth = () => {
    return useContext(AuthContext);
}

