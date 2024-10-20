// @ts-check
import { inputErrorClass } from './config.js';
import { showNotification } from './components/notification.js';

/**
 * Sends a request and handles form submission.
 * @param {HTMLFormElement} form
 * @param {Function} cb
 */
export async function submitForm(form, cb) {
  clearFormErrors(form);

  const formData = new FormData(form);
  const actionUrl = form.getAttribute('action');

  /** @type {HTMLInputElement|null} */
  const methodInput = form.querySelector('input[name="_method"]');

  let method = 'POST';

  if (methodInput) method = methodInput.value.toUpperCase();

  // Convert FormData to a plain object
  const payload = {};

  formData.forEach((value, key) => {
    // Convert numeric fields manually
    if (key.endsWith('_id')) {
      payload[key] = Number(value); // Convert to number
    } else {
      payload[key] = value; // Keep as string
    }
  });

  try {
    const response = await fetch(actionUrl, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    /** @type {import('./typedefs').ApiResponse} */
    const { errors, message, data } = await response.json();

    if (!response.ok) {
      if (errors) displayFormErrors(form, errors);

      showNotification(message, 'error');
    } else {
      showNotification(message, 'success');

      if (method !== 'PUT') form.reset();
      cb(data);
    }
  } catch (error) {
    showNotification('An error occurred. Please try again.', 'error');
  }
}

/**
 *
 * @param {HTMLFormElement} form
 * @param {import("./typedefs").ValidationError[]} errors
 */
function displayFormErrors(form, errors) {
  errors.forEach(({ field, error }) => {
    const input = form.querySelector(`[name="${field}"]`);
    if (input) {
      const helpText = input.nextElementSibling;
      input.classList.add(inputErrorClass);
      if (helpText) helpText.textContent = error;
    }
  });
}

/**
 * Removes the error styles from the form inputs.
 * @param {HTMLFormElement} form
 */
function clearFormErrors(form) {
  form.querySelectorAll('.' + inputErrorClass).forEach((input) => {
    input.classList.remove(inputErrorClass);
    const nextEl = input.nextElementSibling;
    if (nextEl) nextEl.textContent = '';
  });
}
