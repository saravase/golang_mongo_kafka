package db

import (
	"golang_mongo_kafka/data"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "primz2207"
	username   = ""
	password   = ""
	collection = "plants"
)

func NewMongoStore() *data.MongoStore {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		os.Exit(1)
	}

	return &data.MongoStore{
		Session:    session,
		Host:       hosts,
		Database:   database,
		Username:   username,
		Password:   password,
		Collection: collection,
	}
}
