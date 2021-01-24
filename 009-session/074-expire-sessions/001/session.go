package main

import (
	"log"
	"net/http"
	"time"
)

func loggedIn(w http.ResponseWriter, r *http.Request) bool {
	var tmp bool

	if c, err := r.Cookie("session"); err != nil {
		tmp = false
	} else if _, ok := dbSessions[c.Value]; ok {
		c.MaxAge = int(sessionLength)
		http.SetCookie(w, c)
		tmp = true
	}

	return tmp
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	var u user

	if c, err := r.Cookie("session"); err != nil {
		u = user{}
	} else if uID, ok := dbSessions[c.Value]; ok {
		c.MaxAge = int(sessionLength)
		http.SetCookie(w, c)
		u = dbUsers[uID.nick]
	}

	return u
}

func cleanSessions() {
	log.Println("before clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
}

func showSessions() {
	log.Println("***")
	for k, v := range dbSessions {
		log.Println(k, v.nick)
	}
	log.Println("")
}
