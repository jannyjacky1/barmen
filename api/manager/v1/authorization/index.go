package authorization

import (
	"database/sql"
	"github.com/jannyjacky1/barmen/api/manager/v1/common"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
)

const tableName = "tbl_admins"

func Auth(c echo.Context) error {

	model := new(common.AuthForm)

	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	if err := c.Validate(model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	db := c.Get("db").(*sql.DB)
	var pwdHash string

	err := db.QueryRow("SELECT password_hash FROM "+tableName+" WHERE username=$1", model.Username).Scan(&pwdHash)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Message{Message: "Incorrect login or password"})
	}

	if pwdHash == "" || !passwordValidate(model.Password, pwdHash) {
		return c.JSON(http.StatusBadRequest, common.Message{Message: "Incorrect login or password"})
	}

	token := generateToken()
	tokenExpire := time.Now().Add(48 * time.Hour)

	db.Query("UPDATE tbl_admins SET token=$1, token_expire=$2 WHERE username=$3", token, tokenExpire, model.Username)

	return c.JSON(http.StatusOK, common.AuthData{
		Token:       token,
		TokenExpire: tokenExpire,
	})

}

func passwordValidate(pwd string, hashPwd string) bool {

	bytePwd := []byte(pwd)
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}

	return true

}

func generateToken() string {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	token := string(bytes)

	return token
}

func TokenValidator(s string, c echo.Context) (bool, error) {

	database := c.Get("db").(*sql.DB)
	var id int

	err := database.QueryRow("SELECT id FROM tbl_admins WHERE token=$1 AND token_expire>$2", s, time.Now().Format("2006-01-02 15:04:05.999999999")).Scan(&id)

	return id > 0, err
}
