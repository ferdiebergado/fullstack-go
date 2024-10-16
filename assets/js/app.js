// @ts-check

// Get all anchor tags in the topnav
const navLinks = document.querySelectorAll('.top-nav a');

navLinks.forEach((link) => {
  // Check if the href attribute matches the current URL path
  link.addEventListener('click', function() {
    // Get the current URL path
    const currentPath = window.location.pathname;

    // Extract the root path (first two segments)
    const rootPath = '/' + currentPath.split('/')[1];

    // Remove 'active' class from all links first
    navLinks.forEach((link) => {
      link.classList.remove('active');
    });

    if (rootPath === link.getAttribute('href') && currentPath !== '/') {
      link.classList.add('active'); // Add the 'active' class
    }

  })
});

/**
 * Displays notification on success or error.
 * @param {string} message
 * @param {string} type
 */
function showNotification(message, type) {
  const notification = document.getElementById('notification');

  if (notification) {
    const header = notification.getElementById('notification-header');
    const body = notification.getElementById('notification-message');

    if (type === 'success') {
      if (header) header.textContent = 'Action Completed';
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      if (header) header.textContent = 'Action Failed';
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    if (body) body.textContent = message;

    notification.style.display = 'block';
  }
}


/**
 * Shows the appropriate dialog.
 *
 * @param {string} dialogId - The id of the dialog element
 * @param {HTMLSelectElement} select - The select element
 */
function mountDialogForSelect(dialogId, select) {
  /** @type {HTMLDialogElement} */
  const dialog = document.getElementById(dialogId);

  const dialogClose = dialog.querySelector('#dialog-close');

  dialogClose?.addEventListener('click', function () {
    if (dialog) dialog.close();
  });

  window.addEventListener('click', function (event) {
    if (event.target === dialog) dialog.close();
  });

  select?.addEventListener('change', function () {
    showDialog();
  });

  select?.addEventListener('click', function () {
    showDialog();
  });

  function showDialog() {
    const index = select.selectedIndex;
    const optionValue = select.options[index].value;

    if (optionValue === 'add') {
      if (dialog) dialog.showModal();
    }
  }

  return dialog;
}

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
const successBgClass = 'alert-success';
const errorBgClass = 'alert-error';
const errorTextClass = 'form-error';

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

    /** @type {ApiResponse} */
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
   * @param {ValidationError[]} errors - Array of validation error objects.
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

/**
 * Updates the select options with new data.
 *
 * @param {HTMLSelectElement} select
 * @param {CustomEventDetail} detail
 */
function updateSelect(select, detail) {
  // Step 1: Get all options excluding the first option (with an empty value)
  let options = Array.from(select.options).slice(1); // Skip the first option

  // Step 2: Create a new option and add it to the array
  const newOption = document.createElement('option');
  newOption.value = detail.id.toString();
  newOption.text = detail.name;
  newOption.selected = true;
  options.push(newOption);

  // Step 3: Sort the options by their displayed text
  options.sort((a, b) => a.text.localeCompare(b.text));

  // Step 4: Remove all options except the first one (which has an empty value)
  const firstOption = select.options[0]; // Retain the first option
  const secondOption = select.options[1]; // Retain the first option
  select.innerHTML = ''; // Clear all options
  select.add(firstOption); // Add back the first option
  select.add(secondOption);

  // Step 5: Append the sorted options after the first option
  options.forEach((option) => select.add(option));
}
