package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/microcosm-cc/bluemonday"
)
var db *sql.DB
var dbErr error
var htmlStripper *bluemonday.Policy
var notesSanitizer *bluemonday.Policy

func main() {
	htmlStripper = bluemonday.StrictPolicy()
	notesSanitizer = bluemonday.UGCPolicy()
	err := godotenv.Load(".env")
	checkErr(err)

	db, dbErr = sql.Open("sqlite3", os.Getenv("DATABASE_PATH"))
	checkErr(dbErr)
	defer db.Close()
	_, dbErr = db.Exec("PRAGMA foreign_keys = ON;")
	checkErr(dbErr)

	router := gin.Default()

	router.Static("/logic", "frontend/logic")
	router.LoadHTMLGlob("frontend/*.html")

	setUpRoutes(router)

	if os.Getenv("PORT") == "" {
		router.Run(":8080")
	}
	router.Run(":" + os.Getenv("PORT"))
}