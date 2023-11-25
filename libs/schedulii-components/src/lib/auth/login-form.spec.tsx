import { render } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';

import LoginForm from './login-form';

describe('LoginForm', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<LoginForm />, { wrapper: BrowserRouter });
    expect(baseElement).toBeTruthy();
  });
});
