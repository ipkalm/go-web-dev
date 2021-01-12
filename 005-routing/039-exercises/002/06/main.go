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
		fmt.Println(ln)
		if len(ln) == 0 {
			log.Println("end of http request headers")
			break
		}
	}

	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}
	body := "random body"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
