package main

import "database/sql"

type releaseNotes struct {
	VersionCode int    `json:"versionCode"`
	VersionName string `json:"versionName"`
	NotesTxt    string `json:"notesTxt"`
	NotesMd     string `json:"notesMd"`
	NotesHtml   string `json:"notesHtml"`
	UpdatedAt   string  `json:"updatedAt"`
}

func getReleaseNotesOfRelease(releaseId int) (releaseNotes, bool) {
	var VersionCode int
	var VersionName string
	var NotesTxt string
	var NotesMd string
	var NotesHtml string
	var UpdatedAt string
	err := db.QueryRow("SELECT versionCode, versionName, notesTxt, notesMd, notesHtml, updatedAt from release WHERE id = ? AND hidden = 0", releaseId).Scan(&VersionCode, &VersionName, &NotesTxt, &NotesMd, &NotesHtml, &UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return releaseNotes{}, false
	}
	// record exists
	return releaseNotes{
		VersionCode: VersionCode,
		VersionName: VersionName,
		NotesTxt: NotesTxt,
		NotesMd: NotesMd,
		NotesHtml: NotesHtml,
		UpdatedAt: UpdatedAt,
	}, true
}

func getReleaseNotesByAppIdAndVersionCode(appId int, versionCode int, latestFlag bool) (releaseNotes, bool) {
	// returns releaseNotes, exists
	notes := new(releaseNotes)
	var err error
	if !latestFlag {
		err = db.QueryRow("SELECT versionCode, versionName, notesTxt, notesMd, notesHtml, updatedAt from release WHERE appId = ? AND versionCode = ? AND hidden = 0", appId, versionCode).Scan(&notes.VersionCode, &notes.VersionName, &notes.NotesTxt, &notes.NotesMd, &notes.NotesHtml, &notes.UpdatedAt)
	} else {
		maxVersionCode := getMaxVersionCodeOfApp(appId)
		if maxVersionCode == -1 {
			return *notes, false
		}
		err = db.QueryRow("SELECT versionCode, versionName, notesTxt, notesMd, notesHtml, updatedAt from release WHERE appId = ? AND versionCode = ? AND hidden = 0", appId, maxVersionCode).Scan(&notes.VersionCode, &notes.VersionName, &notes.NotesTxt, &notes.NotesMd, &notes.NotesHtml, &notes.UpdatedAt)
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
	err := db.QueryRow("SELECT MAX(versionCode) from release WHERE appId = ? AND hidden = 0", appId).Scan(&versionCode)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return -1
	}
	// record exists
	return versionCode
}