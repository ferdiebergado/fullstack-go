// @ts-check
import { mountDialogForSelect, sendRequest, updateSelect } from './ui.js';

/**
 * Sets up form handling for activity-related forms.
 */
export function setupActivityForms() {
  /** @type {HTMLFormElement} */
  const createActivityForm = document.getElementById('create-activity-form');

  /** @type {HTMLFormElement} */
  const editActivityForm = document.getElementById('edit-activity-form');

  /** @type {HTMLFormElement} */
  const createVenueForm = document.getElementById('create-venue-form');

  /** @type {HTMLFormElement} */
  const createHostForm = document.getElementById('create-host-form');

  /** @type {HTMLSelectElement} */
  const venueSelect = document.getElementById('venue_id');

  /** @type {HTMLSelectElement} */
  const hostSelect = document.getElementById('host_id');

  const hostDialog = mountDialogForSelect('create-host-dialog', hostSelect);
  hostSelect?.addEventListener(
    'HostCreated',
    /** @param {import('./typedefs.js').MyCustomEventInit} event */
    (event) => {
      updateSelect(hostSelect, event.detail);
    }
  );

  const venueDialog = mountDialogForSelect('create-venue-dialog', venueSelect);
  venueSelect?.addEventListener(
    'VenueCreated',
    /** @param {import('./typedefs.js').MyCustomEventInit} event */ (event) => {
      updateSelect(venueSelect, event.detail);
    }
  );

  createActivityForm?.addEventListener('submit', (event) => {
    event.preventDefault();
    sendRequest(createActivityForm, () => {});
  });

  editActivityForm?.addEventListener('submit', (event) => {
    event.preventDefault();
    sendRequest(editActivityForm, () => {});
  });

  createVenueForm?.addEventListener('submit', (event) => {
    event.preventDefault();
    sendRequest(createVenueForm, (data) => {
      venueSelect?.dispatchEvent(
        new CustomEvent('VenueCreated', { detail: data })
      );
      venueDialog.close();
    });
  });

  createHostForm?.addEventListener('submit', (event) => {
    event.preventDefault();
    sendRequest(createHostForm, (data) => {
      hostSelect?.dispatchEvent(
        new CustomEvent('HostCreated', { detail: data })
      );
      hostDialog.close();
    });
  });
}
