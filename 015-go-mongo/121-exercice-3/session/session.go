package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ipkalm/go-web-dev/015-go-mongo/121-exercice-3/model"
)

// SessionLength session lenght
const SessionLength int = 30

var (
	// DbUsers users database
	DbUsers = map[string]model.User{}
	// DbSessions sessions database
	DbSessions = map[string]model.Session{}
	// DbSessionsCleaned DbSessionsCleaned
	DbSessionsCleaned time.Time
)

// GetUser return model.User
func GetUser(w http.ResponseWriter, req *http.Request) model.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New().String()
		c = &http.Cookie{
			Name:  "session",
			Value: sID,
		}

	}
	c.MaxAge = SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u model.User
	if s, ok := DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		DbSessions[c.Value] = s
		u = DbUsers[s.Un]
	}
	return u
}

// AlreadyLoggedIn bool
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		DbSessions[c.Value] = s
	}
	_, ok = DbUsers[s.Un]
	// refresh session
	c.MaxAge = SessionLength
	http.SetCookie(w, c)
	return ok
}

// CleanSessions clean sessions
func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// ShowSessions for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
