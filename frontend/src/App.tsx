import type { ReactElement } from "react";
import React from "react";
import { Routes, Route } from "react-router-dom";
import Calendar from "./components/Calendar";
import ProtectedRoute from "./components/ProtectedRoute";
import CreateEventPage from "./pages/CreateEventPage";
import HomePage from "./pages/HomePage";
import LandingPage from "./pages/LandingPage";

const App = (): ReactElement => (
    <div className="App">
        <Routes>
            <Route index element={<LandingPage />} />
            <Route
                path="/home"
                element={
                    <ProtectedRoute>
                        <HomePage />
                    </ProtectedRoute>
                }
            />
            <Route
                path="/create-event"
                element={
                    <ProtectedRoute>
                        <CreateEventPage />
                    </ProtectedRoute>
                }
            />
            <Route
                path="/availability"
                element={
                    <ProtectedRoute>
                        <Calendar
                            numDays={6}
                            minTime={9}
                            maxTime={17}
                            startDate={new Date()}
                            columns={6}
                            rows={6}
                        />
                    </ProtectedRoute>
                }
            />
        </Routes>
    </div>
);

export default App;
