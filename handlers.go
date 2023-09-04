package main

import "github.com/gin-gonic/gin"

// Auth
func createAccount(c *gin.Context) {}
func deleteAccount(c *gin.Context) {}
func updateAccount(c *gin.Context) {}
func login(c *gin.Context) {}
func logout(c *gin.Context) {}

// Dashboard
func getApps(c *gin.Context) {}
func createApp(c *gin.Context) {}
func deleteApp(c *gin.Context) {}
func hideApp(c *gin.Context) {}

// App
func getReleases(c *gin.Context) {}
func createRelease(c *gin.Context) {}
func deleteRelease(c *gin.Context) {}
func hideRelease(c *gin.Context) {}

// Release
func getReleaseNotesTXT(c *gin.Context) {}
func getReleaseNotesMD(c *gin.Context) {}
func getReleaseNotesHTML(c *gin.Context) {}
func updateReleaseNotesTXT(c *gin.Context) {}
func updateReleaseNotesMD(c *gin.Context) {}
func updateReleaseNotesHTML(c *gin.Context) {}