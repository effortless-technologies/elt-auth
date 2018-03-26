package main

import (
	"net/http"

	"github.com/effortless-technologies/elt-auth/server"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func accessible(c echo.Context) error {

	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	//franchiseId := claims["franchise_id"].(float64)
	role := claims["role"].(string)
	return c.String(http.StatusOK, ""+name+" "+role)
}

func main() {

	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", server.Login)

	e.GET("/", accessible)
	e.GET("/users", server.GetUsers)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":7000"))
}

