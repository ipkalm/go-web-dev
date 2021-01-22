package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/batch", batch)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "my-val",
	})

	fmt.Fprintln(w, "cookie written, check your browser")
	fmt.Fprintln(w, "open dev tools -> app -> cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "your cookie #1: ", c1)

	c2, err := r.Cookie("my-1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "your cookie #2: ", c2)

	c3, err := r.Cookie("my-2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "your cookie #3: ", c3)
}

func batch(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-1",
		Value: "my-val1",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "my-2",
		Value: "my-val2",
	})
}
