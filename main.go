package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// route urls
	//
	// Auth
	router.POST("createAccount", createAccount)
	router.DELETE("deleteAccount", deleteAccount)
	router.POST("updateAccount", updateAccount)
	router.GET("login", login)
	router.GET("logout", logout)

	// Global
	router.GET("getApps", getApps)
	router.POST("createApp", createApp)
	router.DELETE("deleteApp", deleteApp)
	router.POST("hideApp", hideApp)

	// App-specific
	router.GET("getReleases", getReleases)
	router.POST("createRelease", createRelease)
	router.DELETE("deleteRelease", deleteRelease)
	router.POST("hideRelease", hideRelease)

	// Release-specific
	router.GET("getReleaseNotesTXT",getReleaseNotesTXT)
	router.GET("getReleaseNotesMD", getReleaseNotesMD)
	router.GET("getReleaseNotesHTML",getReleaseNotesHTML)
	router.POST("updateReleaseNotesTXT",updateReleaseNotesTXT)
	router.POST("updateReleaseNotesMD",updateReleaseNotesMD)
	router.POST("updateReleaseNotesHTML",updateReleaseNotesHTML)

	router.Run(":8000")
}

// Auth
func createAccount(c *gin.Context) {}
func deleteAccount(c *gin.Context) {}
func updateAccount(c *gin.Context) {}
func login(c *gin.Context) {}
func logout(c *gin.Context) {}

// Global
func getApps(c *gin.Context) {}
func createApp(c *gin.Context) {}
func deleteApp(c *gin.Context) {}
func hideApp(c *gin.Context) {}

// App-specific
func getReleases(c *gin.Context) {}
func createRelease(c *gin.Context) {}
func deleteRelease(c *gin.Context) {}
func hideRelease(c *gin.Context) {}

// Release-specific
func getReleaseNotesTXT(c *gin.Context) {}
func getReleaseNotesMD(c *gin.Context) {}
func getReleaseNotesHTML(c *gin.Context) {}
func updateReleaseNotesTXT(c *gin.Context) {}
func updateReleaseNotesMD(c *gin.Context) {}
func updateReleaseNotesHTML(c *gin.Context) {}