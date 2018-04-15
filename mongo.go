package main

import (
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
func insertToCollection(user UserInfo, db *mgo.Database, collection string) (*mgo.ChangeInfo, error) {
	c := db.C(collection)
	m := make(map[string]string, 1)
	m["email"] = user.Email
	info, errInUpsert := c.Upsert(m, &UserInfo{
		user.Email,
		user.Token,
	})
	return info, errInUpsert
}

func findFromCollection(email string, c *mgo.Collection) (UserInfo, error) {
	user := UserInfo{}
	err := c.Find(bson.M{"email": email}).One(&user)
	return user, err
}
