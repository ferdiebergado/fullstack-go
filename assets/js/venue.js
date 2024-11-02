// @ts-check
'use strict';

/**
 * @typedef {Object} Venue
 * @property {number} id
 * @property {string} name
 * @property {number} division_id
 */

/**
 * @typedef {Venue[]} Venues
 */

import { mountDialogForSelect } from './components/dialog.js';
import { submitForm } from './form.js';
import { updateSelect } from './components/select.js';

const createVenueForm = /** @type {HTMLFormElement} */ (
  document.getElementById('create-venue-form')
);
const venueSelect = /** @type {HTMLSelectElement|null} */ (
  document.getElementById('venue_id')
);
const venueDialog = mountDialogForSelect('create-venue-dialog', venueSelect);

export function handleVenueForm() {
  createVenueForm?.addEventListener('submit', function (event) {
    event.preventDefault();

    submitForm(
      this,
      /** @param {Venues} data */ (data) => {
        venueSelect?.dispatchEvent(
          new CustomEvent('VenueCreated', { detail: data[0] })
        );
        venueDialog.close();
      }
    );
  });
}

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
