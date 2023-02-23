import { render, screen } from "@testing-library/react";
import React from "react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import LandingPage from "../../pages/LandingPage";

const renderComponent = (): void => {
    render(
        <BrowserRouter>
            <Routes>
                <Route path="*" element={<LandingPage />} />
            </Routes>
        </BrowserRouter>
    );
};

describe("Landing Page", () => {
    it("renders LandingPage component title", () => {
        renderComponent();

        expect(screen.getByText("Schedulii")).toBeInTheDocument();
    });

    it("renders LandingPage component login button", () => {
        renderComponent();

        expect(screen.getByText("Login with Google")).toBeInTheDocument();
    });
});
