package main

import (
	"net/http"

	"github.com/google/uuid"
)

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:   "session",
			Value:  uuid.New().String(),
			MaxAge: 300,
		}
		http.SetCookie(w, c)
	}
	return c
}
