package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func checkUsernameExists(uname string) bool {
	stmnt, err := db.Prepare("SELECT * FROM USER WHERE username = ?")
	checkErr(err)

	res, err := stmnt.Exec(uname)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect > 0
}

func registerNewUser(uname string, pswd string) {
	currentTimeStamp := string(time.Now().Unix())
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pswd),bcrypt.DefaultCost)
	checkErr(err)
	hashedPasswordString := string(hashedPassword)
	stmnt, err := db.Prepare("INSERT INTO USER values(username, password, createdAt, updatedAt) values(?,?,?,?);")
	checkErr(err)
	res, err := stmnt.Exec(uname, hashedPasswordString, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	
}