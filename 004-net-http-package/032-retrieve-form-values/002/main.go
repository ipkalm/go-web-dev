package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var t *template.Template

type tmp int

func (m tmp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		Subm   url.Values
	}{
		Method: r.Method,
		Subm:   r.Form,
	}

	err = t.ExecuteTemplate(w, "index.gohtml", data)
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
