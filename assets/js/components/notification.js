// @ts-check
import { errorBgClass, successBgClass } from '../config.js';

/**
 * Displays a notification on success or error.
 * @param {string} message
 * @param {string} type
 */
export function showNotification(message, type) {
  const notification = document.getElementById('notification');
  const header = document.getElementById('notification-header');
  const body = document.getElementById('notification-message');

  if (notification && header && body) {
    if (type === 'success') {
      header.textContent = 'Action Completed';
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      header.textContent = 'Action Failed';
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    body.textContent = message;
    notification.style.display = 'block';
  }
}
