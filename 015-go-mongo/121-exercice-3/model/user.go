package model

// User provide user model
type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}
