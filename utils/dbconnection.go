package utils

import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

func Connect(uri string, dbName string) (*mgo.Database, error) {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(dbName), nil
}

