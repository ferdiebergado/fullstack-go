// @ts-check
'use strict';

/** @type {HTMLTableElement} */
const table = document.getElementById('dynamicTable');
const headers = JSON.parse(table?.getAttribute('data-headers'));
const apiUrl = table?.getAttribute('data-url');

/** @type {HTMLInputElement} */
const filterInput = document.getElementById('filterInput');
const paginationControls = document.getElementById('paginationControls');

/** @type {HTMLSelectElement} */
const rowsPerPageSelect = document.getElementById('rowsPerPage');

const pageJumpInput = document.getElementById('pageJumpInput');
const paginationInfo = document.getElementById('paginationInfo');

const firstButton = document.getElementById('firstButton');
const prevButton = document.getElementById('prevButton');
const nextButton = document.getElementById('nextButton');
const lastButton = document.getElementById('lastButton');

let data = [];
let currentPage = 1;
let totalPages = 1;
let totalItems = 0;
let rowsPerPage = parseInt(rowsPerPageSelect.value, 10);
let sortColumn = null;
let sortDirection = 1;

// Fetch data from API endpoint with pagination
async function fetchData(page = 1, query = '') {
  try {
    const params = new URLSearchParams({
      page: String(page),
      limit: String(rowsPerPage),
      sortCol: sortColumn,
      sortDir: String(sortDirection),
      search: encodeURIComponent(query),
    });

    const response = await fetch(`${apiUrl}?${params.toString()}`);

    if (!response.ok) throw new Error('Failed to fetch data');

    const jsonData = await response.json();

    // Process the response with pagination
    data = jsonData.data || [];
    const pagination = jsonData.meta.pagination;
    currentPage = pagination.page;
    totalPages = pagination.total_pages;
    totalItems = pagination.total_items;

    renderTable();
    updatePagination();
  } catch (error) {
    console.error('Fetch error:', error);
  }
}

/**
 * @typedef {Object} TableHeader
 * @property {string} field
 * @property {string} label
 */

// Render table with optimized DOM manipulation
function renderTable() {
  // Create table headers only once
  if (!table?.tHead) {
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    headers.forEach(
      /**
       * @param {TableHeader} header
       * @param {number} index
       */
      (header, index) => {
        const th = document.createElement('th');
        th.textContent = header.label;
        th.addEventListener('click', () => sortTable(header.field));
        headerRow.appendChild(th);
      }
    );
    thead.appendChild(headerRow);
    table.appendChild(thead);
  }

  // Create table body in a document fragment for optimized rendering
  const tbody = document.createElement('tbody');
  data.forEach((row) => {
    const tr = document.createElement('tr');
    headers.forEach(
      /** @param {TableHeader} header */ (header) => {
        const td = document.createElement('td');
        td.textContent = sanitize(row[header.field]);
        tr.appendChild(td);
      }
    );
    tbody.appendChild(tr);
  });

  // Replace existing tbody to minimize reflows
  if (table.tBodies[0]) {
    table.removeChild(table.tBodies[0]);
  }
  table.appendChild(tbody);
}

// Render simplified pagination controls
function updatePagination() {
  if (paginationControls) {
    // First Page Button
    firstButton.disabled = currentPage === 1;

    // Previous Page Button
    prevButton.disabled = currentPage === 1;

    // Next Page Button
    nextButton.disabled = currentPage === totalPages;

    // Last Page Button
    lastButton.disabled = currentPage === totalPages;

    const startRecord = (currentPage - 1) * rowsPerPage + 1;
    const endRecord = Math.min(currentPage * rowsPerPage, totalItems);
    paginationInfo.textContent = `Page ${currentPage} of ${totalPages} - Showing ${startRecord}-${endRecord} of ${totalItems} records`;
  }
}

/** @param {number} page */
function changePage(page) {
  fetchData(page, filterInput?.value);
}

function updateRowsPerPage() {
  rowsPerPage = parseInt(rowsPerPageSelect.value, 10);
  fetchData(1, filterInput.value); // Fetch from the first page with new rows per page setting
}

/** @param {string} field */
function sortTable(field) {
  if (sortColumn === field) {
    sortDirection *= -1;
  } else {
    sortDirection = 1;
  }
  sortColumn = field;
  // data.sort((a, b) => {
  //   const header = headers[columnIndex].toLowerCase();
  //   return (a[header] > b[header] ? 1 : -1) * sortDirection;
  // });
  fetchData();
}

function filterTable() {
  const query = filterInput.value.toLowerCase();
  fetchData(1, query); // Start from the first page when filtering
}

// Sanitize data to prevent XSS attacks
/** @param {string} text */
function sanitize(text) {
  const element = document.createElement('div');
  element.textContent = text;
  return element.innerHTML;
}

// Attach event listener for filtering
filterInput.addEventListener('input', filterTable);
rowsPerPageSelect.addEventListener('change', updateRowsPerPage);
pageJumpInput.addEventListener('change', () => {
  const targetPage = parseInt(pageJumpInput.value, 10);
  if (targetPage >= 1 && targetPage <= totalPages) {
    changePage(targetPage);
  } else {
    pageJumpInput.value = currentPage;
  }
});

pageJumpInput.max = totalPages;
pageJumpInput.value = currentPage;
firstButton.addEventListener('click', () => changePage(1));
prevButton.addEventListener('click', () => changePage(currentPage - 1));
nextButton.addEventListener('click', () => changePage(currentPage + 1));
lastButton.addEventListener('click', () => changePage(totalPages));

// Initial fetch and render
fetchData();
