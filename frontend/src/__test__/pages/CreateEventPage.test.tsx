import { render, screen } from "@testing-library/react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import CreateEventPage from "../../pages/CreateEventPage";

const renderComponent = (): void => {
    render(
        <BrowserRouter>
            <Routes>
                <Route path="*" element={<CreateEventPage />} />
            </Routes>
        </BrowserRouter>
    );
};

describe("Create Event Page", () => {
    it("renders CreateEventPage component title", () => {
        renderComponent();
    });

    it("renders an EventForm correctly", () => {
        renderComponent();

        const form = screen.findByRole("form");

        expect(form).toBeInTheDocument;
    });
});
