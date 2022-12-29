import React from 'react';
import Navbar from '../components/Navbar';
import '../pages/Homepage.css';

const Homepage = () => {
  return (
    <>
    <Navbar />
      <section className="container">
        <div id="left">
          <h1>Your Calendar</h1>
        </div>
        <div id="right">
          <h1>Your Events</h1>
        </div>
      </section>
    </>
  )
}

export default Homepage
