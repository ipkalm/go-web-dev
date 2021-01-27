package main

import (
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/120-exercice-2/controller"
	"github.com/ipkalm/go-web-dev/015-go-mongo/120-exercice-2/model"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controller.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() (model.UserDB, string) {
	return make(model.UserDB), "db.json"
}
