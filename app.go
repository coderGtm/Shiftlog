package main

import (
	"database/sql"
	"strconv"
	"time"
)

type appRelease struct {
	Id          int    `json:"id"`
	AppId       int    `json:"appId"`
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
	NotesTxt    string `json:"notesTxt"`
	NotesMd     string `json:"notesMd"`
	NotesHtml   string `json:"notesHtml"`
	Data		string `json:"data"`
	Hidden      bool   `json:"hidden"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func deleteAppRelease(releaseId int) {
	stmnt, err := db.Prepare("DELETE FROM release WHERE id = ?")
	checkErr(err)
	defer stmnt.Close()
	_, err = stmnt.Exec(releaseId)
	checkErr(err)
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

func createReleaseForApp(userId int, appId int, versionCode int, versionName string) appRelease {
	currentTimeStamp := time.Now().Unix()
	stmnt, err := db.Prepare("INSERT INTO RELEASE(appId, versionCode, versionName, hidden, createdAt, updatedAt) VALUES(?,?,?,?,?,?);")
	checkErr(err)
	defer stmnt.Close()
	res, err := stmnt.Exec(appId, versionCode, versionName, 0, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	releaseId, err := res.LastInsertId()
	checkErr(err)

	return appRelease{
		int(releaseId), appId, versionCode, versionName, "", "", "", "", false, currentTimeStamp, currentTimeStamp,
	}
}

func getReleasesOfApp(appId int) []*appRelease {
	releases := make([]*appRelease, 0)
	rows, err := db.Query("SELECT id, versionCode, versionName, data, hidden, createdAt, updatedAt from release WHERE appId = ?", appId)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		release := new(appRelease)
		err := rows.Scan(&release.Id, &release.VersionCode, &release.VersionName, &release.Data, &release.Hidden, &release.CreatedAt, &release.UpdatedAt)
		checkErr(err)
		releases = append(releases, release)
	}
	err = rows.Err()
	checkErr(err)

	return releases
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

func deleteAppById(appId int) {
	stmnt, err := db.Prepare("DELETE FROM app WHERE id = ?")
	checkErr(err)
	defer stmnt.Close()
	_, err = stmnt.Exec(appId)
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

func updateReleaseById(id int, name string, code int, data string, hidden int) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	stmnt, err := db.Prepare("UPDATE release SET versionName = ?, versionCode = ?, data = ?, hidden = ?, updatedAt = ? WHERE id = ?")
	checkErr(err)
	defer stmnt.Close()
	_, err = stmnt.Exec(name, code, data, hidden, currentTimeStamp, id)
	checkErr(err)
}