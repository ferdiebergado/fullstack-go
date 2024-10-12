// @ts-check
/** @type {HTMLFormElement} */
const createActivityForm = document.getElementById('create-activity-form');

/** @type {HTMLFormElement} */
const editActivityForm = document.getElementById('edit-activity-form');

/** @type {HTMLFormElement} */
const createVenueForm = document.getElementById('create-venue-form');

/** @type {HTMLFormElement} */
const createHostForm = document.getElementById('create-host-form');

/** @type {HTMLSelectElement | null} */
const venueSelect = document.getElementById('venue_id');

/** @type {HTMLSelectElement | null} */
const hostSelect = document.getElementById('host_id');

const hostDialog = mountDialogForSelect('create-host-dialog', hostSelect);

hostSelect?.addEventListener(
  'HostCreated',

  /** @param {MyCustomEventInit} event */
  function (event) {
    updateSelect(hostSelect, event.detail);
  }
);

const venueDialog = mountDialogForSelect('create-venue-dialog', venueSelect);

venueSelect?.addEventListener(
  'VenueCreated',

  /** @param {MyCustomEventInit} event */
  function (event) {
    updateSelect(venueSelect, event.detail);
  }
);

createActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();

  sendRequest(createActivityForm, function () {});
});

editActivityForm?.addEventListener('submit', function (event) {
  event.preventDefault();

  sendRequest(editActivityForm, function () {});
});

createVenueForm.addEventListener('submit', function (event) {
  event.preventDefault();

  sendRequest(createVenueForm, function (data) {
    const venueCreated = new CustomEvent('VenueCreated', { detail: data });
    venueSelect?.dispatchEvent(venueCreated);
    venueDialog.close();
  });
});

createHostForm.addEventListener('submit', function (event) {
  event.preventDefault();

  sendRequest(createHostForm, function (data) {
    const hostCreated = new CustomEvent('HostCreated', { detail: data });
    hostSelect?.dispatchEvent(hostCreated);
    hostDialog.close();
  });
});
