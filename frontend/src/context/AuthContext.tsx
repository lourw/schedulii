import { createContext, useContext } from "react";

export interface AppContextInterface {
    token: string | null;
    onLogin: () => void;
    onLogout: () => void;
}

export const AuthContext = createContext<AppContextInterface | null>(null);

export const useAuth = (): AppContextInterface | null =>
    useContext(AuthContext);
