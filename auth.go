package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type createAccountSuccessResponse struct {
	Username  string `json:"username"`
	AuthToken string `json:"authToken"`
}
type loginSuccessResponse struct {
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

func registerNewUser(uname string, pswd string) string {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	hashedPasswordString := hashPasswordString(pswd)
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
	return authToken
}

func hashPasswordString(s string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	checkErr(err)
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString
}

func createJWT(userId int64) string {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token_lifespan, err := strconv.Atoi(os.Getenv("JWT_HOUR_LIFESPAN"))
	checkErr(err)

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	checkErr(err)
	return tokenString
}

func verifyAndLogin(uname string, pswd string) (bool, string) {
	var hashedPassword string
	var userId int64
	err := db.QueryRow("SELECT id, password from USER WHERE username = ?;", uname).Scan(&userId, &hashedPassword)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened!
			checkErr(err)
		}
		// record does not exist
		return false, ""
	}
	// record exists
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pswd)) == nil {
		// password correct
		currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
		authToken := createJWT(userId)
		stmnt, err := db.Prepare("UPDATE USER SET authToken = ?, updatedAt = ? WHERE id = ?;")
		checkErr(err)
		_, err = stmnt.Exec(authToken, currentTimeStamp, userId)
		checkErr(err)
		stmnt.Close()
		return true, authToken
	}
	return false, ""
}

func extractAuthToken(c *gin.Context) string {
	token := c.Query("authToken")
	if token != "" { return token }
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func extractUserIdFromToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func isTokenValid(token string) (uint, bool) {
	// extract user id from token and check if token matches the one in db. Returns user id also
	userId, validationErr := extractUserIdFromToken(token)
	if validationErr == nil {
		var dbToken string
		err := db.QueryRow("SELECT authToken from USER WHERE id = ?;", userId).Scan(&dbToken)
		if err != nil {
			if err != sql.ErrNoRows {
				// a real error happened!
				checkErr(err)
			}
			// record does not exist
			return 0, false
		}
		// record exists
		if token == dbToken {
			return userId, true
		}
	}
	return 0, false
}

func logoutUser(userId uint) {
	currentTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	stmnt, err := db.Prepare("UPDATE USER SET authToken = NULL, updatedAt = ? WHERE id = ?")
	checkErr(err)
	_, err = stmnt.Exec(currentTimeStamp, userId)
	checkErr(err)
}

func deleteUserAccount(userId uint) {
	stmnt, err := db.Prepare("DELETE FROM user WHERE id = ?")
	checkErr(err)
	_, err = stmnt.Exec(userId)
	checkErr(err)
}