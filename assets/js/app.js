// @ts-check

// Get the current URL path
const currentPath = window.location.pathname;

// Extract the root path (first two segments)
const rootPath = '/' + currentPath.split('/')[1];

// Get all anchor tags in the topnav
const navLinks = document.querySelectorAll('.topnav a');

// Remove 'active' class from all links first
navLinks.forEach((link) => {
  link.classList.remove('active');
});

navLinks.forEach((link) => {
  // Check if the href attribute matches the current URL path
  if (rootPath === link.getAttribute('href') && currentPath !== '/') {
    link.classList.add('active'); // Add the 'active' class
  }
});

/**
 * Displays notification on success or error.
 * @param {string} message
 * @param {string} type
 */
function showNotification(message, type) {
  const notification = document.getElementById('notification');

  if (notification) {
    const header = notification.querySelector('#notification-header');
    const body = notification.querySelector('#notification-message');

    if (type === 'success') {
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }

    if (header) header.textContent = type.toUpperCase();

    if (body) body.textContent = message;

    notification.style.display = 'block';
  }
}
