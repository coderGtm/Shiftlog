package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	_ "github.com/mattn/go-sqlite3"
)
var db *sql.DB
var dbErr error
var htmlStripper *bluemonday.Policy

func main() {
	db, dbErr = sql.Open("sqlite3", "development.db")
	checkErr(dbErr)
	defer db.Close()

	htmlStripper = bluemonday.StrictPolicy()

	router := gin.Default()

	setUpRoutes(router)

	router.Run(":8080")
}