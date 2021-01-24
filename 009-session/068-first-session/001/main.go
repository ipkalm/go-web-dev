package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type user struct {
	Nick  string
	Fname string
	Lname string
}

var t *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/", postForm)
	router.GET("/bar", bar)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.New().String()
		c = &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		http.SetCookie(w, c)
	}

	if _, ok := dbUsers[dbSessions[c.Value]]; ok {
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
		return
	}

	t.ExecuteTemplate(w, "index.gohtml", nil)
}

func postForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nick := r.FormValue("nick")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	u := user{nick, fname, lname}

	dbSessions[c.Value] = nick
	dbUsers[nick] = u

	t.ExecuteTemplate(w, "enterance.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nick, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[nick]

	t.ExecuteTemplate(w, "bar.gohtml", u)
}
