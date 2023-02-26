import type { ReactElement } from "react";
import React from "react";
import { Route, Routes } from "react-router-dom";
import Calendar from "./components/Calendar";
import ProtectedRoute from "./components/ProtectedRoute";
import AuthProvider from "./context/AuthProvider";
import CreateEventPage from "./pages/CreateEventPage";
import CreateGroupPage from "./pages/CreateGroupPage";
import HomePage from "./pages/HomePage";
import LandingPage from "./pages/LandingPage";

const App = (): ReactElement => (
    <div className="App">
        <AuthProvider>
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
                    path="/create-group"
                    element={
                        <ProtectedRoute>
                            <CreateGroupPage />
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
        </AuthProvider>
    </div>
);

export default App;
