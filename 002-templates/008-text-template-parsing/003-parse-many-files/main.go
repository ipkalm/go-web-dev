package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error while parsing index template", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("error while execute template", err)
	}

	t, err = t.ParseFiles("about.gohtml", "blog.gohtml")
	if err != nil {
		log.Fatalln("error while parsing about and blog template", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "about.gohtml", nil)
	if err != nil {
		log.Fatalln("error while execute about template", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "blog.gohtml", nil)
	if err != nil {
		log.Fatalln("error while execute blog template", err)
	}
}
