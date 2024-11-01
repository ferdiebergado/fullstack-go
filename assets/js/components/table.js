// @ts-check
'use strict';

import { sanitize } from '../sanitize';

/**
 * @typedef {Object} TableHeader
 * @property {string} field
 * @property {string} label
 */

/**
 * @typedef {Object} TableState
 * @property {number} currentPage
 * @property {number} rowsPerPage
 * @property {string} sortColumn
 * @property {number} sortDirection
 * @property {string} search
 * @property {string} searchCol
 */

const MAX_TEXT_LENGTH = 30;
const TABLE_ROW_HEIGHT = '6rem';

const table = /** @type {HTMLTableElement | null} */ (
  document.getElementById('dynamicTable')
);

const tableHead = /** @type {HTMLTableCellElement | null} */ (
  document.getElementById('dynamicTableHead')
);

const tableBody = /** @type {HTMLTableCellElement | null} */ (
  document.getElementById('dynamicTableBody')
);

/** @type {TableHeader[]} */
const headers = JSON.parse(table?.getAttribute('data-headers') || '[]');

const apiUrl = table?.getAttribute('data-url') || '';

const filterInput = /** @type {HTMLInputElement | null} */ (
  document.getElementById('filterInput')
);

const filterSelect = /** @type {HTMLSelectElement | null} */ (
  document.getElementById('filterSelect')
);

const paginationControls = document.getElementById('paginationControls');

const rowsPerPageSelect = /** @type {HTMLSelectElement | null} */ (
  document.getElementById('rowsPerPage')
);

const pageJumpInput = /** @type {HTMLInputElement | null} */ (
  document.getElementById('pageJumpInput')
);

const pageInfo = document.getElementById('pageInfo');
const recordsInfo = document.getElementById('recordsInfo');

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
let search = '';
let searchCol = '';

