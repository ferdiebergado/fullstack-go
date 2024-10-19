// @ts-check

import { submitForm } from '../form';

/** @type {HTMLFormElement} */
const editActivityForm = document.getElementById('edit-activity-form');

editActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();
  submitForm(this, () => {});
});
