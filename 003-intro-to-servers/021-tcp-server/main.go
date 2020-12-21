package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		s := "message for Fprintf"
		io.WriteString(c, "\nHello, Gopher. It is TCP server\n")
		fmt.Fprintln(c, "how are you?")
		fmt.Fprintf(c, "%v, %T\n", s, s)
		c.Close()
	}
}
