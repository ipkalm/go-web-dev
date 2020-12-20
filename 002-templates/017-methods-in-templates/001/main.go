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

func (p person) SomeProc() int {
	return 42
}

func (p person) TakesArgs(x int) int {
	return x * 2
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	p := person{
		Name: "Rocky",
		Age:  45,
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
