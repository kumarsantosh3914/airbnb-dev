// Jest global setup for the project
// - Silence noisy console during tests (you can remove if you want logs)
jest.spyOn(console, 'error').mockImplementation(() => {});
jest.spyOn(console, 'warn').mockImplementation(() => {});

// Ensure consistent timezone in tests
process.env.TZ = 'UTC';

// Common env defaults for tests (add as needed)
process.env.NODE_ENV = process.env.NODE_ENV || 'test';
