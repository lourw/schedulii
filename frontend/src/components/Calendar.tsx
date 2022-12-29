import React from 'react'
import './Navbar.css'

type PropsType = {
  numDays: number;
  startTime: number;
  endTime: number;
}

type StateType = {
  dates: Array<Array<Date>>
}

export default class Calendar extends React.Component<PropsType, StateType> {
  cellToDate: Map<Element, Date> = new Map();
  renderDateCellWrapper = (time: Date): JSX.Element => {
    // later, attach event handlers for clicking here

    return (
      <div className="grid-wrapper"
        role="presentation"
        key={time.toISOString()}
      // mouse handlers later here
      >
        {this.renderDateCell(time)}
      </div>
      // signature of renderDateCell will need to have selected later
    )
  }

  renderDateCell = (time: Date): JSX.Element => {
    const refSetter = (dateCell: HTMLElement | null) => {
      if (dateCell) {
        this.cellToDate.set(dateCell, time);
      }
    }
    return (
      <div className="date-cell"
        // selected={selected}
        ref={refSetter}> a </div>
    )
  }


  renderFullDateGrid(): Array<JSX.Element> {
    const flattenedDates = [];
    const numDays = 3;
    const numTimes = 6;
    for (let j = 0; j < numTimes; j++) {
      for (let i = 0; i < numDays; i++) {
        // turns 2d array into 1 for easier operations
        flattenedDates.push(this.state.dates[i][j]);
      }
    }
    const dateGridElements = flattenedDates.map(this.renderDateCellWrapper);
    for (let i = 0; i < numTimes; i++) {
      // const index = i * numDays // index of flattenedArray
      // const time = this.state.dates[0][i]
      // insert a time label at the start of every row
      // dateGridElements.splice(index + i, 0, this.renderTimeLabel(time))
    }
    return [
      <div key="topleft" />,
      ...dateGridElements.map((element, index) => React.cloneElement(element, { key: `time-${index}`}))
    ]
  }

  render() {
    return (
      <div className='wrapper'>
        <div className='grid'>
          {this.renderFullDateGrid()}
        </div>
      </div>
    )
  }
}
