import React from 'react';
import css from './event-form.module.css';

export interface EventFormProps {
  formData: {
    eventName: string;
    startDate: string;
    endDate: string;
    earliestTime: string;
    latestTime: string;
    location: string;
  };
  isFormValid: boolean;
  onInputChange: (
    field: keyof EventFormProps['formData'],
    value: string
  ) => void;
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
}

export function EventForm({
  formData,
  isFormValid,
  onInputChange,
  onSubmit,
}: EventFormProps) {
  return (
    <div className={css.container}>
      <form className={css.eventForm} onSubmit={onSubmit}>
        <label>Event Name:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="text"
          placeholder="Event Name"
          value={formData.eventName}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('eventName', e.target.value)
          }
        />
        <label>Start Date:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="date"
          placeholder="Start Date"
          value={formData.startDate}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('startDate', e.target.value)
          }
        />

        <label>End Date:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="date"
          placeholder="End Date"
          value={formData.endDate}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('endDate', e.target.value)
          }
        />

        <label>No earlier than: </label>
        <input
          className="text-slate-500 rounded border px-1"
          type="time"
          placeholder="No earlier than:"
          value={formData.earliestTime}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('earliestTime', e.target.value)
          }
        />
        <label>No later than: </label>
        <input
          className="text-slate-500 rounded border px-1"
          type="time"
          placeholder="No later than:"
          value={formData.latestTime}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('latestTime', e.target.value)
          }
        />

        <label>Location</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="text"
          placeholder="Location"
          value={formData.location}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            onInputChange('location', e.target.value)
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
}

export default EventForm;
