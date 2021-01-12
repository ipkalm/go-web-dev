package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog", dog)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		ErrorHandler(w, err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "me.gohtml", "ipkalm")
	if err != nil {
		ErrorHandler(w, err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		ErrorHandler(w, err)
	}
}

// ErrorHandler take response writer and error and print error
// to logs and to writer
func ErrorHandler(w http.ResponseWriter, e error) {
	err := e.Error()
	fmt.Fprintln(w, err)
	log.Panicln(err)
}
