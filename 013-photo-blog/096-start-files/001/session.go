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

func appendValues(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	p1 := "p1.jpg"
	p2 := "p2.jpg"
	p3 := "p3.jpg"

	s := c.Value
	switch {
	case !strings.Contains(s, p1):
		s += "|" + p1
	case !strings.Contains(s, p2):
		s += "|" + p2
	case !strings.Contains(s, p3):
		s += "|" + p3
	}
	c.Value = s
	c.MaxAge = 300

	http.SetCookie(w, c)

	return c
}
