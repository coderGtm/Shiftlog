function login() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("pwd1").value;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/login", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                // save username and authToken from json response to localStorage
                var response = JSON.parse(xhr.responseText);
                localStorage.setItem("username", response.username);
                localStorage.setItem("authToken", response.authToken);
                alert("Login successful!");
                window.location.href = "/dashboard";
            } else {
                alert("Login failed! " + xhr.responseText);
            }
        }
    }
    // send postform body parameters
    xhr.send("username=" + username + "&password=" + password);
}