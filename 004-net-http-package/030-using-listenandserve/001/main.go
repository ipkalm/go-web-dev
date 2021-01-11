package main

import (
	"fmt"
	"net/http"
)

type pewpew int

func (p pewpew) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "just test how its work")
}

func main() {
	var p pewpew
	http.ListenAndServe(":8080", p)
}
