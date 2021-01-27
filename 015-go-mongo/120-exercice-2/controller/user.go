package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/ipkalm/go-web-dev/015-go-mongo/120-exercice-2/model"
	"github.com/julienschmidt/httprouter"
)

// UserController struct
type UserController struct {
	db    model.UserDB
	fname string
}

// NewUserController retrun UserController struct
func NewUserController(db model.UserDB, fname string) *UserController {
	// open or create file if not exist
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db file open/create")

	// decode data from file to map
	var tmp []byte

	tmp, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if json.Valid(tmp) {
		log.Println("data valid")
		err = json.Unmarshal(tmp, &db)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("unmarshal complete")
	}

	log.Println("read from db file to map complete")

	// close file
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("close file")

	// return controller
	return &UserController{db, fname}
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

	exportDBToFile(uc.db, uc.fname)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser from db
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.db, id)
	fmt.Println(uc.db)

	exportDBToFile(uc.db, uc.fname)

	w.WriteHeader(http.StatusOK) // 200
}

func exportDBToFile(db model.UserDB, fname string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(db)
	if err != nil {
		log.Panic(err)
	}
}
