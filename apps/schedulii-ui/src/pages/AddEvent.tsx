import React, { useState } from 'react';
import css from './AddEvent.module.css';

interface AddEventProps {}

const AddEvent: React.FC<AddEventProps> = () => {
  const [formData, setFormData] = useState({
    eventName: '',
    startDate: '',
    endDate: '',
    earliestTime: '',
    latestTime: '',
    location: '',
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(formData);
  };

  return (
    <div className={css.container}>
      <form className={css.eventForm} onSubmit={handleSubmit}>
        <label>Event Name:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="text"
          placeholder="Event Name"
          value={formData.eventName}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, eventName: e.target.value })
          }
        />
        <label>Start Date:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="date"
          placeholder="Start Date"
          value={formData.startDate}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, startDate: e.target.value })
          }
        />

        <label>End Date:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="date"
          placeholder="End Date"
          value={formData.endDate}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, endDate: e.target.value })
          }
        />

        <label>No earlier than: </label>
        <input
          className="text-slate-500 rounded border px-1"
          type="time"
          placeholder="No earlier than:"
          value={formData.earliestTime}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, earliestTime: e.target.value })
          }
        />
        <label>No later than: </label>
        <input
          className="text-slate-500 rounded border px-1"
          type="time"
          placeholder="No later than:"
          value={formData.latestTime}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, latestTime: e.target.value })
          }
        />

        <label>Location</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="text"
          placeholder="Location"
          value={formData.location}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setFormData({ ...formData, location: e.target.value })
          }
        />

        <button
          className="bg-sky-400 font-medium text-slate-50 px-10 py-2 self-center rounded-full"
          type="submit"
        >
          Submit
        </button>
      </form>
    </div>
  );
};

export default AddEvent;
