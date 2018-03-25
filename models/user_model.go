package models

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username 		string 			`json:"username"`
	Password		string			`json:"password"`
	Role 			string 			`json:"role"`
	FranchiseId		int 			`json:"franchise_id"`
}

func GetUsers() ([]*User, error) {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("users")
	var Users []*User
	err = c.Find(bson.M{}).All(&Users)
	if err != nil {
		return nil, err
	}

	return Users, nil
}

func FindUser(username string) (*User, error) {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("users")
	user := User{}
	err = c.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
