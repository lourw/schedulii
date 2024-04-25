// eslint-disable-next-line @typescript-eslint/no-unused-vars
import Landing from '../pages/Landing';
import { Route, Routes } from 'react-router-dom';
import Home from '../pages/Home';
import AddEvent from '../pages/AddEvent';
import Calendar from '../pages/Calendar';

export function App() {
  return (
    <div className="m-16">
      <h1 className="text-4xl mb-8">Schedulii</h1>
      <Routes>
        <Route path="/" element={<Landing />} />
        <Route path="/home" element={<Home />} />
        <Route path="/add-event" element={<AddEvent />} />
        <Route path="/calendar" element={<Calendar />} />
      </Routes>
    </div>
  );
}

export default App;
