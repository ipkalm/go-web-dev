package main

import (
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/115-mvc/001/controller"
	"github.com/julienschmidt/httprouter"
)

type app uint8

func (a app) run() {
	uc := controller.NewUserController()
	r := httprouter.New()

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
