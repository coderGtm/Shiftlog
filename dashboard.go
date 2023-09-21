package main

import (
	"strconv"
	"time"
)

type userApp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func getAppsOfUser(userId uint) []*userApp {
	apps := make([]*userApp, 0)
	rows, err := db.Query("SELECT id, name, hidden, createdAt, updatedAt from app WHERE userId = ?", userId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		app := new(userApp)
		err := rows.Scan(&app.Id, &app.Name, &app.Hidden, &app.CreatedAt, &app.UpdatedAt)
		checkErr(err)
		apps = append(apps, app)
	}
	err = rows.Err()
	checkErr(err)

	return apps
}

func createAppForUser(userId int, appName string) userApp {
	currentTimeStamp := time.Now().Unix()
	stmnt, err := db.Prepare("INSERT INTO APP(userId, name, hidden, createdAt, updatedAt) VALUES(?,?,?,?,?);")
	checkErr(err)
	defer stmnt.Close()
	res, err := stmnt.Exec(userId, appName, 0, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	appId, err := res.LastInsertId()
	checkErr(err)
	// return appId, name and hidden
	return userApp{
		int(appId), appName, false, currentTimeStamp, currentTimeStamp,
	}
}

func updateAppById(appId int, newName string, hidden int) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	stmnt, err := db.Prepare("UPDATE app SET name = ?, hidden = ?, updatedAt = ? WHERE id = ?")
	checkErr(err)
	defer stmnt.Close()
	_, err = stmnt.Exec(newName, hidden, currentTimeStamp, appId)
	checkErr(err)
}