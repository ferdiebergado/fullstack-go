// @ts-check
/**
 * Represents an API response.
 * @typedef {Object} ApiResponse
 * @property {boolean} success - The status of the response
 * @property {string} message - The message
 * @property {ValidationError[]} errors - The object that contains the errors
 * @property {Object} data - The object that contains the data
 */

/**
 * Represents a validation error object.
 * @typedef {Object} ValidationError
 * @property {string} field - The field that has the validation error.
 * @property {string} error - The error message associated with the field.
 */

/**
 * @typedef {Object} CustomEventDetail
 * @property {number} id - The unique identifier.
 * @property {string} name - The name associated with the event.
 */

/**
 * @typedef {CustomEventInit} MyCustomEventInit
 * @property {CustomEventDetail} detail - Custom detail object for the event.
 */

const inputErrorClass = 'has-error';
const successBgClass = 'alert--success';
const errorBgClass = 'alert--danger';
const errorTextClass = 'form__error';

const notification = document.getElementById('notification');
const modal = document.getElementById('modal');

/** @type {HTMLSelectElement | null} */
const venueSelect = document.getElementById('venue_id');
const forms = document.querySelectorAll('form');

/**
 * Displays notification on success or error.
 * @param {string} message
 * @param {string} type
 */
function showNotification(message, type) {
  if (notification) {
    const header = notification.querySelector('#notification-header');
    const body = notification.querySelector('#notification-message');

    if (type === 'success') {
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    if (header) header.textContent = type.toUpperCase();

    if (body) body.textContent = message;

    notification.style.display = 'block';
  }
}

venueSelect?.addEventListener(
  'Created',

  /** @param {MyCustomEventInit} event */
  function (event) {
    // Step 1: Get all options excluding the first option (with an empty value)
    let options = Array.from(venueSelect.options).slice(1); // Skip the first option

    // Step 2: Create a new option and add it to the array
    const newOption = document.createElement('option');
    newOption.value = event.detail.id;
    newOption.text = event.detail.name;
    newOption.selected = true;
    options.push(newOption);

    // Step 3: Sort the options by their displayed text
    options.sort((a, b) => a.text.localeCompare(b.text));

    // Step 4: Remove all options except the first one (which has an empty value)
    const firstOption = venueSelect.options[0]; // Retain the first option
    venueSelect.innerHTML = ''; // Clear all options
    venueSelect.add(firstOption); // Add back the first option

    // Step 5: Append the sorted options after the first option
    options.forEach((option) => venueSelect.add(option));
  }
);

// Attach submit event listener to all forms on the page
forms.forEach((form) => {
  form.addEventListener('submit', async function (event) {
    event.preventDefault();

    // Clear previous error styles and messages
    clearErrors();

    const formData = new FormData(this);
    // const formJSON = Object.fromEntries(formData.entries());
    const actionUrl = this.getAttribute('action');
    /** @type {HTMLInputElement|null} */
    const methodInput = this.querySelector('input[name="_method"]');
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

      /** @type {ApiResponse} */
      const { errors, message, data } = await response.json();

      if (!response.ok) {
        // Display validation errors if available
        if (errors) {
          displayErrors(errors);
        }

        showNotification(message, 'error');
      } else {
        showNotification(message, 'success');

        if (method !== 'PUT') form.reset(); // Clear the form on success
        if (modal) modal.style.display = 'none';

        const Created = new CustomEvent('Created', { detail: data });
        venueSelect?.dispatchEvent(Created);
      }
    } catch (error) {
      showNotification('An error occurred. Please try again.', 'error');
    }
  });

  /**
   * Handles form validation errors.
   * @param {ValidationError[]} errors - Array of validation error objects.
   */
  function displayErrors(errors) {
    errors.forEach(({ field, error }) => {
      const input = form.querySelector(`[name="${field}"]`);

      if (input) {
        const errorMessage = input.nextElementSibling;
        const parent = input.parentElement;

        if (parent) {
          if (!parent.classList.contains(inputErrorClass)) {
            parent.classList.add(inputErrorClass);
          }
        }

        if (errorMessage) errorMessage.textContent = error;
      }
    });
  }

  // Removes the error classes from forms with errors.
  function clearErrors() {
    const errorInputs = form.querySelectorAll('.' + inputErrorClass);
    errorInputs.forEach((input) => {
      input.classList.remove(inputErrorClass);
      const helpText = input.querySelector('.' + errorTextClass);

      if (helpText) helpText.textContent = '';
    });
  }
});
