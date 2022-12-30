import type { ReactElement } from "react";
import React from "react";
import "./Calendar.css";
import { startOfDay, addMinutes, addHours, addDays } from "date-fns";
import formatDate from "date-fns/format";
import styled from "styled-components";

type PropsType = {
    numDays: number;
    minTime: number;
    maxTime: number;
    startDate: Date;
    columns: number;
    rows: number;
};

type StateType = {
    dates: Array<Array<Date>>;
};

const Grid = styled.div<{ columns: number; rows: number }>`
    display: grid;
    grid-template-columns: auto repeat(${(props): number => props.columns}, 1fr);
    grid-template-rows: auto repeat(${(props): number => props.rows}, 1fr);
    width: 100%;
`;

export default class Calendar extends React.Component<PropsType, StateType> {
    cellToDate: Map<Element, Date> = new Map();

    static getStateFromProps(props: PropsType): Partial<StateType> | null {
        return {
            dates: Calendar.computeDatesMatrix(props)
        };
    }

    static computeDatesMatrix(props: PropsType): Array<Array<Date>> {
        const startTime = startOfDay(props.startDate);
        const dates: Array<Array<Date>> = [];
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
    }

    constructor(props: PropsType) {
        super(props);

        this.state = {
            dates: Calendar.computeDatesMatrix(props)
        };
    }

    renderDateCellWrapper = (time: Date): JSX.Element => (
        // eslint-disable-line
        // later, attach event handlers for clicking here
        <div
            className="grid-wrapper"
            role="presentation"
            key={time.toISOString()}
            // mouse handlers later here
        >
            {this.renderDateCell(time)}
        </div>
        // signature of renderDateCell will need to have selected later
    );

    renderDateCell = (time: Date): JSX.Element => {
        // eslint-disable-line
        const refSetter = (dateCell: HTMLElement | null): void => {
            if (dateCell) {
                this.cellToDate.set(dateCell, time);
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

    renderTimeLabel = (time: Date): JSX.Element => (
        // eslint-disable-line
        <span className="timeText">{formatDate(time, "ha")}</span>
    );

    renderDateLabel = (date: Date): JSX.Element => (
        // eslint-disable-line
        <span className="dateLabel">{formatDate(date, "M/d")}</span>
    );

    renderFullDateGrid(): Array<JSX.Element> {
        // eslint-disable-line
        const flattenedDates = [];
        const numDays = this.state.dates.length;
        const numTimes = this.state.dates[0].length;
        for (let j = 0; j < numTimes; j++) {
            for (let i = 0; i < numDays; i++) {
                // turns 2d array into 1 for easier operations
                flattenedDates.push(this.state.dates[i][j]);
            }
        }
        const dateGridElements = flattenedDates.map(this.renderDateCellWrapper);
        for (let i = 0; i < numTimes; i++) {
            const index = i * numDays; // index of flattenedArray
            const time = this.state.dates[0][i];
            // insert a time label at the start of every row
            dateGridElements.splice(index + i, 0, this.renderTimeLabel(time));
        }
        return [
            <div key="topleft" />,
            // top row with the dates
            ...this.state.dates.map((dayOfTimes, index) =>
                React.cloneElement(this.renderDateLabel(dayOfTimes[0]), {
                    key: `date-${index}`
                })
            ),
            // every row below that
            ...dateGridElements.map((element, index) =>
                React.cloneElement(element, { key: `time-${index}` })
            )
        ];
    }

    render(): ReactElement {
        return (
            <div className="wrapper">
                <Grid
                    columns={this.state.dates.length}
                    rows={this.state.dates[0].length}
                >
                    {this.renderFullDateGrid()}
                </Grid>
            </div>
        );
    }
}
