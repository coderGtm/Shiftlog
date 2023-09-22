document.getElementById('createAccountForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('/api/createAccount', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.status === 200) {
            const responseData = await response.json();
            document.getElementById('createAccountMessage').textContent = 'Account created successfully!';
            // Redirect to the login page or perform other actions
        } else {
            const errorData = await response.json();
            document.getElementById('createAccountMessage').textContent = `Error: ${errorData.message}`;
            // Display error message to the user
        }
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle network or unexpected errors
    }
});
