import { render } from '@testing-library/react';

import ScheduliiComponents from './schedulii-components';

describe('ScheduliiComponents', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<ScheduliiComponents />);
    expect(baseElement).toBeTruthy();
  });
});
