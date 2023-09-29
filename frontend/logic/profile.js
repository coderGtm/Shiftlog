function updateUsername() {
    var username = document.getElementById("username").value;
    // check if username is empty
    if (username == "") {
        alert("Username cannot be empty!");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updateUsername", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer "+ localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var response = JSON.parse(xhr.responseText);
                alert("Username updated successfully!");
                // refresh page
                window.location.reload();
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error updating username! " + xhr.responseText);
            }
        }
    }
    // send postform body parameters
    xhr.send("newUsername=" + username);
}

function updatePassword() {
    var p1 = document.getElementById("password1").value;
    var p2 = document.getElementById("password2").value;
    // check if password is empty
    if (p1 == "") {
        alert("Password cannot be empty!");
        return;
    }
    // check if passwords match
    if (p1 != p2) {
        alert("Passwords do not match!");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/updatePassword", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer "+ localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // response contains new auth token, update localStorage
                var response = JSON.parse(xhr.responseText);
                localStorage.setItem("authToken", response.authToken);
                alert("Password updated successfully!");
                // refresh page
                window.location.reload();
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error updating password! " + xhr.responseText);
            }
        }
    }
    // send postform body parameters
    xhr.send("newPassword=" + p1);
}

function deleteAccount() {
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/api/deleteAccount", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("Authorization", "Bearer "+ localStorage.getItem("authToken"));
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                alert("Account deleted successfully!");
                // redirect to login page
                window.location.href = "/login";
            } else if (xhr.status == 401) {
                // auth token expired or invalid, redirect to login page
                window.location.href = "/login";
            } else {
                alert("Error deleting account! " + xhr.responseText);
            }
        }
    }
    xhr.send();
}


function logout() {
    // remove authToken and username from localStorage and redirect to login page
    localStorage.removeItem("authToken");
    localStorage.removeItem("username");
    window.location.href = "/login";
}