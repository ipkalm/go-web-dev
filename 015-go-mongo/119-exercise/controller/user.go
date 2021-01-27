package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/ipkalm/go-web-dev/015-go-mongo/119-exercise/model"
	"github.com/julienschmidt/httprouter"
)

// UserController struct
type UserController struct {
	db model.UserDB
}

// NewUserController retrun UserController struct
func NewUserController(db model.UserDB) *UserController {
	return &UserController{db}
}

// GetUser get User from db
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// composite literal
	u := model.User{}

	// Fetch user
	u = uc.db[id]

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser in db
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = uuid.New().String()
	uc.db[u.ID] = u

	uj, err := json.Marshal(u)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
	log.Println(uc.db)
}

// DeleteUser from db
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.db, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user: ", id, "\n")
}
