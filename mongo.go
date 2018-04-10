package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserInfo struct {
	Email string
	Token string
}

func connectToDb() (*mgo.Session, error) {
	sess, err := mgo.Dial("localhost:27017")
	sess.SetMode(mgo.Monotonic, true)
	return sess, err
}
func setDB(db string, sess *mgo.Session) *mgo.Database {
	database := sess.DB(db)
	return database
}
func insertToCollection(user UserInfo, db *mgo.Database, collection string) error {
	fmt.Println("Heree", user.Email)
	c := db.C(collection)
	err := c.Insert(&UserInfo{
		user.Email,
		user.Token,
	})
	fmt.Println("Errored in insertting collectiong")
	return err
}

func findFromCollection(email string, c *mgo.Collection) (UserInfo, error) {
	user := UserInfo{}
	err := c.Find(bson.M{"email": email}).One(&user)
	return user, err
}
