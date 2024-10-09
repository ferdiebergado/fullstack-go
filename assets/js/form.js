// @ts-check
const inputErrorClass = 'has-error';
const successBgClass = 'alert--success';
const errorBgClass = 'alert--danger';
const errorTextClass = 'form__error';

const forms = document.querySelectorAll('form');
const notification = document.getElementById('notification');

/**
 * Represents an API response.
 * @typedef {Object} ApiResponse
 * @property {boolean} success - The status of the response
 * @property {string} message - The message
 * @property {ValidationError[]} errors - The object that contains the errors
 * @property {Object} data - The object that contains the data
 */

/**
 * Represents an error object.
 * @typedef {Object} ValidationError
 * @property {string} field - The field that has the validation error.
 * @property {string} error - The error message associated with the field.
 */

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

    if (header) {
      header.textContent = type.toUpperCase();
    }

    if (body) {
      body.textContent = message;
    }

    notification.style.display = 'block';
  }
}

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
      const data = await response.json();

      if (!response.ok) {
        // Display validation errors if available
        console.log(data);

        const { errors, message } = data;
        console.log(message);

        if (errors) {
          displayErrors(errors);
        }

        showNotification(message, 'error');
      } else {
        showNotification('Form submitted successfully!', 'success');
        if (method !== 'PUT') {
          form.reset(); // Clear the form on success
        }
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

        if (errorMessage) {
          errorMessage.textContent = error;
        }
      }
    });
  }

  // Removes the error classes from forms with errors.
  function clearErrors() {
    const errorInputs = form.querySelectorAll('.' + inputErrorClass);
    errorInputs.forEach((input) => {
      input.classList.remove(inputErrorClass);
      const helpText = input.querySelector('.' + errorTextClass);

      if (helpText) {
        helpText.textContent = '';
      }
    });
  }
});
