// @ts-check
/**
 * Updates the select options with new data.
 *
 * @param {HTMLSelectElement} select
 * @param {import("../typedefs").CustomEventDetail} detail
 */
function updateSelect(select, detail) {
  // Step 1: Get all options excluding the first option (with an empty value)
  let options = Array.from(select.options).slice(1); // Skip the first option

  // Step 2: Create a new option and add it to the array
  const newOption = document.createElement('option');
  newOption.value = detail.id.toString();
  newOption.text = detail.name;
  newOption.selected = true;
  options.push(newOption);

  // Step 3: Sort the options by their displayed text
  options.sort((a, b) => a.text.localeCompare(b.text));

  // Step 4: Remove all options except the first one (which has an empty value)
  const firstOption = select.options[0]; // Retain the first option
  const secondOption = select.options[1]; // Retain the first option
  select.innerHTML = ''; // Clear all options
  select.add(firstOption); // Add back the first option
  select.add(secondOption);

  // Step 5: Append the sorted options after the first option
  options.forEach((option) => select.add(option));
}

export default updateSelect;
