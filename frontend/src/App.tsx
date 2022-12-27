import React from 'react';
import { LandingPage, Homepage } from './pages';
import { Route, Routes} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<LandingPage/>}/>
        <Route path="/home" element={<Homepage/>}/>
      </Routes>
    </div>
  );
}

export default App;
