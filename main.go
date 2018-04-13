package main

import (
	"flag"
	"net/http"

	"github.com/effortless-technologies/elt-auth/server"
	"github.com/effortless-technologies/elt-auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
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

	flag.Parse()

	models.MongoAddr = mongoAddr

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

