import type { ReactElement } from "react";
import React from "react";
import "../components/EventCard.css";

type EventCardProps = {
    title: string;
    description: string;
}

const EventCard = ({ title, description }: EventCardProps): ReactElement => (
    <section className="card-container">
        <div id="card-title">{title}</div>
        <div id="card-description">{description}</div>
    </section>
);

export default EventCard;
