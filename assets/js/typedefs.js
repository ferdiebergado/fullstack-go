// @ts-check
'use strict';

/**
 * Represents an API response.
 *
 * @typedef {Object} ApiResponse
 * @property {ResponseMeta} meta - The metadata of the response
 * @property {Object} data - The data of the response
 */

/**
 * Represents metadata of an API response.
 *
 * @typedef {Object} ResponseMeta
 * @property {string} message - The message
 * @property {ValidationError[]} errors - The object that contains the errors
 * @property {PaginationMeta} pagination - The object that contains the pagination metadata
 */

/**
 * Represents a validation error object.
 *
 * @typedef {Object} ValidationError
 * @property {string} field - The field that has the validation error.
 * @property {string} error - The error message associated with the field.
 */

/**
 * Represents metadata of paginated data.
 *
 * @typedef {Object} PaginationMeta
 *
 * @property {number} total_items
 * @property {number} total_pages
 * @property {number} page
 * @property {number} limit
 */

/**
 * Represents a Custom Event Detail.
 *
 * @typedef {Object} CustomEventDetail
 * @property {number} id - The unique identifier.
 * @property {string} name - The name associated with the event.
 */

/**
 * Represents a Custome Event Init.
 *
 * @typedef {CustomEventInit} MyCustomEventInit
 * @property {CustomEventDetail} detail - Custom detail object for the event.
 */

export default {};
