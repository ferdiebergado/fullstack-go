// @ts-check
'use strict';

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
 * @property {string} jumpPage
 */

import { highlightText, truncateText, sanitize } from '../utils';

const MAX_TEXT_LENGTH = 30;
const TABLE_ROW_HEIGHT = '6rem';
const STATE_KEY = 'tableState';
const ROWS_PER_PAGE = 5;

// Elements retrieval
const table = /** @type {HTMLTableElement | null} */ (
  document.getElementById('dynamicTable')
);
const tableHead = /** @type {HTMLTableCellElement | null} */ (
  document.getElementById('dynamicTableHead')
);
const tableBody = /** @type {HTMLTableCellElement | null} */ (
  document.getElementById('dynamicTableBody')
);
const filterInput = /** @type {HTMLInputElement | null} */ (
  document.getElementById('filterInput')
);
const filterSelect = /** @type {HTMLSelectElement | null} */ (
  document.getElementById('filterSelect')
);
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
const refreshButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('refreshButton')
);
const resetButton = /** @type {HTMLButtonElement | null} */ (
  document.getElementById('resetButton')
);

let data = [];
let currentPage = 1;
let totalPages = 1;
let totalItems = 1;
let rowsPerPage = getRowsPerPage();
let sortColumn = '';
let sortDirection = 1;
let search = '';
let searchColumn = '';

// Fetch headers from data attributes
const headers = JSON.parse(table?.dataset.headers || '[]');
const apiUrl = table?.dataset.url || '';

function getRowsPerPage() {
  return Number(rowsPerPageSelect?.value) || ROWS_PER_PAGE;
}

