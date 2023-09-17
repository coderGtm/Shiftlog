package main

import (
	"database/sql"
	"time"
)

func deleteAppRelease(releaseId int) {
	stmnt, err := db.Prepare("DELETE FROM release WHERE id = ?")
	checkErr(err)
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