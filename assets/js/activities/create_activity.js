function createActivityHandler() {
	return {
		// Reactive form data
		formData: {
			title: '',
			start_date: '',
			end_date: '',
			venue: '',
			host: '',
			metadata: {}
		},
		errors: {}, // Object to track validation errors
		notification: {
			message: '',
			type: '',
		},
		loading: false, // To track the loading state for the spinner

		// Form submission handler
		submitForm() {
			console.log("form submitted.");

			// Show spinner by setting loading to true
			this.loading = true;

			// Reset previous errors
			this.clearErrors();

			this.clearNotification();

			// Validate form inputs
			if (!this.formData.title) {
				this.errors.title = 'Title is required.';
			}

			if (!this.formData.start_date) {
				this.errors.start_date = 'Start Date is required.';
			}

			if (!this.formData.end_date) {
				this.errors.end_date = 'End Date is required.';
			}

			// If there are errors, stop the submission
			if (Object.keys(this.errors).length > 0) {
				this.loading = false
				return;
			}

			// Send the data to the API using fetch if valid
			fetch('/api/activities', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(this.formData),
			})
				.then(response => {
					if (!response.ok) {
						this.errors = result.errors;
						this.notification = {
							message: 'Form submission failed!',
							type: 'error'
						}
						throw new Error('Network response was not ok');
					}

					this.notification = { message: 'Form submitted successfully!', type: 'success' };

					return response.json();
				})
				.then(data => {
					console.log(data);

					// Handle success

					// Clear form data
					this.formData = {
						title: '',
						start_date: '',
						end_date: '',
						venue: '',
						host: '',
						metadata: {},
					}

					this.loading = false; // To track the loading state for the spinner
				})
				.catch(error => {
					console.log(error);

					// Handle error
					this.notification = {
						message: 'Form submission failed!',
						type: 'error'
					}

					this.loading = false; // To track the loading state for the spinner
				});
		},

		clearErrors() {
			this.errors = {};
		},

		clearNotification() {
			this.notification = { message: '', type: '' };
		},
	}
}
