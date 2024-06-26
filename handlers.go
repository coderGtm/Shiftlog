package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth
func createAccount(c *gin.Context) {
	// get sanatized parameters
	username := htmlStripper.Sanitize(c.PostForm("username"))
	password := htmlStripper.Sanitize(c.PostForm("password"))

	// check for empty params
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if username != c.PostForm("username") || password != c.PostForm("password") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal username or password!")
		return
	}

	// check is username already used
	if checkUsernameExists(username) {
		c.IndentedJSON(http.StatusConflict, "Username Already Exists")
		return
	}
	authToken := registerNewUser(username, password)

	data := createAccountSuccessResponse{
		Username:  username,
		AuthToken: authToken,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func deleteAccount(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}
	deleteUserAccount(userId)
	c.IndentedJSON(http.StatusOK, "User Account Deleted Successfully!")
}
func updateUsername(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get sanatized parameters
	newUsername := htmlStripper.Sanitize(c.PostForm("newUsername"))

	// check for empty params
	if strings.Trim(newUsername, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if newUsername != c.PostForm("newUsername") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal username provided!")
		return
	}

	// check is username already used
	if checkUsernameExists(newUsername) {
		c.IndentedJSON(http.StatusConflict, "This username is already taken!")
		return
	}

	updateUsernameById(int(userId), newUsername)
	c.IndentedJSON(http.StatusOK, "Username updated successfully!")
}

func updatePassword(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get sanatized parameters
	newPassword := htmlStripper.Sanitize(c.PostForm("newPassword"))

	// check for empty params
	if strings.Trim(newPassword, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if newPassword != c.PostForm("newPassword") {
		c.IndentedJSON(http.StatusBadRequest, "Password contain illegal charachters!")
		return
	}

	// check if same password
	if checkIfSamePassword(int(userId), newPassword) {
		c.IndentedJSON(http.StatusConflict, "New Password cannot be same as old Password!")
		return
	}

	authToken = updatePasswordById(int(userId), newPassword)
	data := passwordUpdateSuccessResponse{
		AuthToken: authToken,
	}
	c.IndentedJSON(http.StatusOK, data)
}

func login(c *gin.Context) {
	// get sanatized parameters
	username := htmlStripper.Sanitize(c.PostForm("username"))
	password := htmlStripper.Sanitize(c.PostForm("password"))

	// check for empty params
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if username != c.PostForm("username") || password != c.PostForm("password") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal username or password!")
		return
	}

	// check if username exists
	if !checkUsernameExists(username) {
		c.IndentedJSON(http.StatusConflict, "Username does not exist!")
		return
	}

	authenticated, authToken := verifyAndLogin(username, password)

	if !authenticated {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid login credentials!")
		return
	} else {
		data := loginSuccessResponse{
			Username: username,
			AuthToken: authToken,
		}
		c.IndentedJSON(http.StatusOK, data)
		return
	}
}

func logout(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}
	logoutUser(userId)
	c.IndentedJSON(http.StatusOK, "Logged out!")
}

// Dashboard
func getApps(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	apps := getAppsOfUser(userId)
	c.IndentedJSON(http.StatusOK, apps)
}

func createApp(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get input name
	appName := htmlStripper.Sanitize(c.PostForm("appName"))

	if appName != c.PostForm("appName") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal app Name")
		return
	}

	if strings.Trim(appName, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty app Name is not allowed.")
		return
	}

	app := createAppForUser(int(userId), appName)
	c.IndentedJSON(http.StatusOK, app)
}
func deleteApp(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get input id
	appId := htmlStripper.Sanitize(c.Query("appId"))
	println(appId)

	if appId != c.Query("appId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal app Id")
		return
	}
	intAppId, err := strconv.Atoi(appId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "appId must be an Integer.")
		return
	}
	if isAppOfUser(intAppId, int(userId)) {
		deleteAppById(intAppId)
		c.IndentedJSON(http.StatusOK, "App deleted successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized deletion!")
}
func updateApp(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get sanatized parameters
	// get input id
	appId := htmlStripper.Sanitize(c.PostForm("appId"))
	newName := htmlStripper.Sanitize(c.PostForm("name"))
	newHidden := htmlStripper.Sanitize(c.PostForm("hidden"))

	// check for empty params
	if strings.Trim(appId, " ") == "" || strings.Trim(newName, " ") == "" || strings.Trim(newHidden, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if newName != c.PostForm("name") || newHidden != c.PostForm("hidden") || appId != c.PostForm("appId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal values provided!")
		return
	}
	var hidden int
	switch newHidden {
	case "true":
		hidden = 1
	case "false":
		hidden = 0
	default:
		c.IndentedJSON(http.StatusBadRequest, "Hiddden parameter must have a 'true' or 'false' value")
	}

	intAppId, err := strconv.Atoi(appId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "appId must be an Integer.")
		return
	}
	if isAppOfUser(intAppId, int(userId)) {
		updateAppById(intAppId, newName, hidden)
		c.IndentedJSON(http.StatusOK, "App updated successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized update!")
}

