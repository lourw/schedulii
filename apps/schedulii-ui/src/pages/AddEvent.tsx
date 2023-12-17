import css from './AddEvent.module.css';

interface AddEventProps {

}

const AddEvent: React.FC<AddEventProps> = () => {

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log('Form submitted');
  }
  return (
    <div className={css.container}>
      <form className={css.eventForm} onSubmit={handleSubmit}>
        <label>Event Name</label>
        <input type="text" placeholder="Event Name" />
        <label>Start Date</label>
        <input type="date" placeholder="Start Date" />

        <label>End Date</label>
        <input type="date" placeholder="End Date" />

        <label>No earlier than: </label>
        <input type="time" placeholder="No earlier than:" />
        <label>No later than: </label>
        <input type="time" placeholder="No later than:" />

        <label>Location</label>
        <input type="text" placeholder="Location" />

        <button type="submit">Submit</button>

      </form>
    </div>
  );
}

export default AddEvent;
