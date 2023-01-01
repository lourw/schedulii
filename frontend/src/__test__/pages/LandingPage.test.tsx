import React from "react";
import { render, screen } from "@testing-library/react";

import LandingPage from "../../pages/LandingPage";

describe("Landing Page", () => {
    it("renders LandingPage component title", () => {
        render(<LandingPage />);

        expect(screen.getByText("Schedulii")).toBeInTheDocument();
    });

    it("renders LandingPage component login button", () => {
        render(<LandingPage />);

        expect(screen.getByText("Login with Google")).toBeInTheDocument();
    });
});
