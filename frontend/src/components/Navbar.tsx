import React from 'react'
import { NavLink } from 'react-router-dom'
import '../components/Navbar.css'

const Navbar = () => {
  return (
    <nav className="navbar"> 
        <div className="nav-elements">
            <ul>
                <li>
                    <NavLink to="/home">Home</NavLink>
                </li>
                <li>
                    <NavLink to="/create-event"> Create New Event</NavLink> 
                </li>
            </ul>
        </div>
    </nav>
  )
}

export default Navbar