import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import App from './App';

// Mock global fetch function
global.fetch = jest.fn(() =>
  Promise.resolve({
    ok: true,
    text: () => Promise.resolve('Hello from Go!')
  })
) as jest.Mock;

describe('App component', () => {
  beforeEach(() => {
    // Clear all mocks before each test to ensure test isolation
    (fetch as jest.Mock).mockClear();
  });

  it('fetches and displays data', async () => {
    render(<App />);

    // Wait for the component to update with the fetched data
    // This ensures that asynchronous side effects are completed
    await waitFor(() => {
      // Using regex in getByText for flexible text matching
      expect(screen.getByText(/Data from Go:/i)).toBeInTheDocument();
    });

    // Check that fetch was called exactly once
    // This validates that the network request was made correctly
    expect(fetch).toHaveBeenCalledTimes(1);
  });

  it('handles fetch error', async () => {
    // Mock fetch to simulate an error response
    // This tests the component's error handling capability
    (fetch as jest.Mock).mockImplementationOnce(() =>
      Promise.resolve({
        ok: false,
        status: 404,
        text: () => Promise.resolve('')
      })
    );

    render(<App />);

    // Wait for the component to update with the error message
    // This is crucial for testing the error state of the component
    await waitFor(() => {
      expect(screen.getByText(/Error fetching data:/)).toBeInTheDocument();
    });

    // Ensuring that fetch was called, even in error scenarios
    expect(fetch).toHaveBeenCalledTimes(1);
  });

  // Future testing scenarios could include:
  // - Testing user interactions (like clicks or form submissions).
  // - More complex async behavior.
  // - Snapshot testing for the UI components.
  // - Isolating the component further to test specific functionalities.
});

// Note: Integration into a CI/CD pipeline and use of code quality tools
// like ESLint or Prettier are recommended as part of the development workflow.
