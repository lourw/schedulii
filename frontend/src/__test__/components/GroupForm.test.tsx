import { render, screen } from "@testing-library/react";
import React from "react";
import GroupForm from "../../components/GroupForm";

describe("GroupForm", () => {
    it("renders all input fields on GroupForm component", () => {
        render(<GroupForm />);

        const groupNameField = screen.getByLabelText("Group Name:");
        const groupUrlField = screen.getByLabelText("Group URL:");
        const availableStartHourField = screen.getByLabelText(
            "Available Start Hour:"
        );
        const availableEndHourField = screen.getByLabelText(
            "Available End Hour:"
        );

        expect(groupNameField).toBeInTheDocument();
        expect(groupUrlField).toBeInTheDocument();
        expect(availableStartHourField).toBeInTheDocument();
        expect(availableEndHourField).toBeInTheDocument();
    });
});
