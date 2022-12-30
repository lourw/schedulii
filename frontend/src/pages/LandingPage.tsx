import type { ReactElement } from "react";
import React from "react";
import "../App.css";
import { useAuth } from "../context/AuthContext";
import type { AppContextInterface } from "../context/AuthContext";

const LandingPage = (): ReactElement => {
    const appContext: AppContextInterface | null = useAuth();

    return (
        <div className="wrapper">
            <div className="header" />
            <div className="title">
                <span>Schedulii</span>
                <button className="loginButton" onClick={appContext?.onLogin}>
                    Login with Google
                </button>
            </div>

            <p>Token: {appContext?.token}</p>

            <div className="footer" />
        </div>
    );
};

export default LandingPage;
