package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.POST("/", index)
	log.Fatalln(http.ListenAndServe("127.0.0.1:8080", r))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := getCookie(w, r)

	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("photo")
		if err != nil {
			log.Panicln(err)
		}
		defer f.Close()

		ext := strings.Split(h.Filename, ".")[1]
		hash := sha1.New()
		_, err = io.Copy(hash, f)
		if err != nil {
			log.Panicln(err)
		}

		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		wd, err := os.Getwd()
		if err != nil {
			log.Panicln(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)

		nf, err := os.Create(path)
		f.Seek(0, 0)
		_, err = io.Copy(nf, f)
		if err != nil {
			log.Panicln(err)
		}

		c = appendValues(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	t.ExecuteTemplate(w, "index.gohtml", xs)
}
