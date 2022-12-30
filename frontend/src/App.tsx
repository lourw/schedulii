import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Homepage from './pages/Homepage';
import CreateEvent from './pages/CreateEvent';
import ProtectedRoute from './components/ProtectedRoute';
import AuthProvider from './context/AuthProvider';

const App = () => {
  return (
    <div className="App">
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route index element={<LandingPage/>} />
            <Route path="/home" element={
              <ProtectedRoute>
                <Homepage />
              </ProtectedRoute>}/>
            <Route path="/create-event" element={
              <ProtectedRoute>
                <CreateEvent />
              </ProtectedRoute>}/>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </div>
  );
}

export default App;
