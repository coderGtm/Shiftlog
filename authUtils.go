package main

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type createAccountSuccessResponse struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
}

func checkUsernameExists(uname string) bool {
	stmnt, err := db.Prepare("SELECT * FROM USER WHERE username = ?")
	checkErr(err)

	res, err := stmnt.Exec(uname)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	return affect > 0
}

func registerNewUser(uname string, pswd string) int64 {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pswd),bcrypt.DefaultCost)
	checkErr(err)
	hashedPasswordString := string(hashedPassword)
	stmnt, err := db.Prepare("INSERT INTO USER values(username, password, createdAt, updatedAt) values(?,?,?,?);")
	checkErr(err)
	res, err := stmnt.Exec(uname, hashedPasswordString, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}