import { render, screen } from "@testing-library/react";
import EventForm from "../../components/EventForm";

describe("EventForm", () => {
    it("renders all input fields on EventForm component", () => {
        render(<EventForm />);

        const eventIdField = screen.getByLabelText("Event ID:");
        const groupIdField = screen.getByLabelText("Group ID:");
        const eventNameField = screen.getByLabelText("Event Name:");
        const startTimeField = screen.getByLabelText("Start Time:");
        const endTimeField = screen.getByLabelText("End Time:");

        expect(eventIdField).toBeInTheDocument();
        expect(groupIdField).toBeInTheDocument();
        expect(eventNameField).toBeInTheDocument();
        expect(startTimeField).toBeInTheDocument();
        expect(endTimeField).toBeInTheDocument();
    });
});
