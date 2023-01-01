import { render } from "@testing-library/react";
import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "../App";

test("render react application", () => {
    render(
        <BrowserRouter>
            <Routes>
                <Route path="*" element={<App />} />
            </Routes>
        </BrowserRouter>
    );
});
