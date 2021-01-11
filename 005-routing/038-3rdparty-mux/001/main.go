package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mx := httprouter.New()
	mx.GET("/", index)
	mx.GET("/about", about)
	mx.GET("/apply", apply)
	mx.POST("/apply", apllyProc)
	mx.GET("/blog/:cat/:post", blogGetPost)
	mx.GET("/blog", blogGet)

	http.ListenAndServe(":8080", mx)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func apllyProc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Panicln(err)
	}

	err = t.ExecuteTemplate(w, "applyProc.gohtml", r.PostForm)
	HandleError(w, err)
}

func blogGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := t.ExecuteTemplate(w, "blog.gohtml", "list of articles")
	HandleError(w, err)
}

func blogGetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := t.ExecuteTemplate(w, "blogPost.gohtml", ps)
	HandleError(w, err)
}

// HandleError write error to log and to ResponseWriter
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panicln(err)
	}
}
