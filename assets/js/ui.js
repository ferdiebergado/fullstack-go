// @ts-check
export const inputErrorClass = 'has-error';
export const successBgClass = 'alert-success';
export const errorBgClass = 'alert-error';

/**
 * Initializes the top navigation links.
 */
export function initializeNavLinks() {
  const navLinks = document.querySelectorAll('.top-nav a');
  navLinks.forEach((link) => {
    link.addEventListener('click', function () {
      const currentPath = window.location.pathname;
      const rootPath = '/' + currentPath.split('/')[1];

      navLinks.forEach((link) => link.classList.remove('active'));
      if (rootPath === link.getAttribute('href') && currentPath !== '/') {
        link.classList.add('active');
      }
    });
  });
}

/**
 * Displays a notification on success or error.
 * @param {string} message
 * @param {string} type
 */
export function showNotification(message, type) {
  const notification = document.getElementById('notification');
  const header = document.getElementById('notification-header');
  const body = document.getElementById('notification-message');

  if (notification && header && body) {
    if (type === 'success') {
      header.textContent = 'Action Completed';
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      header.textContent = 'Action Failed';
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    body.textContent = message;
    notification.style.display = 'block';
  }
}

/**
 * Mounts a dialog for select elements.
 * @param {string} dialogId
 * @param {HTMLSelectElement} select
 */
export function mountDialogForSelect(dialogId, select) {
  /** @type {HTMLDialogElement} */
  const dialog = document.getElementById(dialogId);
  const dialogClose = dialog.querySelector('#dialog-close');

  dialogClose?.addEventListener('click', () => dialog?.close());
  window.addEventListener('click', (event) => {
    if (event.target === dialog) dialog.close();
  });

  select?.addEventListener('change', showDialog);
  select?.addEventListener('click', showDialog);

  function showDialog() {
    const optionValue = select.options[select.selectedIndex].value;
    if (optionValue === 'add') {
      dialog?.showModal();
    }
  }

  return dialog;
}

/**
 * Sends a request and handles form submission.
 * @param {HTMLFormElement} form
 * @param {Function} cb
 */
export async function sendRequest(form, cb) {
  clearErrors(form);

  const formData = new FormData(form);
  const actionUrl = form.getAttribute('action');

  /** @type {HTMLInputElement} */
  const methodInput = form.querySelector('input[name="_method"]');

  let method = 'POST';

  if (methodInput) method = methodInput.value.toUpperCase();

  const payload = Object.fromEntries(formData.entries());

  try {
    const response = await fetch(actionUrl, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    const { errors, message, data } = await response.json();
    if (!response.ok) {
      displayErrors(form, errors);
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
function displayErrors(form, errors) {
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
function clearErrors(form) {
  form.querySelectorAll('.' + inputErrorClass).forEach((input) => {
    input.classList.remove(inputErrorClass);
    const nextEl = input.nextElementSibling;
    if (nextEl) nextEl.textContent = '';
  });
}

/**
 * Updates the select element with new options.
 * @param {HTMLSelectElement} select
 * @param {Object} detail
 */
export function updateSelect(select, detail) {
  let options = Array.from(select.options).slice(1);
  const newOption = document.createElement('option');
  newOption.value = detail.id.toString();
  newOption.text = detail.name;
  newOption.selected = true;
  options.push(newOption);
  options.sort((a, b) => a.text.localeCompare(b.text));

  const firstOption = select.options[0];
  select.innerHTML = '';
  select.add(firstOption);
  options.forEach((option) => select.add(option));
}
