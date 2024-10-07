// @ts-check
const inputErrorClass = 'has-error';
const successBgClass = 'alert--success';
const errorBgClass = 'alert--danger';
const errorTextClass = 'form__error';

// const inputError = 'w3-border-red';
// const bgSuccess = 'w3-green';
// const bgError = 'w3-red';

const notification = document.getElementById('notification');

/**
 * Represents an error object.
 * @typedef {Object} ValidationError
 * @property {string} field - The field that has the validation error.
 * @property {string} error - The error message associated with the field.
 */

/**
 * Handles form validation errors.
 * @param {ValidationError[]} errors - Array of validation error objects.
 */
function displayErrors(errors) {
  errors.forEach(({ field, error }) => {
    const input = document.querySelector(`[name="${field}"]`);

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

function clearErrors() {
  const errorInputs = document.querySelectorAll('.' + inputErrorClass);
  errorInputs.forEach((input) => {
    input.classList.remove(inputErrorClass);
    const helpText = input.querySelector('.' + errorTextClass);

    if (helpText) {
      helpText.textContent = '';
    }
  });
}

/**
 *
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

const form = document.getElementsByTagName('form')[0];

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

    const data = await response.json();

    if (!response.ok) {
      // Display validation errors if available
      if (data.errors) {
        displayErrors(data.errors);
      }

      showNotification('Submission failed. Please try again.', 'error');
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
