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

	// Dashboard
	router.GET("getApps", getApps)
	router.POST("createApp", createApp)
	router.DELETE("deleteApp", deleteApp)
	router.POST("hideApp", hideApp)

	// App
	router.GET("getReleases", getReleases)
	router.POST("createRelease", createRelease)
	router.DELETE("deleteRelease", deleteRelease)
	router.POST("hideRelease", hideRelease)

	// Release
	router.GET("getReleaseNotesTXT",getReleaseNotesTXT)
	router.GET("getReleaseNotesMD", getReleaseNotesMD)
	router.GET("getReleaseNotesHTML",getReleaseNotesHTML)
	router.POST("updateReleaseNotesTXT",updateReleaseNotesTXT)
	router.POST("updateReleaseNotesMD",updateReleaseNotesMD)
	router.POST("updateReleaseNotesHTML",updateReleaseNotesHTML)

	router.Run(":8000")
}