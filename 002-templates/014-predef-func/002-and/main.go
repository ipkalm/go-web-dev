package main

import (
	"log"
	"os"
	"text/template"
)

var t *template.Template

type user struct {
	Nickname string
	Email    string
	Admin    bool
}

func init() {
	t = template.Must(template.New("").ParseFiles("index.gohtml"))
}

func main() {
	users := []user{
		{
			Nickname: "Arnold",
			Email:    "arni@gmail.com",
			Admin:    false,
		},
		{
			Nickname: "bro1337",
			Email:    "1337@bro.com",
			Admin:    true,
		},
		{
			Nickname: "",
			Email:    "gopher@great.com",
			Admin:    false,
		},
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
