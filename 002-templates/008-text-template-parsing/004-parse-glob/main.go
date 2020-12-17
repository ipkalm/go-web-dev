package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.ParseGlob("./*.gohtml")
	if err != nil {
		log.Fatalln("error while glob parsing templates", err)
	}

	// this execute first template in *t
	err = t.Execute(os.Stdout, nil)
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
