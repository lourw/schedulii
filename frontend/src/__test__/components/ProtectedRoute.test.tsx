/* eslint-disable @typescript-eslint/no-empty-function */
import React from "react";
import { render, screen } from "@testing-library/react";
import ProtectedRoute from "../../components/ProtectedRoute";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";
import type AuthContextDataType from "../../context/AuthContextDataType";

describe("Protected Route", () => {
    let validContext: AuthContextDataType;

    beforeEach(() => {
        validContext = {
            token: "invalid-token",
            onLogin: (): void => { },
            onLogout: (): void => { },
        };
    });

    it("should render default route when no user is validated", async () => {
        validContext.token = null;

        render(
            <MemoryRouter initialEntries={["/protected"]}>
                <Routes>
                    <Route path="/" element={
                        <p>Home</p>
                    } />
                    <Route path="/protected" element={
                        <AuthContext.Provider value={validContext}>
                            <ProtectedRoute><p>Child</p></ProtectedRoute>
                        </AuthContext.Provider>
                    } />
                </Routes>
            </MemoryRouter>
        );

        expect(screen.getByText("Home")).toBeInTheDocument();
    });

    it("should render children if user is validated", () => {
        // localStorage.setItem("token", JSON.stringify("test-token"));
        render(
            <MemoryRouter initialEntries={["/protected"]}>
                <Routes>
                    <Route path="/" element={
                        <p>Home</p>
                    } />
                    <Route path="/protected" element={
                        <AuthContext.Provider value={validContext}>
                            <ProtectedRoute><p>Child</p></ProtectedRoute>
                        </AuthContext.Provider>
                    } />
                </Routes>
            </MemoryRouter>
        );
        expect(screen.getByText("Child")).toBeInTheDocument();
    });
});
