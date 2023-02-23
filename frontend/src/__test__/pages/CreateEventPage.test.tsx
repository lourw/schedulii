import { render } from "@testing-library/react";
import React from "react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import CreateEventPage from "../../pages/CreateEventPage";

describe("Create Event Page", () => {
    it("renders CreateEventPage component title", () => {
        render(
            <BrowserRouter>
                <Routes>
                    <Route path="*" element={<CreateEventPage />} />
                </Routes>
            </BrowserRouter>
        );
    });
});
