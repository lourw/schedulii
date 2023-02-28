import type { ChangeEvent, ReactElement } from "react";
import React from "react";
import "./GroupForm.css";

type FormData = {
    groupId: string;
    groupName: string;
    groupUrl: string;
    availableStartHour: string;
    availableEndHour: string;
};

const GroupForm = (): ReactElement => {
    const [formData, setFormData] = React.useState<FormData>({
        groupId: "",
        groupName: "",
        groupUrl: "",
        availableStartHour: "",
        availableEndHour: ""
    });

    //HandleSumbit
    const handleChange = (event: ChangeEvent<HTMLInputElement>): void =>
        setFormData((prevFormData) => ({
            ...prevFormData,
            [event.target.name]: event.target.value
        }));

    return (
        <form>
            <label htmlFor="group-name">Group Name: </label>
            <input
                type="text"
                id="group-name"
                name="groupName"
                value={formData.groupName}
                onChange={handleChange}
            />

            <label htmlFor="group-url">Group URL: </label>
            <input
                type="text"
                id="group-url"
                name="groupUrl"
                value={formData.groupUrl}
                onChange={handleChange}
            />

            <label htmlFor="available-start-hour">Available Start Hour:</label>
            <input
                type="number"
                id="available-start-hour"
                name="availableStartHour"
                value={formData.availableStartHour}
                onChange={handleChange}
            />

            <label htmlFor="available-end-hour">Available End Hour:</label>
            <input
                type="number"
                id="available-end-hour"
                name="availableEndHour"
                value={formData.availableEndHour}
                onChange={handleChange}
            />
            <button type="submit">Create Group</button>
        </form>
    );
};

export default GroupForm;
