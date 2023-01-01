import React from "react";
import { render } from "@testing-library/react";

import HomePage from "../../pages/HomePage";
import { BrowserRouter, Route, Routes } from "react-router-dom";

describe("Home Page", () => {
    it("renders HomePage component", () => {
        render(
            <BrowserRouter>
                <Routes>
                    <Route path="*" element={<HomePage />} />
                </Routes>
            </BrowserRouter>
        );
    });
});
