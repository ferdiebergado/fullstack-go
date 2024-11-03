// @ts-check
'use strict';

import { deepEqual, formDataToObject, getFormData, submitForm } from '../form';
import { handleHostForm, watchHost } from '../host.js';
import { handleVenueForm, watchVenue } from '../venue.js';

const editActivityForm = /** @type {HTMLFormElement|null} */ (
  document.getElementById('edit-activity-form')
);

const data = getFormData(editActivityForm);

editActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();

  const formData = new FormData(this);
  const formDataObject = formDataToObject(formData);
  const isUnchanged = deepEqual(data, formDataObject);

  !isUnchanged && submitForm(this, () => {});
});

handleVenueForm();
handleHostForm();
watchVenue();
watchHost();
