// @ts-check
import { errorBgClass, successBgClass } from '../config';

/**
 * Displays notification on success or error.
 * @param {string} message
 * @param {string} type
 */
function showNotification(message, type) {
  const notification = document.getElementById('notification');
  const header = document.getElementById('notification-header');
  const body = document.getElementById('notification-message');

  if (notification) {
    if (type === 'success') {
      if (header) header.textContent = 'Action Completed';
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      if (header) header.textContent = 'Action Failed';
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    if (body) body.textContent = message;

    notification.style.display = 'block';
  }
}

export default showNotification;
