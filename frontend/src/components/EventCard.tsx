import React from 'react'
import '../components/EventCard.css'

interface EventCardProps {
  title: string,
  description: string;
}

function EventCard({title, description}: EventCardProps)  {
  return (
    <section className="card-container">
      <div id="card-title">
        {title}
      </div>
      <div id="card-description">
        {description}
      </div>
    </section>
  )
}

export default EventCard
