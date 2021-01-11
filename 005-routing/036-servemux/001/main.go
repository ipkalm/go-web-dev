package main

import (
	"io"
	"net/http"
)

type metal int

func (m metal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/progressive":
		io.WriteString(w, "<a href=\"https://open.spotify.com/playlist/37i9dQZF1DX5wgKYQVRARv\">prog</a>")
	case "/stoner-doom":
		io.WriteString(w, "<a href=\"https://open.spotify.com/playlist/5yaqdcihcfLkaIxcR2STSi\">stoner doom</a>")
	}
}

func main() {
	var m metal
	http.ListenAndServe(":8080", m)
}
