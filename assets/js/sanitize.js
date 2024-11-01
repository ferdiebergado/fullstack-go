// @ts-check
'use strict';

/**
 * Converts special HTML characters into their corresponding HTML entities,
 * which helps prevent XSS (Cross-Site Scripting) attacks
 * when displaying user-generated content.
 *
 * @param {FormDataEntryValue|string} input
 * @returns {string|Blob}
 */
export function sanitize(input) {
  if (typeof input === 'string') {
    // Create a div element to leverage the browser's HTML parsing
    const div = document.createElement('div');

    // Set the text content of the div to the input value
    // This automatically escapes any HTML tags
    div.textContent = input;

    // Return the sanitized string
    return div.innerHTML;
  } else if (input instanceof File) {
    return input;
  }

  return '';
}
