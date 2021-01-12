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
		if ln == "" {
			log.Printf("end of http req headers")
			break
		}
	}
	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}

	body := "Method:\t" + method + "\nURI:\t" + uri + "\n"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
