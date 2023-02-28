import type { ChangeEvent, ReactElement } from "react";
import React from "react";
import "./EventForm.css";

type FormData = {
    eventId: number;
    groupId: number;
    eventName: string;
    startTime: string;
    endTime: string;
};

const EventForm = (): ReactElement => {
    const [formData, setFormData] = React.useState<FormData>({
        eventId: 0,
        groupId: 0,
        eventName: "",
        startTime: "",
        endTime: ""
    });

    const handleChange = (event: ChangeEvent<HTMLInputElement>): void => {
        const { name, value } = event.target;

        if (name === "startTime" || name === "endTime") {
            setFormData((prevFormData) => ({
                ...prevFormData,
                [name]: new Date(value).toISOString()
            }));
        } else {
            setFormData((prevFormData) => ({
                ...prevFormData,
                [name]: value
            }));
        }
    };

    return (
        <form>
            <label htmlFor="event-id">Event ID: </label>
            <input
                type="number"
                id="event-id"
                name="eventId"
                value={formData.eventId}
                onChange={handleChange}
            />

            <label htmlFor="group-id">Group ID: </label>
            <input
                type="number"
                id="group-id"
                name="groupId"
                value={formData.groupId}
                onChange={handleChange}
            />

            <label htmlFor="event-name">Event Name: </label>
            <input
                type="text"
                id="event-name"
                name="eventName"
                value={formData.eventName}
                onChange={handleChange}
            />

            <label htmlFor="start-time">Start Time:</label>
            <input
                type="datetime-local"
                id="start-time"
                name="startTime"
                value={formData.startTime}
                onChange={handleChange}
            />

            <label htmlFor="end-time">End Time:</label>
            <input
                type="datetime-local"
                id="end-time"
                name="endTime"
                value={formData.endTime}
                onChange={handleChange}
            />
            <button type="submit">Create Event</button>
        </form>
    );
};

export default EventForm;
