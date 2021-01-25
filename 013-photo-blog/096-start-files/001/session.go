package main

import (
	"net/http"
	"strings"

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

func appendValues(w http.ResponseWriter, c *http.Cookie, n string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, n) {
		s += "|" + n
	}
	c.Value = s
	c.MaxAge = 300

	http.SetCookie(w, c)

	return c
}
