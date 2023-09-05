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

	// check is username already used
	if checkUsernameExists(username) {
		c.IndentedJSON(http.StatusConflict, "Username Already Exists")
		return
	}
	newUserId := registerNewUser(username, password)

	data := createAccountSuccessResponse {
		Id: newUserId,
		Username: username,
	}
	c.IndentedJSON(http.StatusOK, data)
}
func deleteAccount(c *gin.Context) {}
func updateAccount(c *gin.Context) {}
func login(c *gin.Context) {}
func logout(c *gin.Context) {}

// Dashboard
func getApps(c *gin.Context) {}
func createApp(c *gin.Context) {}
func deleteApp(c *gin.Context) {}
func updateApp(c *gin.Context) {}

// App
func getReleases(c *gin.Context) {}
func createRelease(c *gin.Context) {}
func deleteRelease(c *gin.Context) {}
func updateRelease(c *gin.Context) {}

// Release
func getReleaseNotesTXT(c *gin.Context) {}
func getReleaseNotesMD(c *gin.Context) {}
func getReleaseNotesHTML(c *gin.Context) {}
func updateReleaseNotesTXT(c *gin.Context) {}
func updateReleaseNotesMD(c *gin.Context) {}
func updateReleaseNotesHTML(c *gin.Context) {}