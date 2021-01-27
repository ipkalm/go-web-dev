package model

import (
	"gopkg.in/mgo.v2/bson"
)

// User provide user model for json
type User struct {
	Name   string        `json:"name"`
	Gender string        `json:"gender"`
	Age    uint8         `json:"age"`
	ID     bson.ObjectId `json:"id"`
}
