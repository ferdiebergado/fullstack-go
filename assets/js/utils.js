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

/**
 *
 * @param {string} originalText
 * @param {number} maxLength
 *
 * @returns {string}
 */
export function truncateText(originalText, maxLength) {
  if (originalText.length > maxLength) {
    return originalText.substring(0, maxLength) + '...';
  }

  return originalText;
}

/**
 * Highlight text from a given text.
 *
 * @param {string} text
 * @param {string} searchTerm
 * @returns {string}
 */
export function highlightText(text, searchTerm) {
  if (!searchTerm) return text;

  // Escape special characters in the text to highlight
  const escapedText = searchTerm.replace(/[-\/\\^$.*+?()[\]{}|]/g, '\\$&');
  // Create a case-insensitive regular expression
  const regex = new RegExp(`(${escapedText})`, 'gi');
  // Replace matches with highlighted span
  const highlightedText = text.replace(regex, '<mark>$1</mark>');

  return highlightedText;
}