// Fetch data from API endpoint with pagination
async function fetchData(page = 1) {
  try {
    search = filterInput?.value.toLocaleLowerCase() || '';

    const params = new URLSearchParams({
      page: String(page),
      limit: String(rowsPerPage),
      sortCol: sortColumn,
      sortDir: String(sortDirection),
      search: encodeURIComponent(search),
      searchCol: filterSelect?.value || '',
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

    renderTableBody();
    updatePagination();
    saveState();
  } catch (error) {
    console.error('Fetch error:', error);
  }
}

function renderFilterSelect() {
  const optionFragment = document.createDocumentFragment();

  headers.forEach((header) => {
    const option = document.createElement('option');

    option.value = header.field;
    option.textContent = header.label;
    optionFragment.appendChild(option);
  });

  if (filterSelect) {
    filterSelect.innerHTML = '';
    filterSelect.appendChild(optionFragment);
  }
}

function renderTableHead() {
  if (tableHead) {
    const headerRow = document.createElement('tr');

    headers.forEach(
      /**
       * @param {TableHeader} header
       */
      (header) => {
        const th = document.createElement('th');
        th.textContent = header.label;
        th.title = 'Click to sort';
        th.style.cursor = 'pointer';
        th.addEventListener('click', () => sortTable(header.field));
        headerRow.appendChild(th);
      }
    );

    const th = document.createElement('th');
    th.textContent = 'Actions';
    headerRow.appendChild(th);

    tableHead.appendChild(headerRow);
  }
}

// Render table with optimized DOM manipulation
function renderTableBody() {
  // Create table body in a document fragment for optimized rendering
  const bodyFragment = document.createDocumentFragment();

  if (totalItems > 0) {
    data.forEach((row) => {
      const tr = document.createElement('tr');
      tr.style.height = TABLE_ROW_HEIGHT;

      headers.forEach(
        /** @param {TableHeader} header */ (header) => {
          const td = document.createElement('td');
          const fieldValue = sanitize(row[header.field]);
          if (typeof fieldValue === 'string') {
            td.style.wordBreak = 'break-all';
            td.title = fieldValue;
            td.textContent = truncateText(fieldValue, MAX_TEXT_LENGTH);
          }
          tr.appendChild(td);
        }
      );

      const prefix = '/api';

      const htmlEndpoint = apiUrl.startsWith('/api')
        ? apiUrl.replace(prefix, '')
        : apiUrl;

      const actionCell = document.createElement('td');
      const infoLink = document.createElement('a');
      infoLink.href = `${htmlEndpoint}/${row.id}`;
      infoLink.textContent = 'Info';
      infoLink.classList.add('btn', 'btn-small', 'btn-primary');
      actionCell.appendChild(infoLink);

      const editLink = document.createElement('a');
      editLink.href = `${htmlEndpoint}/${row.id}/edit`;
      editLink.text = 'Edit';
      editLink.classList.add('btn', 'btn-small', 'btn-secondary');
      actionCell.appendChild(editLink);

      tr.appendChild(actionCell);

      bodyFragment.appendChild(tr);
    });
  } else {
    const tr = document.createElement('tr');
    const td = document.createElement('td');

    td.colSpan = headers.length + 1;
    td.textContent = 'No data.';

    tr.appendChild(td);

    bodyFragment.appendChild(tr);
  }

  if (tableBody) {
    tableBody.innerHTML = '';
    tableBody.appendChild(bodyFragment);
  }
}

// Render simplified pagination controls
function updatePagination() {
  if (totalItems > 0) {
    if (paginationControls) {
      // First Page Button
      firstButton && (firstButton.disabled = currentPage === 1);

      // Previous Page Button
      prevButton && (prevButton.disabled = currentPage === 1);

      // Next Page Button
      nextButton && (nextButton.disabled = currentPage === totalPages);

      // Last Page Button
      lastButton && (lastButton.disabled = currentPage === totalPages);
    }
    pageJumpInput && (pageJumpInput.disabled = false);
    rowsPerPageSelect && (rowsPerPageSelect.disabled = false);

    pageInfo && (pageInfo.textContent = `Page ${currentPage} of ${totalPages}`);

    const startRecord = (currentPage - 1) * rowsPerPage + 1;
    const endRecord = Math.min(currentPage * rowsPerPage, totalItems);
    recordsInfo &&
      (recordsInfo.textContent = `Record ${startRecord}-${endRecord} of ${totalItems} records`);
  } else {
    firstButton && (firstButton.disabled = true);
    prevButton && (prevButton.disabled = true);
    nextButton && (nextButton.disabled = true);
    lastButton && (lastButton.disabled = true);
    pageJumpInput && (pageJumpInput.disabled = true);
    rowsPerPageSelect && (rowsPerPageSelect.disabled = true);
    pageInfo && (pageInfo.textContent = '');
    recordsInfo && (recordsInfo.textContent = '');
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

function saveState() {
  /** @type {TableState} */
  const tableState = {
    sortColumn,
    sortDirection,
    search,
    currentPage,
    rowsPerPage,
    searchCol,
  };

  localStorage.setItem('tableState', JSON.stringify(tableState));
}

function retrieveState() {
  const savedState = localStorage.getItem('tableState');

  if (savedState) {
    /** @type {TableState} */
    const tableState = JSON.parse(savedState);

    if (tableState) {
      currentPage = tableState.currentPage;
      rowsPerPage = tableState.rowsPerPage;
      sortColumn = tableState.sortColumn;
      sortDirection = tableState.sortDirection;
      search = tableState.search;
      searchCol = tableState.searchCol;

      filterInput && (filterInput.value = search);
    }
  }
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

/** @param {HTMLInputElement} input  */
function jumpToPage(input) {
  const targetPage = parseInt(input.value, 10);

  if (
    targetPage >= 1 &&
    targetPage <= totalPages &&
    targetPage !== currentPage
  ) {
    changePage(targetPage);
  } else if (targetPage < 1) {
    changePage(1);
    input.value = String(1);
  }
}

/**
 * Creates a debounced function that delays invoking the provided function
 * until after a specified delay in milliseconds has elapsed since the last
 * time the debounced function was invoked.
 *
 * @param {Function} func - The function to debounce.
 * @param {number} delay - The number of milliseconds to delay.
 * @returns {...*} A new debounced function that takes the same parameters as `func`.
 */
function debounce(func, delay) {
  let timeout; // Holds the timeout ID

  /**
   * The inner function that will be called after the delay.
   *
   * @param {...*} args - The arguments to pass to the original function.
   */
  return function (...args) {
    // The returned function accepts any arguments
    clearTimeout(timeout); // Clear the previous timeout
    timeout = setTimeout(() => func.apply(this, args), delay); // Set a new timeout
  };
}

/**
 *
 * @param {string} originalText
 * @param {number} maxLength
 *
 * @returns {string}
 */
function truncateText(originalText, maxLength) {
  if (originalText.length > maxLength) {
    return originalText.substring(0, maxLength) + '...';
  }

  return originalText;
}

// Attach event listener for filtering
// filterInput?.addEventListener('input', debounce(fetchData, 300));
filterInput?.addEventListener('keydown', function (event) {
  if (event.key === 'Enter') fetchData();
});

rowsPerPageSelect?.addEventListener('change', updateRowsPerPage);

pageJumpInput?.addEventListener('keydown', function (event) {
  if (event.key === 'Enter') jumpToPage(this);
});

pageJumpInput && (pageJumpInput.value = String(currentPage));

firstButton?.addEventListener('click', () => changePage(1));
prevButton?.addEventListener('click', () => changePage(currentPage - 1));
nextButton?.addEventListener('click', () => changePage(currentPage + 1));
lastButton?.addEventListener('click', () => changePage(totalPages));

// Initial fetch and render
retrieveState();
renderFilterSelect();
renderTableHead();
fetchData(currentPage);
