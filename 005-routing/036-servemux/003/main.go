package main

import (
	"io"
	"net/http"
)

type stdoom bool
type prog bool

func (m stdoom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "StonerDoom")
}

func (m prog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Progressive")
}

func main() {
	var sd stdoom
	var p prog

	http.Handle("/sd", sd)
	http.Handle("/p", p)

	http.ListenAndServe(":8080", nil)
}
