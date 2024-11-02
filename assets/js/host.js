// @ts-check
'use strict';

/**
 * @typedef {Object} Host
 * @property {number} id
 * @property {string} name
 */

/**
 * @typedef {Host[]} Hosts
 */
import { mountDialogForSelect } from './components/dialog.js';
import { updateSelect } from './components/select.js';
import { submitForm } from './form.js';

const createHostForm = /** @type {HTMLFormElement|null} */ (
  document.getElementById('create-host-form')
);
const hostSelect = /** @type {HTMLSelectElement|null} */ (
  document.getElementById('host_id')
);
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

    submitForm(
      this,
      /** @param {Hosts} data */ (data) => {
        hostSelect?.dispatchEvent(
          new CustomEvent('HostCreated', { detail: data[0] })
        );

        hostDialog.close();
      }
    );
  });
}
