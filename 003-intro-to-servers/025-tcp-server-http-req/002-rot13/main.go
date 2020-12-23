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
		go handler(c)
	}
}

func handler(c net.Conn) {
	defer c.Close()

	sc := bufio.NewScanner(c)
	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}

	for sc.Scan() {
		tmp := []rune{}
		l := strings.ToLower(sc.Text())

		for _, v := range l {
			if v >= 97 && v <= 109 {
				v += 13
			} else if v >= 110 && v <= 122 {
				v -= 13
			} else {
				log.Println("Not a text characters", v)
			}

			tmp = append(tmp, v)
		}

		fmt.Fprintf(c, "%v\n", string(tmp))
	}
}
