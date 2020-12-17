package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templateStorage, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error while parsing template", err)
	}

	err = templateStorage.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("some error while execute template", err)
	}
}
