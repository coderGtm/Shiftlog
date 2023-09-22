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
	db, dbErr = sql.Open("sqlite3", "development.db")
	checkErr(dbErr)
	defer db.Close()
	_, dbErr = db.Exec("PRAGMA foreign_keys = ON;")
	checkErr(dbErr)

	htmlStripper = bluemonday.StrictPolicy()
	notesSanitizer = bluemonday.UGCPolicy()
	err := godotenv.Load(".env")
	checkErr(err)

	router := gin.Default()

	setUpRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}