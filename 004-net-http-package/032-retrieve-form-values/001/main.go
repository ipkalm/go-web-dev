package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

type tmp int

func (m tmp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(r.Form)

	err = t.ExecuteTemplate(w, "index.gohtml", r.Form)
	if err != nil {
		log.Panicln(err)
	}
}

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t1 tmp
	err := http.ListenAndServe(":8080", t1)
	if err != nil {
		log.Fatalln(err)
	}
}
