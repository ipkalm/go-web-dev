package main

import (
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

type person struct {
	FName string
	LName string
	Sub   bool
}

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fn := r.FormValue("first")
	ln := r.FormValue("last")
	sb := r.FormValue("subscribe") == "on"

	err := t.ExecuteTemplate(w, "index.gohtml", person{fn, ln, sb})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
