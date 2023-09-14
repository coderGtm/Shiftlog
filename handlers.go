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
func updateAccount(c *gin.Context) {}

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
	appId := htmlStripper.Sanitize(c.PostForm("appId"))

	if appId != c.PostForm("appId") {
		c.IndentedJSON(http.StatusBadRequest, "Illegal app Id")
		return
	}
	intAppId, err := strconv.Atoi(appId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "appId must be an Integer.")
		return
	}
	if isAppOfUser(intAppId, int(userId)) {
		deleteUserApp(userId, uint(intAppId))
		c.IndentedJSON(http.StatusOK, "App deleted successfully!")
		return
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorized deletion!")
}
func updateApp(c *gin.Context) {}

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
func deleteRelease(c *gin.Context) {}
func updateRelease(c *gin.Context) {}

// Release
func getReleaseNotesTxt(c *gin.Context)     {}
func getReleaseNotesMd(c *gin.Context)      {}
func getReleaseNotesHtml(c *gin.Context)    {}
func updateReleaseNotes(c *gin.Context)  {}
