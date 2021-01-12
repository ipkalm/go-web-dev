package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		ln := sc.Text()
		if ln == "" {
			log.Printf("end of http req headers")
			break
		}
	}
	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}

	body := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
	<h1>HOLY COW THIS IS LOW LEVEL</h1>
</body>
</html>`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
