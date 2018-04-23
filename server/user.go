package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/effortless-technologies/elt-auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {

	type userPayload struct {
		Username 		string 			`json:"username"`
		Password 		string			`json:"password"`
	}

	up := new(userPayload)
	if err := c.Bind(up); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hp, err := hashPassword(up.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	up.Username = strings.ToLower(up.Username)

	u := models.NewUser()
	u.Username = up.Username
	u.Password = hp
	if err := u.CreateUser(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}

func DeleteUser(c echo.Context) error {

	id := c.Param("id")

	if err := models.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func GetUsers(c echo.Context) error {

	u, err := models.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, u)
}

func Login(c echo.Context) error {

	type userPayload struct {
		Username 		string 			`json:"username"`
		Password 		string			`json:"password"`
	}

	up := new(userPayload)
	if err := c.Bind(up); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	up.Username = strings.ToLower(up.Username)

	u, err := models.FindUser(up.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	match := checkPasswordHash(up.Password, u.Password)

	if match == true {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Name
		claims["franchise_id"] = u.FranchiseId
		claims["role"] = u.Role
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
