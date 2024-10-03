const notification = document.getElementById("notification")
const notificationHeader = document.getElementById("notification-header")
const notificationMessage = document.getElementById("notification-message")

document.getElementById('activity_form').addEventListener('submit', async function (event) {
	event.preventDefault(); // Prevent default form submission

	const formData = new FormData(this);
	const jsonData = Object.fromEntries(formData); // Convert form data to JSON

	console.log(jsonData);

	try {
		const method = this.querySelector('input[name="_method"]').value.toUpperCase() || 'POST'; // Use PUT if specified, else POST
		const response = await fetch(this.action, {
			method,
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(jsonData) // Send data as JSON
		});

		if (!response.ok) {
			throw new Error('Network response was not ok ' + response.statusText);
		}

		// const result = await response.json();
		// console.log('Success:', result); // Handle success (e.g., display a message)
		notification.style.display = "block";
		notificationHeader.textContent = "Success!";
		notificationMessage.textContent = "Form submitted successfully.";
	} catch (error) {
		console.error('Error:', error); // Handle error (e.g., display an error message)
		notification.style.display = "block";
		notificationHeader.textContent = "Error";
		notificationMessage.textContent = "Failed submitting the form.";
	}
});