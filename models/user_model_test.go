package models

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

var mongoAddr = flag.String(
	"mongoAddr",
	"localhost:27017",
	"database service address",
)

var testId *bson.ObjectId

func TestUsers_CreateUsers(t *testing.T) {

	Convey("If users exist", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When retrieving users", func() {
			u := NewUser()
			So(u, ShouldNotBeNil)
			err := u.CreateUser()
			So(err, ShouldBeNil)
			id := u.Id
			So(id, ShouldNotBeNil)
			testId = id
			So(testId, ShouldNotBeNil)

			Convey("A list of users to be returned", func() {
				user, err := FindUserById(testId.Hex())
				So(err, ShouldBeNil)
				So(user, ShouldNotBeNil)
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

func TestUsers_FindUserById(t *testing.T) {

	Convey("If finding a user with username", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When retrieving the user by id", func() {
			user, err := FindUserById(testId.Hex())
			So(err, ShouldBeNil)

			Convey("The user with the specified id should " +
				"be retrieved", func() {
				So(user, ShouldNotBeNil)
				So(user.Id.Hex(), ShouldEqual, testId.Hex())
			})
		})
	})
}

func TestUsersModel_DeleteUser(t *testing.T) {

	Convey("If a properties database exists", t, func() {
		MongoAddr = mongoAddr
		So(MongoAddr, ShouldNotBeNil)

		Convey("When deleting an existing property", func() {
			err := DeleteUser(testId.Hex())
			So(err, ShouldBeNil)

			Convey("A property should have been deleted", func() {
				user, err := FindUserById(testId.Hex())
				So(err, ShouldNotBeNil)
				So(user, ShouldBeNil)
			})
		})
	})
}
