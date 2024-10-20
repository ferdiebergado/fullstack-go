// @ts-check
const table = document.getElementById('data-table');
const tableHead = table?.querySelector('thead');
const tableBody = table?.querySelector('tbody');
const columnsAttr = table?.getAttribute('data-columns');
const apiUrl = table?.getAttribute('data-url');
const firstBtn = document.getElementById('firstPage');
const prevBtn = document.getElementById('prevPage');
const nextBtn = document.getElementById('nextPage');
const lastBtn = document.getElementById('lastPage');
const pagination = document.getElementById('pagination');
const pageInput = document.getElementById('page');
const pageInfo = document.getElementById('pageInfo');
const recordsPerPageSelect = document.getElementById('recordsPerPage');
const totalPagesSpan = document.getElementById('totalPages');

/** @type {HTMLSelectElement} */
const filterSelect = document.getElementById('filter');

let currentPage = 1;
let rowsPerPage = 10;
let totalPages = 1;
/** @type {string[]} */
let columns = [];
let allData = [];
let sortDirection = 1;
let sortColumn = '';

function initializePager() {
  // First and Previous buttons
  firstBtn?.addEventListener('click', () => {
    currentPage = 1;
    fetchAndDisplayData();
  });

  prevBtn?.addEventListener('click', () => {
    currentPage--;
    fetchAndDisplayData();
  });

  // Next and Last buttons
  nextBtn?.addEventListener('click', () => {
    currentPage++;
    fetchAndDisplayData();
  });

  lastBtn?.addEventListener('click', () => {
    currentPage = totalPages;
    fetchAndDisplayData();
  });
}

function setupPageInput() {
  // Jump to page input
  pageInput?.addEventListener('change', (e) => {
    const page = Number(e.target?.value);
    if (page >= 1 && page <= totalPages) {
      currentPage = page;
      fetchAndDisplayData();
    }
  });
}

function initTable() {
  columns = columnsAttr?.split(',') || [];

  document
    .getElementById('search')
    ?.addEventListener('input', debounce(fetchAndDisplayData, 300));

  recordsPerPageSelect?.addEventListener('change', async () => {
    rowsPerPage = parseInt(recordsPerPageSelect.value);
    currentPage = 1;
    fetchAndDisplayData();
  });

  const optionsFragment = document.createDocumentFragment();

  columns.forEach((col) => {
    const filterOpt = document.createElement('option');
    filterOpt.value = col;
    filterOpt.textContent = col;
    optionsFragment.appendChild(filterOpt);
  });

  filterSelect.innerHTML = '';
  filterSelect.appendChild(optionsFragment);

  initializePager();
  setupPageInput();
  fetchAndDisplayData();
}

function fetchAndDisplayData() {
  const searchQuery = document
    .getElementById('search')
    ?.value.trim()
    .toLowerCase();

  const params = new URLSearchParams({
    page: String(currentPage),
    limit: String(rowsPerPage),
    sort: sortColumn,
    sortDir: String(sortDirection),
    search: searchQuery,
  });

  fetch(`${apiUrl}?${params.toString()}`)
    .then((response) => response.json())
    .then((payload) => {
      const { total_items, total_pages, page, data } = payload.data;

      allData = data;
      totalPages = total_pages;
      currentPage = page;
      createTable(columns, allData);
      updatePaginationInfo(total_items, currentPage, totalPages);
    })
    .catch((error) => console.error('Error fetching data:', error));
}

function createTable(columns, data) {
  if (!tableHead?.children.length) {
    tableHead.innerHTML = '';
    const headerRow = document.createElement('tr');
    columns.forEach((col) => {
      const th = document.createElement('th');
      th.textContent = col.toUpperCase();
      th.addEventListener('click', () => sortTable(col));
      headerRow.appendChild(th);
    });
    tableHead?.appendChild(headerRow);
  }

  const bodyFragment = document.createDocumentFragment();
  data.forEach((row) => {
    const tr = document.createElement('tr');
    columns.forEach(
      /** @param {string} col */ (col) => {
        const td = document.createElement('td');
        td.textContent = row[col.trim()];

        tr.appendChild(td);
      }
    );
    bodyFragment.appendChild(tr);
  });

  tableBody.innerHTML = '';
  tableBody?.appendChild(bodyFragment);
}

function updatePaginationInfo(totalItems, currentPage, totalPages) {
  // Update page and record info
  const startRecord = (currentPage - 1) * rowsPerPage + 1;
  const endRecord = Math.min(currentPage * rowsPerPage, totalItems);
  firstBtn.disabled = currentPage === 1;
  prevBtn.disabled = currentPage === 1;
  nextBtn.disabled = currentPage === totalPages;
  lastBtn.disabled = currentPage === totalPages;
  pageInput.max = totalPages;
  pageInput.value = currentPage;
  totalPagesSpan.textContent = totalPages;
  pageInfo.textContent = `Showing ${startRecord} to ${endRecord} of ${totalItems} records`;
}

function sortTable(column) {
  if (sortColumn === column) {
    sortDirection = -sortDirection; // Toggle sort direction
  } else {
    sortDirection = 1; // Default to ascending
  }
  sortColumn = column;

  fetchAndDisplayData();
}

function debounce(func, delay) {
  let timeout;
  return function (...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, args), delay);
  };
}

initTable();
