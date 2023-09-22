document.getElementById('loginForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.status === 200) {
            const responseData = await response.json();
            document.getElementById('loginMessage').textContent = 'Login successful!';
            // Redirect to the dashboard or perform other actions
        } else {
            const errorData = await response.json();
            document.getElementById('loginMessage').textContent = `Error: ${errorData.message}`;
            // Display error message to the user
        }
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle network or unexpected errors
    }
});