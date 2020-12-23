package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	req := getRequest(c)
	response(c, req)
}

func getRequest(c net.Conn) map[string]string {
	req := make(map[string]string)
	sc := bufio.NewScanner(c)

	i := 0
	for sc.Scan() {
		l := sc.Text()
		if i == 0 {
			t := strings.Fields(l)
			req["METHOD"] = t[0]
			req["URI"] = t[1]
			req["HTTP_VER"] = t[2]
			i++
		} else {
			break
		}
	}

	return req
}

func response(c net.Conn, req map[string]string) {
	var innerHTML string
	var codeAndReason string
	var code uint16 = 200

	switch req["URI"] {
	case "/": //index page
		innerHTML = `<nav>
			<ul>
			<li><a href="#">index</a></li>
			<li><a href="/blog/">blog</a></li>
			<li><a href="/about/">about</a></li>
			<li><a href="/signin/">sign in</a></li>
			</ul>
			</nav>
			<h1>INDEX PAGE</h1>`
	case "/about/":
		innerHTML = `<nav>
			<ul>
			<li><a href="/">index</a></li>
			<li><a href="/blog/">blog</a></li>
			<li><a href="#">about</a></li>
			<li><a href="/signin/">sign in</a></li>
			</ul>
			</nav>
			<h1>ABOUT PAGE</h1>`
	case "/blog/":
		innerHTML = `<nav>
			<ul>
			<li><a href="/">index</a></li>
			<li><a href="#">blog</a></li>
			<li><a href="/about/">about</a></li>
			<li><a href="/signin/">sign in</a></li>
			</ul>
			</nav>
			<h1>BLOG PAGE</h1>`
	case "/signin/":
		if req["METHOD"] == "POST" {
			innerHTML = `<nav>
				<ul>
				<li><a href="/">index</a></li>
				<li><a href="/blog/">blog</a></li>
				<li><a href="/about/">about</a></li>
				<li><a href="/signin/">sign in</a></li>
				</ul>
				</nav>
				<h1>THANK YOU FOR REGISTRATION</h1>`
		} else {
			innerHTML = `<nav>
				<ul>
				<li><a href="/">index</a></li>
				<li><a href="#">blog</a></li>
				<li><a href="/about/">about</a></li>
				<li><a href="/signin/">sign in</a></li>
				</ul>
				</nav>
				<h1>SIGN IN</h1>
				<form method="post">
				<label for="name">name</label><br>
				<input type="text" id="name" name="name"><br>
				<input type="submit" value="submit">
				</form>`
		}
	default:
		if strings.HasPrefix(req["URI"], "/signin/?name=") {
			innerHTML = req["URI"]
			break
		}
		innerHTML = `PAGE NOT FOUND`
		code = 404
	}

	messageBody := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
		</head>
		<body>
		` + innerHTML + `<br><br>
		` + fmt.Sprint(req) + `
		</body>
		</html>
		`

	switch code {
	case 200:
		codeAndReason = "200 OK\r\n"
	case 404:
		codeAndReason = "404 Not Found\r\n"
	}

	fmt.Fprintf(c, "%v %v", req["HTTP_VER"], codeAndReason)
	fmt.Fprint(c, "Content-Type: text/html; charset=UTF-8\r\n")
	fmt.Fprint(c, "Server: ipkalmws\r\n")
	fmt.Fprintf(c, "ContentLenght: %d\r\n", len(messageBody))
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, messageBody)
}
