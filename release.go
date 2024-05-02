package main

import (
	"database/sql"
	"strconv"
	"time"
)

type releaseNotes struct {
	ReleaseId   int    `json:"id"`
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
	NotesTxt    string `json:"notesTxt"`
	NotesMd     string `json:"notesMd"`
	NotesHtml   string `json:"notesHtml"`
	Hiddden     bool   `json:"hidden"`
	Data        string `json:"data"`
	UpdatedAt   string `json:"updatedAt"`
}

func getReleaseNotesOfRelease(releaseId int) (releaseNotes, bool) {
	var VersionCode int
	var VersionName string
	var NotesTxt string
	var NotesMd string
	var NotesHtml string
	var Hiddden bool
	var Data string
	var UpdatedAt string
	err := db.QueryRow("SELECT versionCode, versionName, COALESCE(notesTxt, ''), COALESCE(notesMd, ''), COALESCE(notesHtml, ''), hidden, COALESCE(data, ''), updatedAt from release WHERE id = ?;", releaseId).Scan(&VersionCode, &VersionName, &NotesTxt, &NotesMd, &NotesHtml, &Hiddden, &Data, &UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		println(err.Error())
		// record does not exist
		return releaseNotes{}, false
	}
	// record exists
	return releaseNotes{
		ReleaseId:   releaseId,
		VersionCode: VersionCode,
		VersionName: VersionName,
		NotesTxt:    NotesTxt,
		NotesMd:     NotesMd,
		NotesHtml:   NotesHtml,
		Hiddden:     Hiddden,
		Data:        Data,
		UpdatedAt:   UpdatedAt,
	}, true
}

func getReleaseNotesByAppIdAndVersionCode(appId int, versionCode int, latestFlag bool) (releaseNotes, bool) {
	// returns releaseNotes, exists
	notes := new(releaseNotes)
	var err error
	if !latestFlag {
		err = db.QueryRow("SELECT id, versionCode, versionName, notesTxt, notesMd, notesHtml, hidden, COALESCE(data, ''), updatedAt from release WHERE appId = ? AND versionCode = ? AND hidden = 0", appId, versionCode).Scan(&notes.ReleaseId, &notes.VersionCode, &notes.VersionName, &notes.NotesTxt, &notes.NotesMd, &notes.NotesHtml, &notes.Hiddden, &notes.Data, &notes.UpdatedAt)
	} else {
		maxVersionCode := getMaxVersionCodeOfApp(appId)
		if maxVersionCode == -1 {
			return *notes, false
		}
		err = db.QueryRow("SELECT id, versionCode, versionName, notesTxt, notesMd, notesHtml, hidden, COALESCE(data, ''), updatedAt from release WHERE appId = ? AND versionCode = ? AND hidden = 0", appId, maxVersionCode).Scan(&notes.ReleaseId, &notes.VersionCode, &notes.VersionName, &notes.NotesTxt, &notes.NotesMd, &notes.NotesHtml, &notes.Hiddden, &notes.Data, &notes.UpdatedAt)
	}
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return *notes, false
	}
	// record exists
	return *notes, true
}

func getMaxVersionCodeOfApp(appId int) int {
	var versionCode int
	err := db.QueryRow("SELECT IFNULL(MAX(versionCode), 0) from release WHERE appId = ? AND hidden = 0", appId).Scan(&versionCode)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return -1
	}
	//check if versionCode is null
	if versionCode == 0 {
		return -1
	}

	// record exists
	return versionCode
}

func updateReleaseNotesById(id int, notesTxt string, notesMd string, notesHtml string) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	stmnt, err := db.Prepare("UPDATE release SET notesTxt = ?, notesMd = ?, notesHtml = ?, updatedAt = ? WHERE id = ?")
	checkErr(err)
	defer stmnt.Close()
	_, err = stmnt.Exec(notesTxt, notesMd, notesHtml, currentTimeStamp, id)
	checkErr(err)
}
