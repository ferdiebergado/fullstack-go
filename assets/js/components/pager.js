// @ts-check
import { showNotification } from './notification.js';

/** @type {HTMLInputElement} */
const currentPageInput = document.getElementById('currentPage');
const refreshButton = document.getElementById('refresh');
const totalPagesSpan = document.getElementById('totalPages');
/** @type {HTMLSelectElement} */
const recordsPerPageSelect = document.getElementById('recordsPerPage');
const currentRecordsSpan = document.getElementById('currentRecords');
const totalRecordsSpan = document.getElementById('totalRecords');
const tableBody = document.querySelector('#datatable tbody');

let currentPage = parseInt(currentPageInput?.value) || 1;
let totalPages = 1;
let recordsPerPage = parseInt(recordsPerPageSelect?.value);
let totalRecords = 0;

async function fetchData() {
  try {
    const response = await fetch(
      `/api/activities?page=${currentPage}&limit=${recordsPerPage}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
      }
    );

    const { message, data } = await response.json();

    if (!response.ok) {
      showNotification(message, 'error');
    } else {
      totalRecords = data.total_items;
      totalPages = Math.ceil(totalRecords / recordsPerPage);
      renderData(data.data);
      updatePagination();
    }
  } catch {
    showNotification('An error occurred. Please try again.', 'error');
  }
}

/**
 * Renders the table data.
 * @param {Object} data
 */
function renderData(data) {
  if (tableBody) {
    tableBody.innerHTML = '';

    const fragment = document.createDocumentFragment();

    data.forEach((d) => {
      /** @type {HTMLTemplateElement} */
      const template = document
        .getElementById('activity-row')
        ?.content.cloneNode(true);
      if (template) {
        template.querySelector('.title').textContent = d.title;
        template.querySelector('.title')?.setAttribute('title', d.title);
        template.querySelector('.start_date').textContent = d.start_date;
        template.querySelector('.end_date').textContent = d.end_date;
        template.querySelector('.venue').textContent = d.venue;
        template.querySelector('.venue')?.setAttribute('title', d.venue);
        template.querySelector('.region').textContent = d.region;
        template.querySelector('.host').textContent = d.host;
        template.querySelector('.host')?.setAttribute('title', d.host);
        template.querySelector('.info').href = `/activities/${d.id}`;
        template.querySelector('.view').href = `/activities/${d.id}/edit`;

        fragment.appendChild(template);
      }
    });

    tableBody.appendChild(fragment);
  }
}

function updatePagination() {
  currentPageInput.value = currentPage;
  totalPagesSpan.textContent = totalPages;
  currentRecordsSpan.textContent = (currentPage - 1) * recordsPerPage + 1;
  totalRecordsSpan.textContent = totalRecords;
}

async function changePage(newPage) {
  if (newPage >= 1 && newPage <= totalPages) {
    currentPage = newPage;
    await fetchData();
  }
}

document
  .getElementById('firstPage')
  ?.addEventListener('click', async () => await changePage(1));
document
  .getElementById('prevPage')
  ?.addEventListener('click', async () => await changePage(currentPage - 1));
document
  .getElementById('nextPage')
  ?.addEventListener('click', async () => await changePage(currentPage + 1));
document
  .getElementById('lastPage')
  ?.addEventListener('click', async () => await changePage(totalPages));

refreshButton?.addEventListener('click', async () => {
  await fetchData();
});

currentPageInput?.addEventListener('change', async () => {
  const newPage = parseInt(currentPageInput?.value);
  await changePage(newPage);
});

recordsPerPageSelect.addEventListener('change', async () => {
  recordsPerPage = parseInt(recordsPerPageSelect.value);
  currentPage = 1;
  await fetchData();
});

// Initial fetch
fetchData().catch((err) => console.log(err));
