import type { ReactElement } from "react";
import { useEffect } from "react";
import React from "react";
import { startOfDay, addMinutes, addHours, addDays } from "date-fns";
import formatDate from "date-fns/format";
import styled from "styled-components";
import "./Calendar.css";

type CalendarPropsType = {
    numDays: number;
    minTime: number;
    maxTime: number;
    startDate: Date;
    columns: number;
    rows: number;
};

type CalendarDateType = Date[][];

const Grid = styled.div<{ columns: number; rows: number }>`
    display: grid;
    grid-template-columns: auto repeat(${(props): number => props.columns}, 1fr);
    grid-template-rows: auto repeat(${(props): number => props.rows}, 1fr);
    width: 100%;
`;

const Calendar = (props: CalendarPropsType): ReactElement => {
    const [dates, setDates] = React.useState<CalendarDateType>([[new Date(0)]]);
    const cellToDate: Map<Element, Date> = new Map();

    useEffect((): void => {
        setDates(computeDatesMatrix(props));
    }, []);

    const computeDatesMatrix = (props: CalendarPropsType): CalendarDateType => {
        const startTime = startOfDay(props.startDate);
        const dates: Date[][] = [];
        for (let d = 0; d < props.numDays; d++) {
            // create all the chunks for each day in currentDay
            const currentDay = [];
            for (let h = props.minTime; h < props.maxTime; h++) {
                // later, add logic for chunking it into fractions of hours
                currentDay.push(
                    addMinutes(addHours(addDays(startTime, d), h), 0)
                );
            }
            dates.push(currentDay);
        }
        return dates;
    };

    const renderDateCellWrapper = (time: Date): ReactElement => (
        <div
            className="grid-wrapper"
            role="presentation"
            key={time.toISOString()}
            // mouse handlers later here
        >
            {renderDateCell(time)}
        </div>
    );

    const renderDateCell = (time: Date): ReactElement => {
        const refSetter = (dateCell: HTMLElement | null): void => {
            if (dateCell) {
                cellToDate.set(dateCell, time);
            }
        };
        return (
            <div
                className="date-cell"
                // selected={selected}
                ref={refSetter}
            />
        );
    };

    const renderTimeLabel = (time: Date): ReactElement => (
        <span className="timeText">{formatDate(time, "ha")}</span>
    );

    const renderDateLabel = (date: Date): ReactElement => (
        <span className="dateLabel">{formatDate(date, "M/d")}</span>
    );

    const renderFullDateGrid = (): ReactElement[] => {
        // eslint-disable-line
        const flattenedDates = [];

        const numDays = dates.length;
        const numTimes = dates[0].length;

        for (let j = 0; j < numTimes; j++) {
            for (let i = 0; i < numDays; i++) {
                // turns 2d array into 1 for easier operations
                flattenedDates.push(dates[i][j]);
            }
        }
        const dateGridElements = flattenedDates.map(renderDateCellWrapper);
        for (let i = 0; i < numTimes; i++) {
            const index = i * numDays; // index of flattenedArray
            const time = dates[0][i];
            // insert a time label at the start of every row
            dateGridElements.splice(index + i, 0, renderTimeLabel(time));
        }
        return [
            <div key="topleft" />,
            // top row with the dates
            ...dates.map((dayOfTimes, index) =>
                React.cloneElement(renderDateLabel(dayOfTimes[0]), {
                    key: `date-${index}`
                })
            ),
            // every row below that
            ...dateGridElements.map((element, index) =>
                React.cloneElement(element, { key: `time-${index}` })
            )
        ];
    };

    return (
        <div className="wrapper">
            <Grid columns={dates.length} rows={dates[0].length}>
                {renderFullDateGrid()}
            </Grid>
        </div>
    );
};

export default Calendar;
