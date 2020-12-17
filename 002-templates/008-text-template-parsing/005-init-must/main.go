package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	// this execute first template in *t
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("error while execute index template", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error while execute index template", err)
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
