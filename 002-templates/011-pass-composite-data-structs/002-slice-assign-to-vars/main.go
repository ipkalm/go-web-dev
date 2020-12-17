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
	data := []string{"Mayhem", "1349", "Darkthrone", "Burzum"}
	err := t.Execute(os.Stdout, data)

	if err != nil {
		log.Println("error while execute teamplate", err)
	}
}