// App
func getReleases(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}
	// get input id
	appId := htmlStripper.Sanitize(c.Query("appId"))

	if appId != c.Query("appId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal app Id")
		return
	}
	i_appId, err := strconv.Atoi(appId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "App ID must be an Integer.")
		return
	}

	if isAppOfUser(i_appId, int(userId)) {
		releases := getReleasesOfApp(i_appId)
		c.IndentedJSON(http.StatusOK, releases)
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized access!")
}
func createRelease(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get inputs
	appId := htmlStripper.Sanitize(c.PostForm("appId"))
	versionName := htmlStripper.Sanitize(c.PostForm("versionName"))
	versionCode := htmlStripper.Sanitize(c.PostForm("versionCode"))

	if appId != c.PostForm("appId") || versionName != c.PostForm("versionName") || versionCode != c.PostForm("versionCode") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal input parameter values")
		return
	}
	i_appId, err := strconv.Atoi(appId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "App ID must be an integer")
		return
	}
	i_versionCode, err := strconv.Atoi(versionCode)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Version Code must be an integer")
		return
	}
	if strings.Trim(versionName, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty Version Name is not allowed.")
		return
	}
	if isAppOfUser(i_appId, int(userId)) {
		if !isReleaseAlreadyPresent(i_appId, i_versionCode) {
			release := createReleaseForApp(int(userId), i_appId, i_versionCode, versionName)
			c.IndentedJSON(http.StatusOK, release)
			return
		} else {
			c.IndentedJSON(http.StatusBadRequest, "This Version Code already exists")
			return
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, "Unauthorized Request!")
		return
	}
}
func deleteRelease(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}
	// get input id
	releaseId := htmlStripper.Sanitize(c.Query("releaseId"))

	if releaseId != c.Query("releaseId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal Release Id")
		return
	}
	i_releaseId, err := strconv.Atoi(releaseId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "releaseId must be an Integer.")
		return
	}
	if isReleaseOfUser(i_releaseId, int(userId)) {
		deleteAppRelease(i_releaseId)
		c.IndentedJSON(http.StatusOK, "Release deleted successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Delete Request Unauthorized!")
}
func updateRelease(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get sanatized parameters
	// get input id
	releaseId := htmlStripper.Sanitize(c.Request.PostFormValue("releaseId"))
	newName := htmlStripper.Sanitize(c.Request.PostFormValue("versionName"))
	newCode := htmlStripper.Sanitize(c.Request.PostFormValue("versionCode"))
	newHidden := htmlStripper.Sanitize(c.Request.PostFormValue("hidden"))
	data := htmlStripper.Sanitize(c.Request.PostFormValue("data"))

	println(releaseId, newName, newCode, newHidden, data)

	// check for empty params
	if strings.Trim(releaseId, " ") == "" || strings.Trim(newName, " ") == "" || strings.Trim(newHidden, " ") == "" || strings.Trim(newCode, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Empty parameters in Request Body")
		return
	}

	// check for illegal params
	if newName != c.Request.PostFormValue("versionName") || newCode != c.Request.PostFormValue("versionCode") || newHidden != c.Request.PostFormValue("hidden") || releaseId != c.Request.PostFormValue("releaseId") || data != c.Request.PostFormValue("data") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal values provided!")
		return
	}
	var hidden int
	switch newHidden {
	case "true":
		hidden = 1
	case "false":
		hidden = 0
	default:
		c.IndentedJSON(http.StatusBadRequest, "Hiddden parameter must have a 'true' or 'false' value")
		return
	}

	intReleaseId, err := strconv.Atoi(releaseId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Release Id must be an Integer.")
		return
	}
	intVersionCode, err := strconv.Atoi(newCode)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Version Code must be an Integer.")
		return
	}
	if isReleaseOfUser(intReleaseId, int(userId)) {
		updateReleaseById(intReleaseId, newName, intVersionCode, data, hidden)
		c.IndentedJSON(http.StatusOK, "Release Details updated successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized update!")
}

