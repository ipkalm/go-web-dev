package main

import (
	"io"
	"net/http"
)

func stdoom(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "StonerDoom")
}

func prog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Progressive")
}

func main() {
	http.HandleFunc("/sd", stdoom)
	http.HandleFunc("/p", prog)

	http.ListenAndServe(":8080", nil)
}
