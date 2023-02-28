import { render, screen } from "@testing-library/react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import CreateGroupPage from "../../pages/CreateGroupPage";

const renderComponent = (): void => {
    render(
        <BrowserRouter>
            <Routes>
                <Route path="*" element={<CreateGroupPage />} />
            </Routes>
        </BrowserRouter>
    );
};

describe("Create Group Page", () => {
    it("renders CreateEventPage component title", () => {
        renderComponent();
    });

    it("renders a GroupForm correctly", () => {
        renderComponent();

        const form = screen.findByRole("form");

        expect(form).toBeInTheDocument;
    });
});
