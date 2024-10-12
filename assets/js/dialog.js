// @ts-check
/**
 * Shows the appropriate dialog.
 *
 * @param {string} dialogId - The id of the dialog element
 * @param {HTMLSelectElement} select - The select element
 */
function mountDialogForSelect(dialogId, select) {
  /** @type {HTMLDialogElement} */
  const dialog = document.getElementById(dialogId);

  const dialogClose = dialog.querySelector('#dialog-close');

  dialogClose?.addEventListener('click', function () {
    if (dialog) dialog.close();
  });

  window.addEventListener('click', function (event) {
    if (event.target === dialog) dialog.close();
  });

  select?.addEventListener('change', function () {
    showDialog();
  });

  select?.addEventListener('click', function () {
    showDialog();
  });

  function showDialog() {
    const index = select.selectedIndex;
    const optionValue = select.options[index].value;

    if (optionValue === 'add') {
      if (dialog) dialog.showModal();
    }
  }

  return dialog;
}
