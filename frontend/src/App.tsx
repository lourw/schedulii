import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Homepage from './pages/Homepage';
import CreateEvent from './pages/CreateEvent';
import ProtectedRoute from './components/ProtectedRoute';
import AuthProvider from './context/AuthProvider';
import Calendar from './components/Calendar';

const App = () => {
  return (
    <div className="App">
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route index element={<LandingPage />} />
            <Route path="/home" element={
              <ProtectedRoute>
                <Homepage />
              </ProtectedRoute>} />
            <Route path="/create-event" element={
              <ProtectedRoute>
                <CreateEvent />
              </ProtectedRoute>} />
            <Route path="/availability" element={
              <ProtectedRoute>
                <Calendar
                  numDays={6}
                  minTime={9}
                  maxTime={17}
                  startDate={new Date()}
                  columns={6}
                  rows={6} />
              </ProtectedRoute>} />
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </div>
  );
}

export default App;
