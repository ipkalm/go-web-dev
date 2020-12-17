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

type concert struct {
	City string
	Date string
}

type tour struct {
	Bands    []band
	Concerts []concert
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	bands := []band{
		{"Immortal", "Black Metal"},
		{"Sleep", "Stoner Doom"},
		{"Led Zeppelin", "Hard Rock"},
		{"Meshuggah", "Djent"},
	}

	concerts := []concert{
		{"Oslo", "12.12.2020"},
		{"Moscow", "14.12.2020"},
		{"Toronto", "20.12.2020"},
	}

	tourne := tour{
		Bands:    bands,
		Concerts: concerts,
	}

	err := t.Execute(os.Stdout, tourne)

	if err != nil {
		log.Println("error while execute teamplate", err)
	}
}