// Fetch data from API endpoint with pagination
async function fetchData(page = 1) {
  try {
    search = filterInput?.value.toLocaleLowerCase() || '';
    searchColumn = filterSelect?.value || '';
    rowsPerPage = getRowsPerPage();
    currentPage = page;

    const params = new URLSearchParams({
      page: String(page),
      limit: String(rowsPerPage),
      sortCol: sortColumn,
      sortDir: String(sortDirection),
      search: encodeURIComponent(search),
      searchCol: searchColumn,
    });

    const response = await fetch(`${apiUrl}?${params.toString()}`);

    if (!response.ok) throw new Error('Failed to fetch data');

    /** @type {import('../typedefs').ApiResponse} */
    const jsonData = await response.json();
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
  console.log(searchColumn);

  if (filterSelect) {
    filterSelect.innerHTML = headers
      .map(
        /** @param {TableHeader} header */
        (header) =>
          `<option value="${header.field}" ${
            searchColumn === header.field ? 'selected' : ''
          }>${header.label}</option>`
      )
      .join('');
  }
}

function renderTableHead() {
  if (tableHead) {
    const headerFragment = document.createDocumentFragment();
    const headerRow = document.createElement('tr');

    headers.forEach(({ field, label }) => {
      const th = document.createElement('th');
      th.textContent = label;
      th.title = 'Click to sort';
      th.style.cursor = 'pointer';
      th.addEventListener('click', () => sortTable(th, field));
      headerRow.appendChild(th);
    });

    const actionsHeader = document.createElement('th');
    actionsHeader.textContent = 'Actions';
    headerRow.appendChild(actionsHeader);
    headerFragment.appendChild(headerRow);

    tableHead.appendChild(headerFragment);
  }
}

// Render table with optimized DOM manipulation
function renderTableBody() {
  const bodyFragment = document.createDocumentFragment();
  if (totalItems > 0) {
    data.forEach((row) => {
      const tr = document.createElement('tr');
      tr.style.height = TABLE_ROW_HEIGHT;

      headers.forEach(({ field }) => {
        const td = document.createElement('td');
        const fieldValue = sanitize(row[field]);

        if (typeof fieldValue === 'string') {
          td.style.wordBreak = 'break-all';
          td.title = fieldValue;
          const truncatedText = truncateText(fieldValue, MAX_TEXT_LENGTH);
          td.innerHTML =
            field === searchColumn
              ? highlightText(truncatedText, filterInput?.value || '')
              : truncatedText;
        }

        tr.appendChild(td);
      });

      const actionCell = createActionCell(row.id);
      tr.appendChild(actionCell);
      bodyFragment.appendChild(tr);
    });
  } else {
    const noDataRow = document.createElement('tr');
    const noDataCell = document.createElement('td');
    noDataCell.colSpan = headers.length + 1;
    noDataCell.textContent = 'No data.';
    noDataRow.appendChild(noDataCell);
    bodyFragment.appendChild(noDataRow);
  }

  if (tableBody) {
    tableBody.innerHTML = ''; // Clear existing content
    tableBody.appendChild(bodyFragment);
  }
}

function createActionCell(rowId) {
  const prefix = '/api';
  const htmlEndpoint = apiUrl.startsWith('/api')
    ? apiUrl.replace(prefix, '')
    : apiUrl;
  const actionCell = document.createElement('td');
  const infoLink = createLink(
    `Info`,
    `${htmlEndpoint}/${rowId}`,
    'btn btn-small btn-primary'
  );
  const editLink = createLink(
    `Edit`,
    `${htmlEndpoint}/${rowId}/edit`,
    'btn btn-small btn-secondary'
  );

  actionCell.appendChild(infoLink);
  actionCell.appendChild(editLink);
  return actionCell;
}

function createLink(text, href, className) {
  const link = document.createElement('a');
  link.href = href;
  link.textContent = text;
  link.className = className;
  return link;
}

// Render simplified pagination controls
function updatePagination() {
  if (totalItems > 0) {
    setButtonState(firstButton, currentPage === 1);
    setButtonState(prevButton, currentPage === 1);
    setButtonState(nextButton, currentPage === totalPages);
    setButtonState(lastButton, currentPage === totalPages);

    pageJumpInput.disabled = false;
    rowsPerPageSelect.disabled = false;

    pageInfo.textContent = `Page ${currentPage} of ${totalPages}`;
    const startRecord = (currentPage - 1) * rowsPerPage + 1;
    const endRecord = Math.min(currentPage * rowsPerPage, totalItems);
    recordsInfo.textContent = `Record ${startRecord}-${endRecord} of ${totalItems} records`;
  } else {
    setButtonState(firstButton, true);
    setButtonState(prevButton, true);
    setButtonState(nextButton, true);
    setButtonState(lastButton, true);
    pageJumpInput.disabled = true;
    rowsPerPageSelect.disabled = true;
    pageInfo.textContent = '';
    recordsInfo.textContent = '';
  }
}

/**
 *
 * @param {HTMLButtonElement} button
 * @param {boolean} isDisabled
 */
function setButtonState(button, isDisabled) {
  button && (button.disabled = isDisabled);
}

/** @param {number} page */
function changePage(page) {
  if (page >= 1 && page <= totalPages) {
    fetchData(page);
  }
}

function updateRowsPerPage() {
  rowsPerPage = Number(rowsPerPageSelect?.value);
  fetchData(); // Fetch from the first page with new rows per page setting
}

/**
 * @param {HTMLTableCellElement} th
 *  @param {string} field */
function sortTable(th, field) {
  sortDirection = sortColumn === field ? -sortDirection : 1;
  sortColumn = field;
  const sortOrder = th.dataset.sort === 'asc' ? 'desc' : 'asc';

  // Reset other headers' sort attribute
  Array.from(table?.querySelectorAll('th')).forEach((header) => {
    header.dataset.sort = 'none';
  });

  // Set the selected header's sort attribute
  th.dataset.sort = sortOrder;
  fetchData();
}

/** @param {HTMLInputElement} input  */
function jumpToPage(input) {
  const targetPage = Number(input.value);
  if (targetPage >= 1 && targetPage <= totalPages) {
    changePage(targetPage);
  } else {
    changePage(Math.max(targetPage, 1));
    input.value = '1';
  }
}

function saveState() {
  /** @type {TableState} */
  const tableState = {
    sortColumn,
    sortDirection,
    search,
    currentPage,
    rowsPerPage,
    searchCol: searchColumn,
    jumpPage: pageJumpInput?.value || '1',
  };

  console.log('Saved state:', tableState);

  localStorage.setItem(STATE_KEY, JSON.stringify(tableState));
}

function retrieveState() {
  const savedState = localStorage.getItem(STATE_KEY);

  if (savedState) {
    /** @type {TableState} */
    const tableState = JSON.parse(savedState);

    console.log('Retrieved state:', tableState);

    currentPage = tableState.currentPage || 1;
    rowsPerPageSelect.value =
      String(tableState.rowsPerPage) || String(getRowsPerPage());
    sortColumn = tableState.sortColumn || '';
    sortDirection = tableState.sortDirection || 1;
    filterInput.value = tableState.search;
    filterSelect.value = tableState.searchCol;
    searchColumn = tableState.searchCol;
    search = tableState.search;
    pageJumpInput.value = tableState.jumpPage;

    console.log(filterSelect?.value);
  }
}

// Event listeners
filterInput?.addEventListener(
  'keypress',
  (e) => e.key === 'Enter' && fetchData()
);
filterSelect?.addEventListener(
  'change',
  (e) => (searchColumn = e.target?.value)
);
rowsPerPageSelect?.addEventListener('change', updateRowsPerPage);
pageJumpInput?.addEventListener('keypress', (e) => {
  if (e.key === 'Enter') jumpToPage(e.target);
});
firstButton?.addEventListener('click', () => changePage(1));
prevButton?.addEventListener('click', () => changePage(currentPage - 1));
nextButton?.addEventListener('click', () => changePage(currentPage + 1));
lastButton?.addEventListener('click', () => changePage(totalPages));
refreshButton?.addEventListener('click', () => fetchData(currentPage));
resetButton?.addEventListener('click', () => {
  filterInput.value = '';
  filterSelect.options[0].selected = true;
  pageJumpInput.value = '1';
  rowsPerPageSelect.value = ROWS_PER_PAGE;
  fetchData();
});

// Initial render and state retrieval
retrieveState();
renderFilterSelect();
renderTableHead();
fetchData(currentPage);
