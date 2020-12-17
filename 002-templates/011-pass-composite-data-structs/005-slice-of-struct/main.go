package main

import (
	"log"
	"os"
	"text/template"
)

type band struct {
	Name  string
	Genre string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	data := []band{
		{"Immortal", "Black Metal"},
		{"Sleep", "Stoner Doom"},
		{"Led Zeppelin", "Hard Rock"},
		{"Meshuggah", "Djent"},
	}
	err := t.Execute(os.Stdout, data)

	if err != nil {
		log.Println("error while execute teamplate", err)
	}
}
