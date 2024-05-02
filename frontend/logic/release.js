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
    preservedReleaseId = releaseId;
    
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getReleaseNotes?releaseId=" + releaseId, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var releaseData = JSON.parse(xhr.responseText);
                releaseId = releaseData.id;
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
    console.log(releaseData);
    // Update page elements with fetched release data
    document.getElementById("versionNameHeader").innerHTML = "<strong>"+releaseData.versionName+"</strong>"
    document.getElementById("versionName").textContent = releaseData.versionName;
    document.getElementById("uVersionCode").value = releaseData.versionCode;
    document.getElementById("uVersionName").value = releaseData.versionName;
    document.getElementById("versionCode").textContent = releaseData.versionCode;
    document.getElementById("uHidden").checked = releaseData.hidden;
    document.getElementById("hidden").textContent = releaseData.hidden ? "Yes" : "No";
    document.getElementById("data").textContent = releaseData.data;
    document.getElementById("remarks").textContent = releaseData.data;
    document.getElementById("api").textContent = "/api/getReleaseNotes?releaseId=" + releaseData.id;

    // updatedAt is a timestamp, the number of seconds elapsed since January 1, 1970 UTC, convert to readable date and time in am pm format
    var updatedAt = new Date(releaseData.updatedAt * 1000);
    var hours = updatedAt.getHours();
    var minutes = updatedAt.getMinutes();
    var ampm = hours >= 12 ? 'pm' : 'am';
    hours = hours % 12;
    hours = hours ? hours : 12; // the hour '0' should be '12'
    minutes = minutes < 10 ? '0'+minutes : minutes;
    var strTime = hours + ':' + minutes + ' ' + ampm;
    document.getElementById("lastModified").textContent = updatedAt.toDateString() + " " + strTime;

    setReleaseNotes(releaseData);
}

function setReleaseNotes(releaseData) {
    var rTxt = releaseData.notesTxt;
    var rMd = releaseData.notesMd;
    var rHtml = releaseData.notesHtml;
    
    var txtDiv = document.getElementById("releaseNotesText");
    var mdDiv = document.getElementById("releaseNotesMarkdown");
    var htmlDiv = document.getElementById("releaseNotesHTML");

    if (rTxt == null || rTxt == "") {
        txtDiv.innerHTML = "<em>No release notes provided!</em>";
        document.getElementById("text-edit").textContent = "";
    }
    else {
        txtDiv.innerHTML = "<pre>"+rTxt+"</pre>";
        document.getElementById("text-edit").textContent = rTxt;
    }

    if (rMd == null || rMd == "") {
        mdDiv.innerHTML = "<em>No release notes provided!</em>";
        document.getElementById("markdown-edit").textContent = "";
    }
    else {
        mdDiv.innerHTML = "<md-block>"+rMd+"</md-block>";
        document.getElementById("markdown-edit").textContent = rMd;
    }

    if (rHtml == null || rHtml == "") {
        htmlDiv.innerHTML = "<em>No release notes provided!</em>";
        document.getElementById("html-edit").textContent = "";
    }
    else {
        htmlDiv.innerHTML = "<iframe srcdoc='"+rHtml+"' onload='resizeIframe(this)'></iframe>";
        document.getElementById("html-edit").textContent = rHtml;
    }

    sessionStorage.setItem("rTxt", rTxt);
    sessionStorage.setItem("rMd", rMd);
    sessionStorage.setItem("rHtml", rHtml);
}

function updateRelease() {
    var versionCode = document.getElementById("uVersionCode").value;
    var versionName = document.getElementById("uVersionName").value;
    var hidden = document.getElementById("uHidden").checked;
    var data = document.getElementById("data").value;

    var authToken = localStorage.getItem("authToken");

    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updateRelease", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer " + authToken);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // Release updated successfully, you can redirect or show a success message
                alert("Release updated successfully!");
                window.location.reload();
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                alert("Session expired, please login again!");
                window.location.href = "/login";
            } else {
                alert("Error updating release! " + xhr.responseText);
            }
        }
    };
    var params = "releaseId=" + preservedReleaseId + "&versionCode=" + versionCode + "&versionName=" + versionName + "&hidden=" + hidden + "&data=" + data;
    xhr.send(params);
}

function updateText() {
    var rTxt = document.getElementById("text-edit").value;
    var rMd = sessionStorage.getItem("rMd");
    var rHtml = sessionStorage.getItem("rHtml");

    updateReleaseNotes(rTxt, rMd, rHtml);
}

function updateMarkdown() {
    var rTxt = sessionStorage.getItem("rTxt");
    var rMd = document.getElementById("markdown-edit").value;
    var rHtml = sessionStorage.getItem("rHtml");

    updateReleaseNotes(rTxt, rMd, rHtml);
}

function updateHtml() {
    var rTxt = sessionStorage.getItem("rTxt");
    var rMd = sessionStorage.getItem("rMd");
    var rHtml = document.getElementById("html-edit").value;

    updateReleaseNotes(rTxt, rMd, rHtml);
}

function updateReleaseNotes(rTxt, rMd, rHtml) {
    var authToken = localStorage.getItem("authToken");

    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updateReleaseNotes", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer " + authToken);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // Release updated successfully, you can redirect or show a success message
                alert("Release notes updated successfully!");
                window.location.reload();
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                alert("Session expired, please login again!");
                window.location.href = "/login";
            } else {
                alert("Error updating release notes! " + xhr.responseText);
            }
        }
    };
    var params = "releaseId=" + preservedReleaseId + "&notesTxt=" + rTxt + "&notesMd=" + rMd + "&notesHtml=" + rHtml;
    xhr.send(params);
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

function resizeIframe(obj) {

    obj.style.height = obj.contentWindow.document.body.scrollHeight + 'px';

  }

var preservedReleaseId = -1;