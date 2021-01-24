package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type user struct {
	Nick     string
	Fname    string
	Lname    string
	Password string
}

var t *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/signup", signup)
	r.GET("/bar", bar)
	r.GET("/admin", admin)

	r.POST("/signup", signup)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUser(r)
	err := t.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Panicln(err)
	}
}

func bar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUser(r)

	if !loggedIn(r) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	err := t.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		log.Panicln(err)
	}
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if loggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		nick := r.FormValue("nick")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		pwd := r.FormValue("pwd")

		if _, ok := dbUsers[nick]; ok {
			http.Error(w, "nick already taken", http.StatusForbidden)
			return
		}

		sID := uuid.New().String()
		c := &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		dbSessions[sID] = nick
		http.SetCookie(w, c)

		dbUsers[nick] = user{nick, fname, lname, pwd}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := t.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Panicln(err)
	}
}

func admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "admin.gohtml", dbSessions)
	if err != nil {
		log.Panicln(err)
	}
}

func loggedIn(r *http.Request) bool {
	var tmp bool

	if c, err := r.Cookie("session"); err != nil {
		tmp = false
	} else if _, ok := dbSessions[c.Value]; ok {
		tmp = true
	}

	return tmp
}

func getUser(r *http.Request) user {
	var u user

	if c, err := r.Cookie("session"); err != nil {
		u = user{}
	} else if uID, ok := dbSessions[c.Value]; ok {
		u = dbUsers[uID]
	}

	return u
}
