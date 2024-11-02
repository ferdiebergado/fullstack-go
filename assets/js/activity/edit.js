// @ts-check
'use strict';

import { submitForm } from '../form';

const editActivityForm = /** @type {HTMLFormElement|null} */ (
  document.getElementById('edit-activity-form')
);

editActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();
  submitForm(this, () => {});
});
