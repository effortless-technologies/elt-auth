package main

import (
	"flag"

	"github.com/effortless-technologies/elt-auth/server"
	"github.com/effortless-technologies/elt-auth/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

func main() {

	flag.Parse()

	models.MongoAddr = mongoAddr

	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", server.Login)
	e.POST("/users", server.CreateUser)
	e.GET("/users", server.GetUsers)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	e.Logger.Fatal(e.Start(":7000"))
}

