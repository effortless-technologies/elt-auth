package server

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/effortless-technologies/elt-auth/models"

	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	//"gopkg.in/mgo.v2/bson"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

var userPayload = `
{
	"username": "test_gec",
	"password": "1234"
}
`

func TestUsers_Login(t *testing.T) {
	Convey("If a test user & a database exists", t, func() {
		models.MongoAddr = mongoAddr
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(
			echo.POST, "/", strings.NewReader(userPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/login")

		Convey("When calling the POST/login handler", func() {
			err := Login(c)
			So(err, ShouldBeNil)

			Convey("Then a .jwt should be returned with a " +
				"a status code of 200, func()", func() {
					So(rec.Code, ShouldEqual, 200)

					type token struct {
						Token 		string			`json:"token"`
					}
					payload, _ := ioutil.ReadAll(rec.Body)
					var t *token
					err = json.Unmarshal([]byte(payload), &t)
					So(err, ShouldBeNil)
			})
		})
	})
}

func TestUsers_GetUsers(t *testing.T) {
	Convey("If a test user & a database exists", t, func() {
		models.MongoAddr = mongoAddr
		So(models.MongoAddr, ShouldNotBeNil)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		So(req, ShouldNotBeNil)

		rec := httptest.NewRecorder()
		So(rec, ShouldNotBeNil)

		c := e.NewContext(req, rec)
		c.SetPath("/users")

		Convey("When calling the GET/users handler", func() {
			err := GetUsers(c)
			So(err, ShouldBeNil)

			Convey("Then the test user should be returned with a " +
				"status code of 200", func() {
				So(rec.Code, ShouldEqual, 200)

				payload, _ := ioutil.ReadAll(rec.Body)
				var properties []*models.User
				err = json.Unmarshal(payload, &properties)
				So(err, ShouldBeNil)
				So(len(properties), ShouldBeGreaterThan, 0)
			})
		})
	})
}
