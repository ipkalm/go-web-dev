package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.New("").ParseFiles("index.gohtml"))
}

func main() {
	data := struct {
		S1 int
		S2 int
	}{
		77,
		69,
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
