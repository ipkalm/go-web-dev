package main

import "fmt"

// main generate simple html file to stdout.
// you cant pipe it into the file with redirection ">"
// Example: go run main.go > index.html
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

	fmt.Println(t)
}
