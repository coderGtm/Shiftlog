document.body.onload = function() {
    // Check if user is logged in
    if (!localStorage.getItem("authToken")) {
        // Redirect to login page
        window.location.href = "/login";
    }

    // Populate app name from query string
    document.querySelector("h2 strong").textContent = decodeURIComponent(getQueryString("name"));
    document.getElementById("appName").value = decodeURIComponent(getQueryString("name"));
    document.getElementById("hidden").checked = getQueryString("hidden") === "true";

    fetchReleases();
}

function fetchReleases() {
    // Get release list and display in table. API endpoint is /api/getReleases and an auth token is required in the header
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getReleases?appId=" + getQueryString("id"), true);
    xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // Hide loading spinner
                document.getElementById("loader").style.display = "none";
                // Parse JSON response and display in table
                var response = JSON.parse(xhr.responseText);
                var table = document.getElementById("releaseTable");
                if (response.length == 0) {
                    document.getElementById("noReleases").style.display = "block";
                }
                else {
                    for (var i = 0; i < response.length; i++) {
                        var row = table.insertRow(i + 1);
                        // Show serial no, version code, version name, release notes, and last modified
                        var cell0 = row.insertCell(0);
                        cell0.innerHTML = i + 1;
                        var cell1 = row.insertCell(1);
                        cell1.innerHTML = response[i].versionCode;
                        var cell2 = row.insertCell(2);
                        cell2.innerHTML = response[i].versionName;
                        var cell3 = row.insertCell(3);
                        cell3.innerHTML = response[i].hidden;
                        var cell4 = row.insertCell(4);
                        // updatedAt is a timestamp, convert to readable date and time
                        var updatedAt = new Date(response[i].updatedAt * 1000);
                        cell4.innerHTML = updatedAt.toLocaleString();
    
                        // Make row clickable and redirect to release page
                        row.onclick = function() {
                            var id = response[this.rowIndex - 1].id;
                            var versionName = response[this.rowIndex - 1].versionName
                            window.location.href = "/release?versionName=" + encodeURIComponent(versionName) + "&id=" + id;
                        };
                        row.style.cursor = "pointer";
                    }
                }
            } else if (xhr.status == 401) {
                // Auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else if (xhr.status == 400) {
                // Bad request
                alert("Error fetching releases! " + xhr.responseText);
            } else {
                alert("Error fetching releases! " + xhr.responseText);
            }
        }
    };
    xhr.send();
}

function deleteApp() {
    // Delete the app. API endpoint is /api/deleteApp and an auth token is required in the header
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/api/deleteApp?appId=" + getQueryString("id"), true);
    xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                alert("App deleted successfully!");
                // Redirect to dashboard after deleting the app
                window.location.href = "/dashboard";
            } else if (xhr.status == 401) {
                // Auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error deleting app! " + xhr.responseText);
            }
        }
    };
    xhr.send();
}

function updateApp() {
    var appName = document.getElementById("appName").value;
    var hidden = document.getElementById("hidden").checked;

    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updateApp", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                alert("App details updated successfully!");
                // Refresh the page after updating app details
                window.location.reload();
            } else if (xhr.status == 401) {
                // Auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error updating app details! " + xhr.responseText);
            }
        }
    };

    var params = "appId=" + getQueryString("id") + "&name=" + appName + "&hidden=" + hidden;
    xhr.send(params);
}

function createRelease() {
    var versionCode = document.getElementById("versionCode").value;
    var versionName = document.getElementById("versionName").value;
    var appId = getQueryString("id");

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/createRelease", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var response = JSON.parse(xhr.responseText);
                var id = response.id;
                alert("Release created successfully!");
                // Redirect to release page after creating the release
                window.location.href = "/release?versionName=" + encodeURIComponent(versionName) + "&id=" + id;
            } else if (xhr.status == 401) {
                // Auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error creating release! " + xhr.responseText);
            }
        }
    };

    var params = "versionCode=" + versionCode + "&versionName=" + versionName + "&appId=" + appId;
    xhr.send(params);
}

function logout() {
    // Remove authToken and username from localStorage and redirect to login page
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