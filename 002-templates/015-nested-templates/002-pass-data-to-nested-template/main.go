package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", "donate")

	if err != nil {
		log.Fatalln(err)
	}
}
