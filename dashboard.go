package main

import (
	"database/sql"
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
	res, err := stmnt.Exec(userId, appName, 0, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	appId, err := res.LastInsertId()
	checkErr(err)
	// return appId, name and hidden
	return userApp{
		int(appId), appName, false, currentTimeStamp, currentTimeStamp,
	}
}

func isAppOfUser(appId int, userId int) bool {
	var dbUserId int
		err := db.QueryRow("SELECT userId from APP WHERE id = ?;", appId).Scan(&dbUserId)
		if err != nil {
			if err != sql.ErrNoRows {
				// a real error happened!
				checkErr(err)
			}
			// record does not exist
			return false
		}
		// record exists
		if userId == dbUserId {
			return true
		}
		return false
}

func deleteUserApp(userId uint, appId uint) {
	stmnt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	checkErr(err)
	_, err = stmnt.Exec(userId)
	checkErr(err)
}

func isReleaseAlreadyPresent(appId int, versionCode int) bool {
	var releaseId int
	err := db.QueryRow("SELECT id FROM release WHERE appId = ? AND versionCode = ?", appId, versionCode).Scan(&releaseId)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return false
	}
	return true
}

func createReleaseForApp(userId int, appId int, versionCode int, versionName string) appRelease {
	currentTimeStamp := time.Now().Unix()
	stmnt, err := db.Prepare("INSERT INTO RELEASE(appId, versionCode, versionName, hidden, createdAt, updatedAt) VALUES(?,?,?,?,?,?);")
	checkErr(err)
	res, err := stmnt.Exec(appId, versionCode, versionName, 0, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	releaseId, err := res.LastInsertId()
	checkErr(err)

	return appRelease{
		int(releaseId), appId, versionCode, versionName, "", "", "", false, currentTimeStamp, currentTimeStamp,
	}
}

func getReleasesOfApp(appId int) []*appRelease {
	releases := make([]*appRelease, 0)
	rows, err := db.Query("SELECT id, versionCode, versionName, hidden, createdAt, updatedAt from release WHERE appId = ?", appId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		release := new(appRelease)
		err := rows.Scan(&release.Id, &release.VersionCode, &release.VersionName, &release.Hidden, &release.CreatedAt, &release.UpdatedAt)
		checkErr(err)
		releases = append(releases, release)
	}
	err = rows.Err()
	checkErr(err)

	return releases
}

func isReleaseOfUser(releaseId int, userId int) bool {
	var dbAppId int
		err := db.QueryRow("SELECT appId from release WHERE id = ?;", releaseId).Scan(&dbAppId)
		if err != nil {
			if err != sql.ErrNoRows {
				// a real error happened!
				checkErr(err)
			}
			// record does not exist
			return false
		}
		// record exists
		if isAppOfUser(dbAppId, userId) {
			return true
		}
		return false
}

func deleteAppRelease(releaseId int) {
	stmnt, err := db.Prepare("DELETE FROM release WHERE id = ?")
	checkErr(err)
	_, err = stmnt.Exec(releaseId)
	checkErr(err)
}