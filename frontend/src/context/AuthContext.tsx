import { createContext, useContext } from "react";
import type AppContextDataType from "./AuthContextDataType";

export const AuthContext = createContext<AppContextDataType | null>(null);

export const useAuth = (): AppContextDataType | null => useContext(AuthContext);
