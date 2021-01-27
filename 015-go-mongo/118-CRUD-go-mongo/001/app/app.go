package app

import (
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/118-CRUD-go-mongo/001/controller"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

// Run launch app
func Run() {
	uc := controller.NewUserController(getSession())
	r := httprouter.New()

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://root:WTzHXg80bRKUAOQC2mhF@localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	return s
}
