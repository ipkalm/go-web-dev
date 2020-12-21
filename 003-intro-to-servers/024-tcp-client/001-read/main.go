package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	c, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	defer c.Close()

	bs, err := ioutil.ReadAll(c)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
