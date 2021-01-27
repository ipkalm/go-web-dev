package app

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ipkalm/go-web-dev/015-go-mongo/121-exercice-3/controller"
	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// Run start the server
func Run() {
	r := httprouter.New()
	uc := controller.NewUserController(tpl)

	r.GET("/", uc.Index)
	r.GET("/bar", uc.Bar)
	r.GET("/signup", uc.Signup)
	r.GET("/login", uc.Login)
	r.GET("/logout", uc.Logout)

	r.POST("/signup", uc.Signup)

	log.Fatal(http.ListenAndServe(":8080", r))
}
