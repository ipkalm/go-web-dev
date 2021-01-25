package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	log.Fatalln(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.ExecuteTemplate(w, "index.gohtml", getCookie(w, r).Value)
}
