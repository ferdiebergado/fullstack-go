// @ts-check

import { initializeNavLinks } from './ui.js';
import { setupActivityForms } from './activity.js';
import { setupPagination } from './components/pager.js';

// Initialize navigation links
initializeNavLinks();

// Set up form handling for activities
setupActivityForms();

// Set up pagination handling
setupPagination();
