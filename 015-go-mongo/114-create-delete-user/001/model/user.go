package model

// User provide user model for json
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    uint8  `json:"age"`
	ID     string `json:"id"`
}
