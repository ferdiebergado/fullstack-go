// @ts-check
import { inputErrorClass } from './config';
import showNotification from './components/notification';

/**
 *
 * @param {HTMLFormElement} form
 * @param {Function} cb
 */
async function sendRequest(form, cb) {
  // Clear previous error styles and messages
  clearErrors();

  const formData = new FormData(form);
  // const formJSON = Object.fromEntries(formData.entries());
  const actionUrl = form.getAttribute('action');
  /** @type {HTMLInputElement|null} */
  const methodInput = form.querySelector('input[name="_method"]');
  const method = methodInput?.value.toUpperCase() || 'POST'; // Use PUT if specified, else POST

  // Convert FormData to a plain object
  const payload = {};
  formData.forEach((value, key) => {
    // Convert numeric fields manually
    console.log(key, value);

    if (key.endsWith('_id')) {
      payload[key] = Number(value); // Convert to number
    } else {
      payload[key] = value; // Keep as string
    }
  });

  try {
    // @ts-ignore
    const response = await fetch(actionUrl, {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    });

    /** @type {import('./typedefs').ApiResponse} */
    const { errors, message, data } = await response.json();

    if (!response.ok) {
      // Display validation errors if available
      if (errors) displayErrors(errors);

      showNotification(message, 'error');
    } else {
      showNotification(message, 'success');

      if (method !== 'PUT') form.reset(); // Clear the form on success

      cb(data);
    }
  } catch (error) {
    showNotification('An error occurred. Please try again.', 'error');
  }

  /**
   * Handles form validation errors.
   * @param {import('./typedefs').ValidationError[]} errors - Array of validation error objects.
   */
  function displayErrors(errors) {
    errors.forEach(({ field, error }) => {
      const input = form.querySelector(`[name="${field}"]`);

      if (input) {
        const helpText = input.nextElementSibling;

        if (!input.classList.contains(inputErrorClass)) {
          input.classList.add(inputErrorClass);
        }

        if (helpText) helpText.textContent = error;
      }
    });
  }

  // Removes the error classes from forms with errors.
  function clearErrors() {
    const inputsWithError = form.querySelectorAll('.' + inputErrorClass);

    inputsWithError.forEach((input) => {
      input.classList.remove(inputErrorClass);
      const helpText = input.nextElementSibling;

      if (helpText) helpText.textContent = '';
    });
  }
}

export default sendRequest;
