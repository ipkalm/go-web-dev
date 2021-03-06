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

		go func(conn net.Conn) {
			defer conn.Close()

			sc := bufio.NewScanner(conn)

			for sc.Scan() {
				ln := sc.Text()
				if len(ln) == 0 {
					log.Println("end of http request headers")
					break
				}
				fmt.Println(ln)
			}
			if err := sc.Err(); err != nil {
				log.Panicln(err)
			}

			log.Println("Code got here.")

			_, err := io.WriteString(c, "I see you connected")
			if err != nil {
				log.Panicln(err)
			}
		}(c)
	}
}
