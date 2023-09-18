package main

import "github.com/gin-gonic/gin"

func setUpRoutes(router *gin.Engine) {
	// Auth
	router.POST("createAccount", createAccount)
	router.DELETE("deleteAccount", deleteAccount)
	router.PUT("updateUsername", updateUsername)
	router.PUT("updatePassword", updatePassword)
	router.POST("login", login)
	router.GET("logout", logout)

	// Dashboard
	router.GET("getApps", getApps)
	router.POST("createApp", createApp)
	router.DELETE("deleteApp", deleteApp)
	router.PUT("updateApp", updateApp)

	// App
	router.GET("getReleases", getReleases)
	router.POST("createRelease", createRelease)
	router.DELETE("deleteRelease", deleteRelease)
	router.POST("updateRelease", updateRelease)

	// Release
	router.GET("getReleaseNotes", getReleaseNotes)
	router.POST("updateReleaseNotesTxt", updateReleaseNotes)
}
