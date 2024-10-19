// @ts-check
/**
 * Mounts a dialog for select elements.
 * @param {string} dialogId
 * @param {HTMLSelectElement} select
 */
export function mountDialogForSelect(dialogId, select) {
  /** @type {HTMLDialogElement} */
  const dialog = document.getElementById(dialogId);
  const dialogClose = dialog.querySelector('#dialog-close');

  dialogClose?.addEventListener('click', () => dialog?.close());
  window.addEventListener('click', (event) => {
    if (event.target === dialog) dialog.close();
  });

  select?.addEventListener('change', showDialog);
  select?.addEventListener('click', showDialog);

  function showDialog() {
    const optionValue = select.options[select.selectedIndex].value;
    if (optionValue === 'add') {
      dialog?.showModal();
    }
  }

  return dialog;
}
