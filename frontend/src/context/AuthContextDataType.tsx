type AppContextDataType = {
    token: string | null;
    onLogin: () => void;
    onLogout: () => void;
};

export default AppContextDataType;
