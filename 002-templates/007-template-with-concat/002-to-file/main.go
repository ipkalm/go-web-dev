package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	x := "1349"

	t := `<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="utf-8">
	  <title>Greetings</title>
	</head>
	<body>
	  <h1> The band is: ` + x + `</h1>
	</body>
	</html>
	`

	f, err := os.Create("./index.html")
	if err != nil {
		log.Fatalln("error while creating file", err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(t))
}
