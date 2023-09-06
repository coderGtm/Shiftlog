package main

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type createAccountSuccessResponse struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	AuthToken string `json:"authToken"`
}

func checkUsernameExists(uname string) bool {
	sqlStmt := `SELECT username FROM USER WHERE username = ?`
    err := db.QueryRow(sqlStmt, uname).Scan(&uname)
    if err != nil {
        if err != sql.ErrNoRows {
            // a real error happened!
			checkErr(err)
        }

        return false
    }

    return true
}

func registerNewUser(uname string, pswd string) (int64, string) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pswd),bcrypt.DefaultCost)
	checkErr(err)
	hashedPasswordString := string(hashedPassword)
	stmnt, err := db.Prepare("INSERT INTO USER(username, password, createdAt, updatedAt) values(?,?,?,?);")
	checkErr(err)
	res, err := stmnt.Exec(uname, hashedPasswordString, currentTimeStamp, currentTimeStamp)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	authToken := createJWT(id)
	stmnt, err = db.Prepare("UPDATE USER SET authToken = ? WHERE id = ?;")
	checkErr(err)
	_, err = stmnt.Exec(authToken, id)
	checkErr(err)
	return id, authToken
}

func createJWT(userId int64) string {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token_lifespan,err := strconv.Atoi(os.Getenv("JWT_HOUR_LIFESPAN"))
	checkErr(err)

	claims := jwt.MapClaims {
		"userId": userId,
		"exp": time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	checkErr(err)
	return tokenString
}