import type { ReactElement } from "react";
import EventForm from "../components/EventForm";
import Navbar from "../components/Navbar";

const CreateEventPage = (): ReactElement => (
    <div>
        <Navbar />
        <EventForm />
    </div>
);

export default CreateEventPage;