// Release
func getReleaseNotes(c *gin.Context) {
	// unprotected endpoint

	// 2 methods, ordered by priority:
	// 1) Directly by release id
	// 2) By app id and version code (latest keyword allowed)

	releaseId := htmlStripper.Sanitize(c.Query("releaseId"))
	appId := htmlStripper.Sanitize(c.Query("appId"))
	versionCode := htmlStripper.Sanitize(c.Query("versionCode"))

	if releaseId != "" {
		if releaseId != c.Query("releaseId") {
			c.IndentedJSON(http.StatusBadRequest, "Illegal Release ID")
			return
		}
		releaseId, err := strconv.Atoi(releaseId)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Release ID must be an Integer!")
			return
		}
		releaseNotes, exists := getReleaseNotesOfRelease(releaseId)
		if !exists {
			c.IndentedJSON(http.StatusNotFound, "Release Notes not found!")
			return
		}
		
		c.IndentedJSON(http.StatusOK, releaseNotes)
		return
	} else if appId != "" && versionCode != "" {
		latestFlag := false
		if appId != c.Query("appId") || versionCode != c.Query("versionCode") {
			c.IndentedJSON(http.StatusBadRequest, "Illegal App ID or Version Code")
			return
		}
		appId, err := strconv.Atoi(appId)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "App ID must be an Integer!")
			return
		}
		i_versionCode := -1
		if versionCode == "latest" {
			latestFlag = true
		} else {
			versionCode, err := strconv.Atoi(versionCode)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, "Invalid Version Code")
				return
			}
			i_versionCode = versionCode
		}
		notes, exists := getReleaseNotesByAppIdAndVersionCode(appId, i_versionCode, latestFlag)
		if !exists {
			c.IndentedJSON(http.StatusNotFound, "Release Notes not found!")
			return
		}
		c.IndentedJSON(http.StatusOK, notes)
		return
	}
	c.IndentedJSON(http.StatusBadRequest, "Missing Parameters!")
}
func updateReleaseNotes(c *gin.Context) {
	authToken := extractAuthToken(c)
	if authToken == "" {
		c.IndentedJSON(http.StatusUnauthorized, "Auth token missing!")
		return
	}
	userId, validToken := isTokenValid(authToken)
	if !validToken {
		c.IndentedJSON(http.StatusUnauthorized, "Invalid Auth Token")
		return
	}

	// get sanatized parameters
	// get input id
	releaseId := htmlStripper.Sanitize(c.Request.PostFormValue("releaseId"))
	notesTxt := htmlStripper.Sanitize(c.Request.PostFormValue("notesTxt"))
	notesMd := notesSanitizer.Sanitize(c.Request.PostFormValue("notesMd"))
	notesHtml := notesSanitizer.Sanitize(c.Request.PostFormValue("notesHtml"))

	println("notesTxt: ", notesTxt)

	// check for empty params
	if strings.Trim(releaseId, " ") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Missing Release ID")
		return
	}

	// check for illegal params
	if releaseId != c.Request.PostFormValue("releaseId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal values for Release ID provided!")
		return
	}

	intReleaseId, err := strconv.Atoi(releaseId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Release Id must be an Integer.")
		return
	}

	if isReleaseOfUser(intReleaseId, int(userId)) {
		updateReleaseNotesById(intReleaseId, notesTxt, notesMd, notesHtml)
		c.IndentedJSON(http.StatusOK, "Release Notes updated successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized update!")
}



// frontend

func getHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func getSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}
func getLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func getDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
func getProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", nil)
}
func getAppPage(c *gin.Context) {
	c.HTML(http.StatusOK, "app.html", nil)
}
func getReleasePage(c *gin.Context) {
	c.HTML(http.StatusOK, "release.html", nil)
}