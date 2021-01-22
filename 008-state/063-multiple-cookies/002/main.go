package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "1",
		})
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if s, err := strconv.Atoi(c.Value); err == nil {
		c.Value = strconv.Itoa(s + 1)
		http.SetCookie(w, c)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "you open this site: ", c.Value)
}
