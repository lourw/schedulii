// eslint-disable-next-line @typescript-eslint/no-unused-vars
import Landing from '../pages/Landing';
import { Route, Routes } from 'react-router-dom';
import Home from '../pages/Home';

export function App() {
  return (
    <>
      <div className="m-16">
        <h1 className="text-4xl mb-8">Schedulii</h1>
        <Routes>
          <Route path="/" element={<Landing />} />
          <Route path="/home" element={<Home />} />
        </Routes>
      </div>
    </>
  );
}

export default App;
