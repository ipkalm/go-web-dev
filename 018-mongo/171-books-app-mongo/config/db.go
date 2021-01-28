package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	// DB database
	DB *mgo.Database
	// Books mongo collection
	Books *mgo.Collection
)

func init() {
	s, err := mgo.Dial("mongodb://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}
	if err = s.Ping(); err != nil {
		log.Panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	log.Println("connected to mongodb")
}
