package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error while parsing template", err)
	}

	f, err := os.Create("./index.html")
	if err != nil {
		log.Println("error while creating file", err)
	}
	defer f.Close()

	err = t.Execute(f, nil)
	if err != nil {
		log.Fatalln("error while execute template", err)
	}
}
