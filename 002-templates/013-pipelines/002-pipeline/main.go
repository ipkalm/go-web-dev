package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var t *template.Template

var fm = template.FuncMap{
	"fDbl":  double,
	"fSq":   square,
	"fSqrt": sqRoot,
}

func init() {
	t = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {
	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}

func double(x int) int {
	return x * 2
}

func square(x int) float64 {
	return float64(x * x)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}
