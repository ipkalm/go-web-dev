package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type dZero struct {
	person
	LicenseToKill bool
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	p1 := dZero{
		person{"Timmyyyy", 9},
		false,
	}
	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
