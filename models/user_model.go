package models

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var MongoAddr *string

type User struct {
	Id 				*bson.ObjectId	`json:"id" bson:"_id"`
	Username 		string 			`json:"username"`
	Password		string			`json:"password"`
	Role 			string 			`json:"role"`
	FranchiseId		int 			`json:"franchise_id"`
	Name 			string			`json:"name"`
}

func NewUser() *User {

	u := new(User)
	id := bson.NewObjectId()
	u.Id = &id

	return u
}

func (u *User) CreateUser() error {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("users")
	_, err = c.UpsertId(u.Id, u)
	if err != nil {
		log.Println("Error creating User: ", err.Error())
		return err
	}

	return nil
}

func GetUsers() ([]*User, error) {

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

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

	session, err := mgo.Dial(*MongoAddr)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("elt").C("users")
	user := User{}
	err = c.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
