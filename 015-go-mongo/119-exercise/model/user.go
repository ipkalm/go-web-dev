package model

// User struct
type User struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
}

// Id was of type string before

// UserDB users database
type UserDB = map[string]User
