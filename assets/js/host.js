// @ts-check

import { mountDialogForSelect } from './components/dialog.js';
import { updateSelect } from './components/select.js';
import { submitForm } from './form.js';

/** @type {HTMLFormElement} */
const createHostForm = document.getElementById('create-host-form');

/** @type {HTMLSelectElement} */
const hostSelect = document.getElementById('host_id');

const hostDialog = mountDialogForSelect('create-host-dialog', hostSelect);

export function watchHost() {
  hostSelect?.addEventListener(
    'HostCreated',
    /** @param {import('./typedefs.js').MyCustomEventInit} event */
    function (event) {
      updateSelect(this, event.detail);
    }
  );
}

export function handleHostForm() {
  createHostForm?.addEventListener('submit', function (event) {
    event.preventDefault();

    submitForm(this, (data) => {
      hostSelect?.dispatchEvent(
        new CustomEvent('HostCreated', { detail: data })
      );
      hostDialog.close();
    });
  });
}
