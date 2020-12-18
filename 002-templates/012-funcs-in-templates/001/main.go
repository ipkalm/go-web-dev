package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var t *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type band struct {
	Name  string
	Genre string
}

type concert struct {
	City string
	Date string
}

func init() {
	t = template.Must(template.New("").Funcs(fm).ParseFiles("./index.gohtml"))
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

	tourne := struct {
		Bands    []band
		Concerts []concert
	}{
		bands,
		concerts,
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", tourne)

	if err != nil {
		log.Println("error while execute teamplate", err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}
