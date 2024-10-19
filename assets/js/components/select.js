// @ts-check
/**
 * Updates the select element with new options.
 * @param {HTMLSelectElement} select
 * @param {Object} detail
 */
export function updateSelect(select, detail) {
  const options = Array.from(select.options).slice(1);

  const newOption = document.createElement('option');
  newOption.value = detail.id.toString();
  newOption.text = detail.name;
  newOption.selected = true;
  options.push(newOption);
  options.sort((a, b) => a.text.localeCompare(b.text));

  const firstOption = select.options[0];
  select.innerHTML = '';
  select.add(firstOption);
  options.forEach((option) => select.add(option));
}
