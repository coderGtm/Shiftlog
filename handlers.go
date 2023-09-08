package main

import (
	"net/http"
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

func deleteAccount(c *gin.Context) {}
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

func logout(c *gin.Context) {}

// Dashboard
func getApps(c *gin.Context) {}
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

	appId, appName, appHidden := createAppForUser(int(userId), appName)
	data := userApp {
		Id: appId,
		Name: appName,
		Hidden: appHidden,
	}
	c.IndentedJSON(http.StatusOK, data)
}
func deleteApp(c *gin.Context) {}
func updateApp(c *gin.Context) {}

// App
func getReleases(c *gin.Context)   {}
func createRelease(c *gin.Context) {}
func deleteRelease(c *gin.Context) {}
func updateRelease(c *gin.Context) {}

// Release
func getReleaseNotesTXT(c *gin.Context)     {}
func getReleaseNotesMD(c *gin.Context)      {}
func getReleaseNotesHTML(c *gin.Context)    {}
func updateReleaseNotesTXT(c *gin.Context)  {}
func updateReleaseNotesMD(c *gin.Context)   {}
func updateReleaseNotesHTML(c *gin.Context) {}
