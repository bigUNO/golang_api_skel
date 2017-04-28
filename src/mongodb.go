package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

const (
	MongoDBHosts = "ds159050.mlab.com:59050"
	AuthDatabase = "golang_api"
	AuthUserName = "zA4gAR8ETb78"
	AuthPassword = "sFZe4JphfWyhH5pASsrARjmcXV6H9ZYJ"
	IsDrop       = false
)

func init() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// create mongo session which manages pool of socket connections
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	if IsDrop {
		err = session.DB("inventory").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	ensureIndex(session)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("golang_api").C("inventory")

	index := mgo.Index{
		Key:        []string{"Id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
