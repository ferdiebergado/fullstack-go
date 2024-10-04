function formHandler() {
	return {
		formData: {
			title: document.getElementById('title').value,
			start_date: document.getElementById('start_date').value,
			end_date: document.getElementById('end_date').value,
			venue: document.getElementById('venue').value,
			host: document.getElementById('host').value,
			metadata: {}
		},
		errors: {},
		notification: {
			message: '',
			type: '',
		},
		async submitForm(e) {
			e.preventDefault();

			const form = e.target;
			const url = form.action;  // Fetch the URL from the form's action attribute

			this.clearErrors();
			this.clearNotification();

			try {
				const response = await fetch(url, {
					method: 'PUT',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify(this.formData)
				});

				const result = await response.json();

				if (!response.ok) {
					this.errors = result.errors;
				} else {
					this.notification = { message: 'Form submitted successfully!', type: 'success' };
					this.formData = {
						title: '',
						start_date: '',
						end_date: '',
						venue: '',
						host: '',
						metadata: {},
					}
				}
			} catch (error) {
				this.notification = { message: 'An error occurred. Please try again later.', type: 'error' };
			}
		},
		clearErrors() {
			this.errors = {};
		},
		clearNotification() {
			this.notification = { message: '', type: '' };
		}
	};
}
