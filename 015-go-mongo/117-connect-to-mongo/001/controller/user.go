package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/117-connect-to-mongo/001/model"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// UserController provide controller for User struct
type UserController struct {
	client *mgo.Session
}

// NewUserController return pointer to UserController struct
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser print user by id in url
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "j",
		Gender: "m",
		Age:    32,
		ID:     p.ByName("id"),
	}
	ujs, err := json.Marshal(u)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", ujs)
}

// CreateUser create User by POST request and store
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Panicln(err)
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser just delete the user by id in url from db
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: code to del user
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "code to del user")
}
