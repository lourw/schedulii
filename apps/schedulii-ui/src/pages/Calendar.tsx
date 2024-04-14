import {
  startOfMonth,
  endOfMonth,
  eachDayOfInterval,
  format,
  subDays,
  addDays,
  subMonths,
  addMonths,
  isSameDay,
} from 'date-fns';

import { useState } from 'react';

const Calendar = () => {
  // Get details about the current month
  const [currentDate, setCurrentDate] = useState(new Date());
  const daysOfWeek = [
    'Sunday',
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday',
  ];
  const firstDayOfMonth = startOfMonth(currentDate);
  const lastDayOfMonth = endOfMonth(currentDate);
  const diffStart = daysOfWeek.indexOf(format(firstDayOfMonth, 'iiii')); // returns how many days before current month to fill in grid
  const diffEnd = 6 - daysOfWeek.indexOf(format(lastDayOfMonth, 'iiii')); // returns how many days after current month to fill in grid

  // Generate array of dates for the current month
  const datesInMonth = eachDayOfInterval({
    start: firstDayOfMonth,
    end: lastDayOfMonth,
  });

  const datesInLastMonth = eachDayOfInterval({
    start: subDays(startOfMonth(currentDate), diffStart),
    end: subDays(startOfMonth(currentDate), 1),
  });

  const datesInNextMonth = eachDayOfInterval({
    start: addDays(endOfMonth(currentDate), 1),
    end: addDays(endOfMonth(currentDate), diffEnd),
  });

  // Classname var
  const otherMonths = 'text-xl text-slate-400';
  const buttons = 'mx-5 p-2 border-4';

  // Render calendar
  return (
    <>
      <div>
        <h1 className="text-3xl">Calendar</h1>
        <div className="flex">
          <button
            className={buttons}
            onClick={() => setCurrentDate(subMonths(currentDate, 1))}
          >
            {' '}
            Left Button{' '}
          </button>
          <h1 className="text-3xl"> {format(currentDate, 'MMMM yyyy')} </h1>
          <button
            className={buttons}
            onClick={() => setCurrentDate(addMonths(currentDate, 1))}
          >
            {' '}
            Right Button{' '}
          </button>
        </div>
      </div>
      <div className="calendar">
        <div className="grid grid-cols-7 bg-fuchsia-50">
          {daysOfWeek.map((day) => (
            <div className="text-2xl">{day}</div>
          ))}

          {datesInLastMonth.map((date) => (
            <div className={otherMonths}> {format(date, 'd')} </div>
          ))}

          {datesInMonth.map((date) => (
            <div
              className={
                'text-xl ' + (isSameDay(new Date(), date) ? 'font-bold' : '')
              }
            >
              {' '}
              {format(date, 'd')}{' '}
            </div>
          ))}

          {datesInNextMonth.map((date) => (
            <div className={otherMonths}> {format(date, 'd')} </div>
          ))}
        </div>
      </div>
    </>
  );
};

export default Calendar;
