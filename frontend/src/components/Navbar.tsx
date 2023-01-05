import type { ReactElement } from "react";
import React from "react";
import { NavLink } from "react-router-dom";
import "../components/Navbar.css";
import { useAuth } from "../context/AuthContext";

const Navbar = (): ReactElement => {
    const appContext = useAuth();

    return (
        <nav className="navbar">
            <div className="nav-elements">
                <ul>
                    <li>
                        <NavLink to="/home">Home</NavLink>
                    </li>
                    <li>
                        <NavLink to="/create-event"> Create New Event</NavLink>{" "}
                        {/* this is a placeholder for events page */}
                    </li>
                    <li>
                        <button onClick={appContext?.onLogout}>Logout</button>
                    </li>
                </ul>
            </div>
        </nav>
    );
};

export default Navbar;
