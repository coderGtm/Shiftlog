document.addEventListener('DOMContentLoaded', async function() {
    try {
        const response = await fetch('/api/getApps', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer <authToken>`, // Replace with the actual auth token
            },
        });

        if (response.status === 200) {
            const appList = await response.json();
            const appListContainer = document.getElementById('appList');

            appList.forEach(app => {
                const appElement = document.createElement('div');
                appElement.className = 'app-card';
                appElement.innerHTML = `
                    <h2>${app.name}</h2>
                    <p>Created at: ${app.createdAt}</p>
                    <p>Updated at: ${app.updatedAt}</p>
                `;
                appListContainer.appendChild(appElement);
            });
        } else {
            // Handle error cases
            const errorData = await response.json();
            console.error('Failed to fetch app list:', errorData);
            // Display error message to the user
        }
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle network or unexpected errors
    }
});
