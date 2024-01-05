import { render, screen, userEvent } from 'design/utils/testing';

import { AdvancedSearchToggleProps } from '../types';

import { AdvancedSearchToggle } from './AdvancedSearchToggle';

test('sets optional fields and calls cb', async () => {
  const mockToggle = jest.fn();
  const props: AdvancedSearchToggleProps = {
    isToggled: false,
    onToggle: mockToggle,
    px: 1, // 4px
    gap: 6, // 40px
    className: 'example-class',
  };

  render(<AdvancedSearchToggle {...props} />);
  expect(screen.getByLabelText('Advanced')).not.toBeChecked();
  expect(screen.getByTestId('wrapper')).toHaveClass('example-class');
  expect(screen.getByTestId('wrapper')).toHaveStyle('gap: 40px');
  expect(screen.getByTestId('wrapper')).toHaveStyle('padding-left: 4px');
  expect(screen.getByTestId('wrapper')).toHaveStyle('padding-right: 4px');

  await userEvent.click(screen.getByLabelText('Advanced'));
  expect(mockToggle).toHaveBeenCalled();
});

test('checks toggle if true', async () => {
  const props: AdvancedSearchToggleProps = {
    isToggled: true,
    onToggle: () => {},
  };

  render(<AdvancedSearchToggle {...props} />);
  expect(screen.getByLabelText('Advanced')).toBeChecked();
  expect(screen.getByTestId('wrapper')).toHaveStyle('gap: 8px'); // default
});
