package main

import "github.com/gin-gonic/gin"


func setUpRoutes(router *gin.Engine) {
	api := router.Group("api")
	// Auth
	api.POST("createAccount", createAccount)
	api.DELETE("deleteAccount", deleteAccount)
	api.PUT("updateUsername", updateUsername)
	api.PUT("updatePassword", updatePassword)
	api.POST("login", login)
	api.GET("logout", logout)

	// Dashboard
	api.GET("getApps", getApps)
	api.POST("createApp", createApp)
	api.DELETE("deleteApp", deleteApp)
	api.PUT("updateApp", updateApp)

	// App
	api.GET("getReleases", getReleases)
	api.POST("createRelease", createRelease)
	api.DELETE("deleteRelease", deleteRelease)
	api.PUT("updateRelease", updateRelease)

	// Release
	api.GET("getReleaseNotes", getReleaseNotes)
	api.PUT("updateReleaseNotes", updateReleaseNotes)

	frontend := router.Group("/")

	frontend.GET("", getHomePage)
	frontend.GET("signup", getSignupPage)
	frontend.GET("login", getLoginPage)
	frontend.GET("dashboard", getDashboardPage)
	frontend.GET("profile", getProfilePage)
	frontend.GET("app", getAppPage)
	frontend.GET("release", getReleasePage)
}
