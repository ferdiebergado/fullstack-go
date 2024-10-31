// @ts-check
'use strict';

/**
 * @typedef {Object} TableHeader
 * @property {string} field
 * @property {string} label
 */

const table = /** @type {HTMLTableElement | null} */ (
  document.getElementById('dynamicTable')
);

/** @type {TableHeader[]} */
const headers = JSON.parse(table?.getAttribute('data-headers') || '[]');

const apiUrl = table?.getAttribute('data-url') || '';

const filterInput = /** @type {HTMLInputElement | null} */ (
  document.getElementById('filterInput')
);
const paginationControls = document.getElementById('paginationControls');

const rowsPerPageSelect = /** @type {HTMLSelectElement | null} */ (
  document.getElementById('rowsPerPage')
);

const pageJumpInput = /** @type {HTMLInputElement | null} */ (
  document.getElementById('pageJumpInput')
);
const paginationInfo = document.getElementById('paginationInfo');

const firstButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('firstButton')
);
const prevButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('prevButton')
);
const nextButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('nextButton')
);
const lastButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('lastButton')
);

let data = [];
let currentPage = 1;
let totalPages = 1;
let totalItems = 0;
let rowsPerPage = parseInt(rowsPerPageSelect?.value || '10', 10);
let sortColumn = null;
let sortDirection = 1;

// Fetch data from API endpoint with pagination
async function fetchData(page = 1) {
  try {
    const params = new URLSearchParams({
      page: String(page),
      limit: String(rowsPerPage),
      sortCol: sortColumn,
      sortDir: String(sortDirection),
      search: encodeURIComponent(filterInput?.value.toLocaleLowerCase() || ''),
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

// Render table with optimized DOM manipulation
function renderTable() {
  // Create table headers only once
  if (!table?.tHead) {
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    headers.forEach(
      /**
       * @param {TableHeader} header
       */
      (header) => {
        const th = document.createElement('th');
        th.textContent = header.label;
        th.addEventListener('click', () => sortTable(header.field));
        headerRow.appendChild(th);
      }
    );

    thead.appendChild(headerRow);
    table?.appendChild(thead);
  }

  // Create table body in a document fragment for optimized rendering
  const tbody = document.createElement('tbody');
  if (totalItems > 0) {
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

      replaceTableBody();
    });
  } else {
    const tr = document.createElement('tr');
    const td = document.createElement('td');

    td.colSpan = headers.length;
    td.textContent = 'No data.';

    tr.appendChild(td);

    replaceTableBody();

    tbody.appendChild(tr);
  }

  table?.appendChild(tbody);
}

// Replace existing tbody to minimize reflows
function replaceTableBody() {
  table?.tBodies[0] && table.removeChild(table.tBodies[0]);
}

// Render simplified pagination controls
function updatePagination() {
  if (paginationControls) {
    // First Page Button
    firstButton && (firstButton.disabled = currentPage === 1);

    // Previous Page Button
    prevButton && (prevButton.disabled = currentPage === 1);

    // Next Page Button
    nextButton && (nextButton.disabled = currentPage === totalPages);

    // Last Page Button
    lastButton && (lastButton.disabled = currentPage === totalPages);

    const startRecord = (currentPage - 1) * rowsPerPage + 1;
    const endRecord = Math.min(currentPage * rowsPerPage, totalItems);

    if (paginationInfo) {
      if (totalItems > 0)
        return (paginationInfo.textContent = `Page ${currentPage} of ${totalPages} - Showing ${startRecord}-${endRecord} of ${totalItems} records`);

      return (paginationInfo.textContent = '');
    }
  }
}

/** @param {number} page */
function changePage(page) {
  fetchData(page);
}

function updateRowsPerPage() {
  rowsPerPage = parseInt(rowsPerPageSelect?.value || '10', 10);

  fetchData(); // Fetch from the first page with new rows per page setting
}

/** @param {string} field */
function sortTable(field) {
  if (sortColumn === field) {
    sortDirection *= -1;
  } else {
    sortDirection = 1;
  }

  sortColumn = field;

  fetchData();
}

// Sanitize data to prevent XSS attacks
/** @param {string} text */
function sanitize(text) {
  const element = document.createElement('div');

  element.textContent = text;

  return element.innerHTML;
}

/**
 *
 * @param {Function} func
 * @param {number} delay
 * @returns
 */
function debounce(func, delay) {
  let timeout;
  return function (...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, args), delay);
  };
}

// Attach event listener for filtering
filterInput?.addEventListener('input', debounce(fetchData, 300));

rowsPerPageSelect?.addEventListener('change', updateRowsPerPage);

pageJumpInput?.addEventListener('change', () => {
  const targetPage = parseInt(pageJumpInput.value, 10);

  if (targetPage >= 1 && targetPage <= totalPages) {
    changePage(targetPage);
  } else {
    if (pageJumpInput) pageJumpInput.value = String(currentPage);
  }
});

if (pageJumpInput) {
  pageJumpInput.max = String(totalPages);
  pageJumpInput.value = String(currentPage);
}

firstButton?.addEventListener('click', () => changePage(1));
prevButton?.addEventListener('click', () => changePage(currentPage - 1));
nextButton?.addEventListener('click', () => changePage(currentPage + 1));
lastButton?.addEventListener('click', () => changePage(totalPages));

// Initial fetch and render
fetchData();
