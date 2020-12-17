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
	data := map[string]string{
		"Mayhem":     "Oslo",
		"1349":       "Oslo",
		"Darkthrone": "Kolbotn",
		"Burzum":     "Bergen",
	}
	err := t.Execute(os.Stdout, data)

	if err != nil {
		log.Println("error while execute teamplate", err)
	}
}
