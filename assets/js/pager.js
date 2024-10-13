// @ts-check
/**
 * @typedef {Object} ResponseData
 * @property {Object[]} data
 */

/**
 * @typedef {Object} PaginationData
 * @property {ResponseData} data
 */

const pagination = document.querySelector('.pagination');

if (pagination) {
  const links = pagination.getElementsByTagName('a');
  for (let index = 0; index < links.length; index++) {
    const link = links[index];

    link.addEventListener('click', async function (event) {
      event.preventDefault();

      try {
        // @ts-ignore
        const response = await fetch(link.href, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/json',
          },
        });

        /** @type {ApiResponse} */
        const { message, data } = await response.json();

        if (!response.ok) {
          // Display validation errors if available

          showNotification(message, 'error');
          return;
        }

        renderData(data);
      } catch (error) {
        showNotification('An error occurred. Please try again.', 'error');
      }
    });
  }
}

/**
 *
 * @param {ResponseData} data
 */
function renderData(data) {
  const table = document.querySelector('table');

  if (table) {
    const tbody = table.tBodies[0];

    removeChildren(tbody);

    data.data.forEach((d) => {
      /** @type {HTMLTemplateElement} */
      // @ts-ignore
      const template = document
        .getElementById('activity-row')
        .content.cloneNode(true);

      if (template) {
        template.querySelector('.title').textContent = d.title;
        template.querySelector('.start_date').textContent = d.start_date;
        template.querySelector('.end_date').textContent = d.end_date;
        template.querySelector('.venue').textContent = d.venue;
        template.querySelector('.region').textContent = d.region;
        template.querySelector('.host').textContent = d.host;
        template.querySelector('.info').href = `/activities/${d.id}`;
        template.querySelector('.view').href = `/activities/${d.id}/edit`;
        tbody.appendChild(template);
      }
      // const row = document.createElement('tr');

      // addCol(row, d.title);
      // addCol(row, d.start_date);
      // addCol(row, d.end_date);
      // addCol(row, d.venue);
      // addCol(row, d.region);
      // addCol(row, d.host);

      // const col = document.createElement('td');
      // const infoLink = document.createElement('a');
      // infoLink.href = `/activities/${d.id}`;
      // infoLink.textContent = 'Info';
      // const viewLink = document.createElement('a');
      // viewLink.href = `/activities/${d.id}/edit`;
      // viewLink.textContent = 'Edit';
      // row.appendChild(col);

      // tbody.appendChild(row);
    });
  }
}

/**
 *
 * @param {ResponseData} data
 */
function updateLinks(data) {}

/**
 * Removes all the children node of the element.
 * @param {HTMLElement} element
 */
function removeChildren(element) {
  while (element.firstChild) {
    element.removeChild(element.firstChild);
  }
}

/**
 *
 * @param {HTMLTableRowElement} row
 * @param {string} value
 */
function addCol(row, value) {
  const col = document.createElement('td');
  col.textContent = value;
  row.appendChild(col);
}
