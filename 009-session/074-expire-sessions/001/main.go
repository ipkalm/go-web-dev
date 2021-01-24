package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Nick     string
	Fname    string
	Lname    string
	Password []byte
	Role     string
}

type session struct {
	nick         string
	lastActivity time.Time
}

var t *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength uint8 = 30

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/bar", bar)
	r.GET("/signup", signup)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/admin", admin)

	r.POST("/signup", signup)
	r.POST("/login", login)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUser(w, r)
	err := t.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Panicln(err)
	}
}

func bar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUser(w, r)

	if !loggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "you are not 007. access denied", http.StatusForbidden)
		return
	}

	err := t.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		log.Panicln(err)
	}
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if loggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		nick := r.FormValue("nick")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		passwd, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("passwd")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		role := r.FormValue("roles")

		if _, ok := dbUsers[nick]; ok {
			http.Error(w, "nick already taken", http.StatusForbidden)
			return
		}

		sID := uuid.New().String()
		c := &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		dbSessions[sID] = session{nick, time.Now()}
		http.SetCookie(w, c)

		dbUsers[nick] = user{nick, fname, lname, passwd, role}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := t.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Panicln(err)
	}
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if loggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	switch r.Method {
	case http.MethodPost:
		n := r.FormValue("nick")
		p := r.FormValue("passwd")
		u, ok := dbUsers[n]
		if !ok {
			http.Error(w, "username or passwd do not match", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "username or passwd do not match", http.StatusForbidden)
			return
		}
		sID := uuid.New().String()
		c := &http.Cookie{
			Name:   "session",
			Value:  sID,
			MaxAge: int(sessionLength),
		}
		dbSessions[sID] = session{n, time.Now()}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	case http.MethodGet:
		err := t.ExecuteTemplate(w, "login.gohtml", nil)
		if err != nil {
			log.Panicln(err)
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !loggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, err := r.Cookie("session")
	if err != nil {
		log.Panicln(err)
	}

	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !loggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	u := getUser(w, r)

	if u.Role != "admin" {
		http.Error(w, "you are not admin. access denied", http.StatusForbidden)
		return
	}

	err := t.ExecuteTemplate(w, "admin.gohtml", dbSessions)
	if err != nil {
		log.Panicln(err)
	}
}
