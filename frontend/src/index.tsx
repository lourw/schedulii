import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Routes, Route} from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Homepage from './pages/Homepage';
import CreateEvent from './pages/CreateEvent';
import Calendar from './components/Calendar';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
      <Routes>
        <Route path="/" element={<LandingPage/>}/>
        <Route path="/home" element={<Homepage/>}/>
        <Route path="/create-event" element={<CreateEvent/>}/>
        <Route path="/availability" element={<Calendar 
                                                numDays={3}
                                                startTime={9}
                                                endTime={15} />}/>
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
