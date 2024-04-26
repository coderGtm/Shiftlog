document.body.onload = function () {
    // check if user is logged in
    if (!localStorage.getItem("authToken")) {
        // redirect to login page
        window.location.href = "/login";
    }
    // populate versionName
    document.getElementById("versionName").innerHTML = "<strong>"+getQueryString("versionName")+"</strong>";
    // fetch release data
    fetchReleaseData();
}

function fetchReleaseData() {
    // Fetch release data using an API call
    var releaseId = getReleaseIdFromQuery(); // Implement the function to get release ID from the URL query
    var authToken = localStorage.getItem("authToken");

    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getReleaseNotes?releaseId=" + releaseId, true);
    xhr.setRequestHeader("Authorization", "Bearer " + authToken);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var releaseData = JSON.parse(xhr.responseText);
                updateReleasePage(releaseData);
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error fetching release data! " + xhr.responseText);
            }
        }
    };
    xhr.send();
}

function updateReleasePage(releaseData) {
    // Update page elements with fetched release data
    document.getElementById("version").value = releaseData.name;
    document.getElementById("versionCode").value = releaseData.versionCode;
    document.getElementById("versionName").value = releaseData.versionName;
    document.getElementById("hidden").checked = releaseData.hidden;
    document.getElementById("releaseNotesText").textContent = releaseData.releaseNotesText;
    document.getElementById("releaseNotesMarkdown").textContent = releaseData.releaseNotesMarkdown;
    document.getElementById("releaseNotesHTML").innerHTML = releaseData.releaseNotesHTML;
    document.getElementById("lastModified").textContent = releaseData.lastModified;
    document.getElementById("appName").textContent = releaseData.appName;
}

function updateRelease() {
    var releaseId = getReleaseIdFromQuery(); // Implement the function to get release ID from the URL query
    var version = document.getElementById("version").value;
    var versionCode = document.getElementById("versionCode").value;
    var versionName = document.getElementById("versionName").value;
    var hidden = document.getElementById("hidden").checked;
    var releaseNotes = document.getElementById("releaseNotes").value;

    // Prepare data for update
    var updateData = {
        version: version,
        versionCode: versionCode,
        versionName: versionName,
        hidden: hidden,
        releaseNotes: releaseNotes
    };

    var authToken = localStorage.getItem("authToken");

    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updateRelease?id=" + releaseId, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.setRequestHeader("Authorization", "Bearer " + authToken);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // Release updated successfully, you can redirect or show a success message
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error updating release! " + xhr.responseText);
            }
        }
    };
    xhr.send(JSON.stringify(updateData));
}

function deleteRelease() {
    var releaseId = getReleaseIdFromQuery(); // Implement the function to get release ID from the URL query
    var authToken = localStorage.getItem("authToken");

    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/api/deleteRelease?releaseId=" + encodeURIComponent(releaseId), true);
    xhr.setRequestHeader("Authorization", "Bearer " + authToken);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // Release deleted successfully, you can redirect or show a success message
                window.location.href = "/dashboard";
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error deleting release! " + xhr.responseText);
            }
        }
    };
    xhr.send();
}

function getReleaseIdFromQuery() {
    // Implement the function to extract release ID from the URL query
    // Example: If the URL is /release?id=123, this function should return 123
    // You can use JavaScript's URLSearchParams or other methods to achieve this
    var queryParams = new URLSearchParams(window.location.search);
    return queryParams.get("id");
}

function logout() {
    // remove authToken and username from localStorage and redirect to login page
    localStorage.removeItem("authToken");
    localStorage.removeItem("username");
    window.location.href = "/login";
}

// Get query string from URL
function getQueryString(field) {
    var href = window.location.href;
    var reg = new RegExp("[?&]" + field + "=([^&#]*)", "i");
    var string = reg.exec(href);
    return string ? string[1] : null;
}