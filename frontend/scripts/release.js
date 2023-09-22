document.addEventListener('DOMContentLoaded', async function() {
    const releaseId = 1; // Replace with the actual release ID

    try {
        const response = await fetch(`/api/getReleaseNotes?releaseId=${releaseId}`, {
            method: 'GET',
        });

        if (response.status === 200) {
            const releaseDetails = await response.json();
            const releaseDetailsContainer = document.getElementById('releaseDetails');

            releaseDetailsContainer.innerHTML = `
                <h2>Version: ${releaseDetails.versionName}</h2>
                <p>Release Notes:</p>
                <div>${releaseDetails.notesHtml}</div>
                <p>Last Updated at: ${releaseDetails.updatedAt}</p>
            `;
        } else {
            // Handle error cases
            const errorData = await response.json();
            console.error('Failed to fetch release notes:', errorData);
            // Display error message to the user
        }
    } catch (error) {
        console.error('An error occurred:', error);
        // Handle network or unexpected errors
    }
});
