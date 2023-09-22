document.addEventListener('DOMContentLoaded', async function() {
    const appId = 1; // Replace with the actual app ID

    try {
        const response = await fetch(`/api/getRelease?appId=${appId}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer <authToken>`, // Replace with the actual auth token
            },
        });

        if (response.status === 200) {
            const releaseList = await response.json();
            const releaseListContainer = document.getElementById('releaseList');

            releaseList.forEach(release => {
                const releaseElement = document.createElement('div');
                releaseElement.className = 'release-card';
                releaseElement.innerHTML = `
                    <h2>Version: ${release.versionName}</h2>
                    <p>Created at: ${release.createdAt}</p>
                    <p>Updated at: ${release.updatedAt}</p>
                `;
                releaseListContainer.appendChild(releaseElement);
            });
        } else {
            // Handle error cases
            const errorData = await response.json();
            console.error('Failed to fetch release info:', errorData);
            // Display error message to the user
        }
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle network or unexpected errors
    }
});