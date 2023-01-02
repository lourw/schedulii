import React from "react";
import { render } from "@testing-library/react";

import CreateEventPage from "../../pages/CreateEventPage";
import { BrowserRouter, Routes, Route } from "react-router-dom";

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
