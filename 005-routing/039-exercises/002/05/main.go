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

	io.WriteString(c, "hello there")
}
