import React, { useEffect, useState } from 'react';
import { EventCard } from '@schedulii/schedulii-components';

interface HomeProps {}

interface Event {
  event_id: number;
  event_name: string;
  start_time: string;
  end_time: string;
}

const API_URL = import.meta.env.VITE_API_URL;

const Home: React.FC<HomeProps> = (props) => {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchEvents = async () => {
      setLoading(true);
      try {
        const response = await fetch(`${API_URL}/events`);
        const data = await response.json();

        setEvents(data);

        /* eslint-disable @typescript-eslint/no-explicit-any */
      } catch (err: any) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchEvents();
  }, []);

  if (loading) return <div>Loading events...</div>;
  if (error) return <div>Error loading events</div>;

  return (
    <div className="events">
      {events.map((event) => (
        <EventCard
          key={event.event_id}
          title={event.event_name}
          startTime={event.start_time}
          endTime={event.end_time}
        />
      ))}
    </div>
  );
};

export default Home;
