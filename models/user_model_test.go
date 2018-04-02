package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsers_GetUsers(t *testing.T) {

	Convey("If users exist", t, func() {

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
		username := "gec"
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
