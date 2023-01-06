import React from "react";
import { render } from "@testing-library/react";

import HomePage from "../../pages/HomePage";
import { BrowserRouter, Routes, Route } from "react-router-dom";

describe("Home Page", () => {
    it("renders HomePage component", () => {
        render(
            <BrowserRouter>
                <Routes>
                    <Route index element={<HomePage />} />
                </Routes>
            </BrowserRouter>
        );
    });
});
