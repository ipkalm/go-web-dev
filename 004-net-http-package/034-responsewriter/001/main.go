package main

import (
	"fmt"
	"log"
	"net/http"
)

type tmp int

func (m tmp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ipkalm-key", "this is from ipkalm")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "<h1>Whatever you want here</h1>")
}

func main() {
	var t1 tmp
	err := http.ListenAndServe(":8080", t1)
	if err != nil {
		log.Fatalln(err)
	}
}
