import styles from './event-card.module.css';

export interface EventCardProps {
  title: string;
  startTime: string;
  endTime: string;
}

export function EventCard(props: EventCardProps) {
  const formatEventDateTime = (date: string) => {
    const dateObject = new Date(date);

    const options: Intl.DateTimeFormatOptions = {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
      hour12: true,
    };

    return dateObject.toLocaleString('en-US', options);
  };

  return (
    <div className={styles.container}>
      <p className={styles.title}>{props.title}</p>
      <p>
        {formatEventDateTime(props.startTime)} -{' '}
        {formatEventDateTime(props.endTime)}
      </p>
    </div>
  );
}

export default EventCard;
