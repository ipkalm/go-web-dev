package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	var s string

	log.Println(r.Method)

	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		log.Printf("\nfile:\t%v\nheader:\t%v\n", f, h)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form><br>`+s)
}
