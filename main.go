package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// route urls
	//
	// Auth
	router.POST("createAccount", createAccount)
	router.POST("deleteAccount", deleteAccount)
	router.POST("updateAccount", updateAccount)
	router.GET("login", login)
	router.GET("logout", logout)

	// Global
	router.GET("getApps", getApps)
	router.POST("createApp", createApp)
	router.DELETE("deleteApp", deleteApp)
	router.POST("hideApp", hideApp)

	router.Run("localhost:8000")
}

// Auth
func createAccount(c *gin.Context) {}
func deleteAccount(c *gin.Context) {}
func updateAccount(c *gin.Context) {}
func login(c *gin.Context) {}
func logout (c *gin.Context) {}

// Global
func getApps(c *gin.Context) {}
func createApp(c *gin.Context) {}
func deleteApp(c *gin.Context) {}
func hideApp(c *gin.Context) {}