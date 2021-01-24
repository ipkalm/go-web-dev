package main

import "net/http"

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
