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
	if s := request(c); s != "" {
		fmt.Println(s)
		response(c, s)
	} else {
		c.Close()
		log.Panic("something wrong in with request func")
	}
}

func request(c net.Conn) string {
	urlReq := ""
	req := make(map[string][]string)

	sc := bufio.NewScanner(c)

	for sc.Scan() {
		l := sc.Text()
		if l != "" {
			t := strings.Fields(l)
			req[t[0]] = t[1:]
		} else {
			break
		}
	}

	if req["GET"] != nil {
		urlReq = req["GET"][0]
	} else {
		log.Panicln("Request method is not GET")
	}

	return urlReq
}

func response(c net.Conn, s string) {
	defer c.Close()

	messageBody := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
</head>
<body>
	<h1>this is your request URL</h1>
	<h2>` + s + `</h2>
</body>
</html>
`

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(c, "Content-Type: text/html; charset=UTF-8\r\n")
	fmt.Fprint(c, "Server: ipkalmws\r\n")
	fmt.Fprintf(c, "ContentLenght: %d\r\n", len(messageBody))
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, messageBody)
}
