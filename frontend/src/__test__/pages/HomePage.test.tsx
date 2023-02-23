import { render } from "@testing-library/react";
import React from "react";

import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomePage from "../../pages/HomePage";

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
