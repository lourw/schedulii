import { render } from '@testing-library/react';
import EventCard from './event-card';

describe('EventCard', () => {
  it('should render successfully', () => {
    const props = {
      title: "Sample Event",
      startTime: "2024-12-02T15:00:00",
      endTime: "2024-12-02T17:30:00"
    };

    const { baseElement } = render(<EventCard {...props} />);
    expect(baseElement).toBeTruthy();
  });
});
