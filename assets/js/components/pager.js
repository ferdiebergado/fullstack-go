// @ts-check
import { showNotification } from './notification.js';

/**
 * Sets up pagination link handling.
 */
// @ts-check
export function setupPagination() {
  const paginationContainer = document.querySelector('.pagination');

  if (paginationContainer) {
    paginationContainer.addEventListener('click', async (event) => {
      event.preventDefault();
      const target = event.target;

      // Only proceed if the clicked element is a pagination link
      if (target?.tagName === 'A') {
        try {
          const response = await fetch(target?.href, {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
              Accept: 'application/json',
            },
          });

          const { message, data } = await response.json();
          if (!response.ok) {
            showNotification(message, 'error');
          } else {
            renderData(data);
          }
        } catch {
          showNotification('An error occurred. Please try again.', 'error');
        }
      }
    });
  }
}

/**
 * Renders the table data.
 * @param {Object} data
 */
function renderData(data) {
  const table = document.querySelector('table');
  const tbody = table?.tBodies[0];

  if (tbody) {
    tbody.innerHTML = '';

    const fragment = document.createDocumentFragment();

    data.data.forEach((d) => {
      /** @type {HTMLTemplateElement} */
      const template = document
        .getElementById('activity-row')
        ?.content.cloneNode(true);
      if (template) {
        template.querySelector('.title').textContent = d.title;
        template.querySelector('.start_date').textContent = d.start_date;
        template.querySelector('.end_date').textContent = d.end_date;
        template.querySelector('.venue').textContent = d.venue;
        template.querySelector('.region').textContent = d.region;
        template.querySelector('.host').textContent = d.host;
        template.querySelector('.info').href = `/activities/${d.id}`;
        template.querySelector('.view').href = `/activities/${d.id}/edit`;

        fragment.appendChild(template);
      }
    });

    tbody.appendChild(fragment);
  }
}

/**
 * Removes all the child nodes of an element.
 * @param {HTMLElement} element
 */
function removeChildren(element) {
  while (element.firstChild) {
    element.removeChild(element.firstChild);
  }
}
