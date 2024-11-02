// @ts-check
'use strict';

/**
 * Creates a debounced function that delays invoking the provided function
 * until after a specified delay in milliseconds has elapsed since the last
 * time the debounced function was invoked.
 *
 * @param {Function} func - The function to debounce.
 * @param {number} delay - The number of milliseconds to delay.
 * @returns {...*} A new debounced function that takes the same parameters as `func`.
 */
export function debounce(func, delay) {
  let timeout; // Holds the timeout ID

  /**
   * The inner function that will be called after the delay.
   *
   * @param {...*} args - The arguments to pass to the original function.
   */
  return function (...args) {
    // The returned function accepts any arguments
    clearTimeout(timeout); // Clear the previous timeout
    timeout = setTimeout(() => func.apply(this, args), delay); // Set a new timeout
  };
}
