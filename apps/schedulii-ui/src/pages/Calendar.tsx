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
  const generalDays = 'text-center p-2 pb-4';
  const otherMonths = 'text-xl text-slate-400';
  const buttons = 'text-3xl flex-auto';

  // Render calendar
  return (
    <>
      <div>
        <div className="flex mb-5">
          <button
            className={buttons}
            onClick={() => setCurrentDate(subMonths(currentDate, 1))}
          >
            {' '}
            ◀️{' '}
          </button>
          <h1 className="text-3xl"> {format(currentDate, 'MMMM yyyy')} </h1>
          <button
            className={buttons}
            onClick={() => setCurrentDate(addMonths(currentDate, 1))}
          >
            {' '}
            ▶️{' '}
          </button>
        </div>
      </div>
      <div>
        <div className="grid grid-cols-7 bg-fuchsia-50">
          {daysOfWeek.map((day) => (
            <div
              className={
                'text-2xl mb-2 pb-2 border-b border-black ' + generalDays
              }
            >
              {day[0]}
            </div>
          ))}
          {datesInLastMonth.map((date) => (
            <div className={otherMonths + ' ' + generalDays}>
              {' '}
              {format(date, 'd')}{' '}
            </div>
          ))}

          {datesInMonth.map((date) => (
            <div
              className={
                'text-xl ' +
                (isSameDay(new Date(), date) ? 'font-bold text-red-800' : '') +
                ' ' +
                generalDays
              }
            >
              {' '}
              {format(date, 'd')}{' '}
            </div>
          ))}

          {datesInNextMonth.map((date) => (
            <div className={otherMonths + ' ' + generalDays}>
              {' '}
              {format(date, 'd')}{' '}
            </div>
          ))}
        </div>
      </div>
    </>
  );
};

export default Calendar;
