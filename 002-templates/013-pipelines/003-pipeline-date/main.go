package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var t *template.Template

var fm = template.FuncMap{
	"fDateDMY": dayMonthYear,
}

func init() {
	t = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

// dayMonthYear return format date
// The layout string used by the Parse function and Format method
// shows by example how the reference time should be represented.
// We stress that one must show how the reference time is formatted,
// not a time of the user's choosing. Thus each layout string is a
// representation of the time stamp,
//	Jan 2 15:04:05 2006 MST
// An easy way to remember this value is that it holds, when presented
// in this order, the values (lined up with the elements above):
//	  1 2  3  4  5    6  -7
// There are some wrinkles illustrated below.
func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}
