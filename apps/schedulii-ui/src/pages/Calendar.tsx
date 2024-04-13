function Calendar() {
  const weekdays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
  const months = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
  ];
  const currentDay = new Date();
  return (
    <>
      Calendar
      <h2>
        {' '}
        {months[currentDay.getMonth()]} {currentDay.getDay()},{' '}
        {currentDay.getFullYear()}
      </h2>
      <div className="calendar-body">
        <div className="table-header">
          {weekdays.map((weekday) => {
            return (
              <div className="weekday">
                <p>{weekday}</p>
              </div>
            );
          })}
        </div>
        <div className="table"></div>
      </div>
    </>
  );
}

export default function App() {
  return (
    <>
      <Calendar />
    </>
  );
}
