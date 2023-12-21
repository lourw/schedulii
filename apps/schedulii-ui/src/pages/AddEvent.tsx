import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import css from './AddEvent.module.css';

interface AddEventProps {}

interface FormData {
  eventName: string;
  startDate: string;
  endDate: string;
  earliestTime: string;
  latestTime: string;
  location: string;
}

const AddEvent: React.FC<AddEventProps> = () => {
  const API_URL = import.meta.env.VITE_API_URL;
  const [formData, setFormData] = useState<FormData>({
    eventName: '',
    startDate: '',
    endDate: '',
    earliestTime: '',
    latestTime: '',
    location: '',
  });
  const [isFormValid, setIsFormValid] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    setIsFormValid(validateForm());
  }, [formData]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (isFormValid) {
      axios
        .post(`${API_URL}/events/add`, {
          event_name: formData.eventName,
          start_time: new Date(formData.startDate),
          end_time: new Date(formData.endDate),
          location: formData.location,
        })
        .then((res) => {
          navigate('/home');
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  const validateForm = () => {
    let isValid = true;
    const isEmptyField = Object.values(formData).some(
      (field) => field.trim().length === 0
    );

    const invalidDateRange = formData.endDate < formData.startDate;

    if (isEmptyField || invalidDateRange) {
      isValid = false;
    }
    return isValid;
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

        {isFormValid ? (
          <button
            className="bg-sky-400 font-medium text-slate-50 px-10 py-2 self-center rounded-full"
            type="submit"
          >
            Submit
          </button>
        ) : (
          <button
            className="bg-slate-400 font-medium text-slate-50 px-10 py-2 self-center rounded-full"
            type="submit"
            disabled
          >
            {' '}
            Submit{' '}
          </button>
        )}
      </form>
    </div>
  );
};

export default AddEvent;
