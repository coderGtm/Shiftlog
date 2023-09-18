package main

import (
	"time"
)

type userApp struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Hidden bool `json:"hidden"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}
type appRelease struct {
	Id     int    `json:"id"`
	AppId  int    `json:"appId"`
	VersionCode  int    `json:"versionCode"`
	VersionName  string `json:"versionName"`
	NotesTxt  string `json:"notesTxt"`
	NotesMd  string `json:"notesMd"`
	NotesHtml  string `json:"notesHtml"`
	Hidden bool `json:"hidden"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
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







