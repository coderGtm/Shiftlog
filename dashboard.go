package main

import (
	"strconv"
	"time"
)

type userApp struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Hidden bool `json:"hidden"`
}

func getAppsOfUser(userId int) {

}

func createAppForUser(userId int, appName string) (int, string, bool) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	stmnt, err := db.Prepare("INSERT INTO APP(userId, name, hidden, createdAt, updatedAt) VALUES(?,?,?,?,?);")
	checkErr(err)
	res, err := stmnt.Exec(userId, appName, 0, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	appId, err := res.LastInsertId()
	checkErr(err)
	// return appId, name and hidden
	return int(appId), appName, false
}
