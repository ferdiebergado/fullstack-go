// @ts-check
'use strict';

import { inputErrorClass } from './config.js';
import { showNotification } from './components/notification.js';
import { sanitize } from './utils.js';

/**
 * Sends a request and handles form submission.
 *
 * @param {HTMLFormElement} form
 * @param {Function} cb
 */
export async function submitForm(form, cb) {
  clearFormErrors(form);

  const formData = new FormData(form);
  const actionUrl = form.getAttribute('action') || '';

  const methodInput = /** @type {HTMLInputElement|null} */ (
    form.querySelector('input[name="_method"]')
  );

  let method = 'POST';

  if (methodInput) method = methodInput.value.toUpperCase();

  // Convert FormData to a plain object
  const payload = {};

  formData.forEach((value, key) => {
    const input = sanitize(value);

    if (typeof input === 'string') {
      // Convert numeric fields manually
      if (key.endsWith('_id')) {
        payload[key] = Number(input); // Convert to number
      } else {
        payload[key] = input; // Keep as string
      }
    }
  });

  try {
    const response = await fetch(actionUrl, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    /** @type {import('./typedefs').ApiResponse} */
    const {
      meta: { message, errors },
      data,
    } = await response.json();

    if (!response.ok) {
      if (errors) displayFormErrors(form, errors);

      showNotification(message, 'error');
    } else {
      showNotification(message, 'success');

      method !== 'PUT' && form.reset();

      cb(data);
    }
  } catch (error) {
    showNotification('An error occurred. Please try again.', 'error');
  }
}

/**
 * Adds error styles to form inputs.
 *
 * @param {HTMLFormElement} form
 * @param {import('./typedefs').ValidationError[]} errors
 */
function displayFormErrors(form, errors) {
  errors.forEach(({ field, error }) => {
    const input = form.querySelector(`[name="${field}"]`);
    if (input) {
      const helpText = input.nextElementSibling;
      input.classList.add(inputErrorClass);
      helpText && (helpText.textContent = error);
    }
  });
}

/**
 * Removes the error styles from the form inputs.
 *
 * @param {HTMLFormElement} form
 */
function clearFormErrors(form) {
  form.querySelectorAll('.' + inputErrorClass).forEach((input) => {
    input.classList.remove(inputErrorClass);
    const nextEl = input.nextElementSibling;
    nextEl && (nextEl.textContent = '');
  });
}

/**
 *
 * @param {HTMLFormElement} form
 * @returns
 */
export function getFormData(form) {
  const formData = {};

  Array.from(form.elements).forEach((element) => {
    if (!element.name) return; // Skip elements without a name attribute

    if (element.type === 'checkbox') {
      formData[element.name] = element.checked;
    } else if (element.type === 'radio') {
      if (element.checked) {
        formData[element.name] = element.value;
      }
    } else {
      formData[element.name] = element.value;
    }
  });

  return formData;
}

/**
 *
 * @param {FormData} formData
 * @returns
 */
export function formDataToObject(formData) {
  const obj = {};
  for (const [key, value] of formData.entries()) {
    obj[key] = value;
  }
  return obj;
}

/**
 *
 * @param {Object} obj1
 * @param {Object} obj2
 * @returns {boolean}
 */
export function deepEqual(obj1, obj2) {
  return JSON.stringify(obj1) === JSON.stringify(obj2);
}
