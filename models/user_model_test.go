package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"flag"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

func TestUsers_CreateUsers(t *testing.T) {

	Convey("If users exist", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When retrieving users", func() {
			u := NewUser()
			So(u, ShouldNotBeNil)
			err := u.CreateUser()
			So(err, ShouldBeNil)
			id := u.Id.Hex()
			So(id, ShouldNotBeNil)

			Convey("A list of users to be returned", func() {
				users, err := GetUsers()
				So(err, ShouldBeNil)
				So(users, ShouldNotBeNil)

				found := false
				for index := range users {
					if users[index].Id.Hex() == id {
						found = true
					}
 				}

 				So(found, ShouldEqual, true)
			})
		})
	})
}

func TestUsers_GetUsers(t *testing.T) {

	Convey("If users exist", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When retrieving users", func() {
			users, err := GetUsers()
			So(err, ShouldBeNil)

			Convey("A list of users to be returned", func() {
				So(users, ShouldNotBeNil)
				So(len(users), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestUsers_FindUser(t *testing.T) {

	Convey("If finding a user with username", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		username := "test_gec"
		So(username, ShouldNotBeNil)

		Convey("When retrieving the user by uesrname", func() {
			user, err := FindUser(username)
			So(err, ShouldBeNil)

			Convey("The user with the specified username should " +
				"be retrieved", func() {
				So(user, ShouldNotBeNil)
				So(user.Username, ShouldEqual, username)
			})
		})
	})
}
