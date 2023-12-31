import { render } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import EventForm from './event-form';

describe('EventForm', () => {
  it('should render successfully', () => {
    const { baseElement } = render(
      <EventForm
        formData={{
          eventName: '',
          startDate: '',
          endDate: '',
          earliestTime: '',
          latestTime: '',
          location: '',
        }}
        isFormValid={false}
        onInputChange={() => {}}
        onSubmit={() => {}}
      />,
      { wrapper: BrowserRouter }
    );
    expect(baseElement).toBeTruthy();
  });
});
