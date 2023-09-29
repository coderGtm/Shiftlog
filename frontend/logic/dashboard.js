document.body.onload = function() {
    // check if user is logged in
    if (!localStorage.getItem("authToken")) {
        // redirect to login page
        window.location.href = "/login";
    }
    // populate username
    document.getElementById("username").innerHTML = localStorage.getItem("username");
    fetchDashboardData();
}

function fetchDashboardData() {
    // get app list and display in table. api endpoint is /api/getApps and an auth token is required in the header
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getApps", true);
    xhr.setRequestHeader("Authorization", "Bearer "+ localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // hide loading spinner
                document.getElementById("loader").style.display = "none";
                // parse json response and display in table
                var response = JSON.parse(xhr.responseText);
                // response format is like [{id:<id>, name: <name>, hidden: [true/false], createdAt: <timestamp>, updatedAt: <timestamp>}, ...]
                var table = document.getElementById("appTable");
                for (var i = 0; i < response.length; i++) {
                    var row = table.insertRow(i + 1);
                    // show serial no, name, hidden, and last modified
                    var cell0 = row.insertCell(0);
                    cell0.innerHTML = i + 1;
                    var cell1 = row.insertCell(1);
                    cell1.innerHTML = response[i].name;
                    var cell2 = row.insertCell(2);
                    cell2.innerHTML = response[i].hidden;
                    var cell3 = row.insertCell(3);
                    // updatedAt is a timestamp, the number of seconds elapsed since January 1, 1970 UTC, convert to readable date and time in am pm format
                    var updatedAt = new Date(response[i].updatedAt * 1000);
                    var hours = updatedAt.getHours();
                    var minutes = updatedAt.getMinutes();
                    var ampm = hours >= 12 ? 'pm' : 'am';
                    hours = hours % 12;
                    hours = hours ? hours : 12; // the hour '0' should be '12'
                    minutes = minutes < 10 ? '0'+minutes : minutes;
                    var strTime = hours + ':' + minutes + ' ' + ampm;
                    cell3.innerHTML = updatedAt.toDateString() + " " + strTime;

                    // make row clickable and redirect to /app?id=<id>?name=<name>
                    row.onclick = function() {
                        var id = response[this.rowIndex - 1].id;
                        var name = response[this.rowIndex - 1].name;
                        var urlEncodedQuery = "?id=" + encodeURIComponent(id) + "&name=" + encodeURIComponent(name);
                        window.location.href = "/app" + urlEncodedQuery;
                    }
                    row.style.cursor = "pointer";
                }
            }
            else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            }
            else if (xhr.status == 400) {
                // bad request
                alert("App not Created! " + xhr.responseText);
            } else {
                alert("Error fetching apps! " + xhr.responseText);
            }
        }
    }
    xhr.send();
}

function createApp() {
    var name = document.getElementById("appName").value;
    // check if name is empty
    if (name == "") {
        alert("App name cannot be empty!");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/createApp", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer "+ localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // redirect to /app?id=<id>?name=<name>
                var response = JSON.parse(xhr.responseText);
                var id = response.id;
                var name = response.name;
                var urlEncodedQuery = "?id=" + encodeURIComponent(id) + "&name=" + encodeURIComponent(name);
                window.location.href = "/app" + urlEncodedQuery;
            } else {
                alert("Error creating app! " + xhr.responseText);
            }
        }
    }
    // send postform body parameters
    xhr.send("appName=" + name);
}

function logout() {
    // remove authToken and username from localStorage and redirect to login page
    localStorage.removeItem("authToken");
    localStorage.removeItem("username");
    window.location.href = "/login";
}