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

const links = pagination?.getElementsByTagName('a');

if (links) {
  for (let index = 0; index < links.length; index++) {
    const link = links[index];

    if (link) {
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
}

/**
 *
 * @param {ResponseData} data
 */
function renderData(data) {
  const table = document.querySelector('table');

  if (table) {
    const tbody = table.tBodies[0];
    tbody.innerHTML = '';
    data.data.forEach((d) => {
      const row = document.createElement('tr');
      const titleCol = document.createElement('td');
      titleCol.textContent = d.title;
      row.appendChild(titleCol);

      const startDateCol = document.createElement('td');
      startDateCol.textContent = d.start_date;
      row.appendChild(startDateCol);

      const endDateCol = document.createElement('td');
      endDateCol.textContent = d.end_date;
      row.appendChild(endDateCol);

      const venueCol = document.createElement('td');
      venueCol.textContent = d.venue;
      row.appendChild(venueCol);

      const regionCol = document.createElement('td');
      regionCol.textContent = d.region;
      row.appendChild(regionCol);

      const hostCol = document.createElement('td');
      hostCol.textContent = d.host;
      row.appendChild(hostCol);

      tbody.appendChild(row);
    });
  }
}
