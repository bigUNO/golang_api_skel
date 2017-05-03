package main

import (
	"log"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

// getSession: create a mongo session and pass the dutch
func getSession() *mgo.Session {
  // read config in
	var config = ReadConfig()
	fmt.Println(config.MongoDBHosts)

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: config.AuthDatabase,
		Username: config.AuthUserName,
		Password: config.AuthPassword,
	}

	// create mongo session which manages pool of socket connections
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)

	if config.IsDrop {
		err = session.DB("inventory").DropDatabase()
		if err != nil {
			panic(err)
		}
	}
	return session
}

// ensureIndex: create index then starting up
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

func init() {
	// create index on start
	s := getSession()
	ensureIndex(s)
}
