package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
			log.Fatalln(err)
		}
		go handle(c)
	}
}

func handle(c net.Conn) {
	err := c.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("connection timeout")
	}

	s := bufio.NewScanner(c)
	for s.Scan() {
		l := s.Text()
		fmt.Println(l)
		fmt.Fprintf(c, "You said \"%s\"?\n", l)
	}
	defer c.Close()

	fmt.Println("code got here")
}
