package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/mrshl", mrshl)
	r.GET("/encd", encd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.Execute(w, nil)
}

func mrshl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		"j",
		"b",
		[]string{"s", "g", "w s o h"},
	}
	json, err := json.Marshal(p)
	if err != nil {
		log.Panicln(err)
	}
	w.Write(json)
}

func encd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		"j",
		"b",
		[]string{"s", "g", "w s o h"},
	}
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panicln(err)
	}
}
