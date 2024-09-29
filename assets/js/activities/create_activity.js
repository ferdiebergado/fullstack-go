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
        showNotification: false,
        notificationMessage: '',
        loading: false, // To track the loading state for the spinner

        // Form submission handler
        submitForm() {
            console.log("form submitted.");

            // Check if already loading to prevent double submission
            if (this.loading) {
                console.log("Form submission already in progress.");
                return; // Prevent further submissions
            }

            // Show spinner by setting loading to true
            this.loading = true;

            // Reset previous errors
            this.errors = {};

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
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    console.log(data);

                    // Handle success
                    this.notificationMessage = 'Form submitted successfully!';
                    this.showNotification = true;

                    // Clear form data
                    this.formData.title = '';
                    this.formData.start_date = '';
                    this.formData.end_date = '';
                    this.formData.venue = '';
                    this.formData.host = '';
                    this.formData.metadata = {};

                    this.loading = false; // To track the loading state for the spinner
                })
                .catch(error => {
                    console.log(error);

                    // Handle error
                    this.notificationMessage = 'Form submission failed!';
                    this.showNotification = true;

                    this.loading = false; // To track the loading state for the spinner
                });
        },

        // Email validation function
        validateEmail(email) {
            const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            return re.test(email);
        }
    }
}