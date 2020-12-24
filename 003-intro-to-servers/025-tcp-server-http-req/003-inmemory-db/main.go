package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

var db []string

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
	sc := bufio.NewScanner(c)
	if err := sc.Err(); err != nil {
		log.Panicln(err)
	}

	for sc.Scan() {
		l := strings.Fields(strings.ToLower(sc.Text()))
		if len(l) == 0 {
			log.Println("empty input")
			continue
		}

		switch l[0] {
		case "insert":
			tmpString := strings.Join(l[1:], " ")
			db = append(db, tmpString)
			fmt.Fprintf(c, "\"%v\" successful added\n", tmpString)
		case "show":
			if len(db) == 0 {
				fmt.Fprintln(c, "db is empty")
			}

			for k, v := range db {
				fmt.Fprintf(c, "№ %v: %v\n", k, v)
			}
		case "drop":
			if len(l) != 2 {
				fmt.Fprintln(c, "Incorrect number of arguments.")
				fmt.Fprintln(c, "I want index of element.")
				fmt.Fprintln(c, "drop 2 for example")

				break
			}
			if l[1] == "all" {
				db = []string{}
				fmt.Fprintln(c, "succefully remove all data")

				break
			}

			index, err := strconv.Atoi(l[1])
			if err != nil {
				fmt.Fprintln(c, "need a number, get a ", l[1])
				break
			}
			db = append(db[:index], db[index+1:]...)
			fmt.Fprintln(c, "delete success")
		default:
			log.Println("incorrect command", l[0])
			fmt.Fprintln(c, "error: I understand only this commands: insert, show, drop")
		}
	}
}
