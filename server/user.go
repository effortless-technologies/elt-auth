package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/effortless-technologies/elt-auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

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

	if up.Password == u.Password {
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
