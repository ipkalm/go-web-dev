package main

import (
	"bufio"
	"fmt"
	"io"
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
			continue
		}

		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	var count int8
	var method, uri string
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		ln := sc.Text()
		if count == 0 {
			statusLine := strings.Fields(ln)
			method = statusLine[0]
			uri = statusLine[1]
			count++
		}
		// fmt.Println(ln)
		if ln == "" {
			log.Printf("end of http req headers")
			break
		}
	}
	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}

	var content string
	switch uri {
	case "/":
		if method == "GET" {
			content = "INDEX<br><ul><li><a href=\"#\">index</a></li><li><a href=\"/apply\">apply</a></li></ul>Method: " + method + "<br>URI: " + uri + "<br>"
		}
	case "/apply":
		switch method {
		case "POST":
			content = "APPLY<br><ul><li><a href=\"/\">index</a></li><li><a href=\"/apply\">apply</a></li></ul>Method: " + method + "<br>URI: " + uri + "<br>"
		case "GET":
			content = "APPLY<br><ul><li><a href=\"/\">index</a></li><li><a href=\"#\">apply</a></li></ul>Method: " + method + "<br>URI: " + uri + `<br><form method="POST" action="/apply">
<input type="text" name="fname" id="fname" placeholder="fname" autofocus>
<input type="submit" value="apply">
</form>`
		}
	}

	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<body>
` + content + `
</body>
</html>`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
