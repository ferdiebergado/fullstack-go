// @ts-check

import { mountDialogForSelect } from './components/dialog.js';
import { submitForm } from './form.js';
import { updateSelect } from './components/select.js';

/** @type {HTMLFormElement} */
const createVenueForm = document.getElementById('create-venue-form');

/** @type {HTMLSelectElement} */
const venueSelect = document.getElementById('venue_id');

export function handleVenueForm() {
  createVenueForm?.addEventListener('submit', function (event) {
    event.preventDefault();

    submitForm(this, (data) => {
      venueSelect?.dispatchEvent(
        new CustomEvent('VenueCreated', { detail: data })
      );
      venueDialog.close();
    });
  });
}

const venueDialog = mountDialogForSelect('create-venue-dialog', venueSelect);

export function watchVenue() {
  venueSelect?.addEventListener(
    'VenueCreated',
    /** @param {import('./typedefs.js').MyCustomEventInit} event */ function (
      event
    ) {
      updateSelect(this, event.detail);
    }
  );
}
