// @ts-check
const inputError = 'w3-border-red';
const bgSuccess = 'w3-green';
const bgError = 'w3-red';

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
      input.classList.add('w3-border-red');
      if (errorMessage) {
        errorMessage.textContent = error;
      }
    }
  });
}

function clearErrors() {
  const errorInputs = document.querySelectorAll('.' + inputError);
  errorInputs.forEach((input) => {
    input.classList.remove(inputError);
    if (input.nextElementSibling) {
      input.nextElementSibling.textContent = '';
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
      notification.classList.remove(bgError);
      notification.classList.add(bgSuccess);
    } else {
      notification.classList.remove(bgSuccess);
      notification.classList.add(bgError);
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
  const formJSON = Object.fromEntries(formData.entries());
  const actionUrl = this.getAttribute('action');
  /** @type {HTMLInputElement|null} */
  const methodInput = this.querySelector('input[name="_method"]');
  const method = methodInput?.value.toUpperCase() || 'POST'; // Use PUT if specified, else POST

  try {
    // @ts-ignore
    const response = await fetch(actionUrl, {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formJSON),
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
