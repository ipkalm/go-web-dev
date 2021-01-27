package controller

import (
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/google/uuid"
	"github.com/ipkalm/go-web-dev/015-go-mongo/121-exercice-3/model"
	"github.com/ipkalm/go-web-dev/015-go-mongo/121-exercice-3/session"
	"golang.org/x/crypto/bcrypt"
)

// UserController provide controller for user model
type UserController struct {
	tpl *template.Template
}

// NewUserController return UserController instance
func NewUserController(t *template.Template) *UserController {
	return &UserController{t}
}

// Index do index
func (uc UserController) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := session.GetUser(w, req)
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// Bar do bar
func (uc UserController) Bar(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// Signup do signup
func (uc UserController) Signup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u model.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := session.DbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.New().String()
		c := &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		session.DbSessions[c.Value] = model.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = model.User{un, bs, f, l, r}
		session.DbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

// Login do login
func (uc UserController) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u model.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := session.DbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.New().String()
		c := &http.Cookie{
			Name:  "session",
			Value: sID,
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		session.DbSessions[c.Value] = model.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.ShowSessions() // for demonstration purposes
	uc.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

// Logout do logout
func (uc UserController) Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(session.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(session.DbSessionsCleaned) > (time.Second * 30) {
		go session.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
