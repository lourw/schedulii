import React from "react";
import { render, screen } from "@testing-library/react";

import LandingPage from "../../pages/LandingPage";
import { BrowserRouter, Routes, Route } from "react-router-dom";

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
