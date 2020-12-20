package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var t *template.Template

func init() {
	t = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Intro to Golang", "4"},
				{"CSCI-130", "Intro web dev with Golang", "4"},
				{"CSCI-140", "Mobile dev with Golang", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "advanced go", "5"},
				{"CSCI-190", "advanced webdev with go", "5"},
				{"CSCI-191", "advanced mobile dev with go", "5"},
			},
		},
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}
