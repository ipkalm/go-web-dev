package main

import (
	"bufio"
	"fmt"
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

		go handle(c)
	}
}

func handle(c net.Conn) {
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
	}
	defer c.Close()
	fmt.Println("Code got here.")
}
