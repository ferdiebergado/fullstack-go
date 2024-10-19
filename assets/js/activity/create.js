// @ts-check

import { submitForm } from '../form.js';
import { handleHostForm, watchHost } from '../host.js';
import { handleVenueForm, watchVenue } from '../venue.js';

/** @type {HTMLFormElement} */
const createActivityForm = document.getElementById('create-activity-form');

createActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();

  submitForm(this, () => {});
});

handleVenueForm();
handleHostForm();
watchVenue();
watchHost();
