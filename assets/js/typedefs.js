// @ts-check
/**
 * Represents an API response.
 * @typedef {Object} ApiResponse
 * @property {boolean} success - The status of the response
 * @property {string} message - The message
 * @property {ValidationError[]} errors - The object that contains the errors
 * @property {Object} data - The object that contains the data
 */

/**
 * Represents a validation error object.
 * @typedef {Object} ValidationError
 * @property {string} field - The field that has the validation error.
 * @property {string} error - The error message associated with the field.
 */

/**
 * @typedef {Object} CustomEventDetail
 * @property {number} id - The unique identifier.
 * @property {string} name - The name associated with the event.
 */

/**
 * @typedef {CustomEventInit} MyCustomEventInit
 * @property {CustomEventDetail} detail - Custom detail object for the event.
 */

export default {};
