package main

import (
	"fmt"
	"net"
)

func main() {
	c, err := net.Dial("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	defer c.Close()

	fmt.Fprintln(c, "client 003-024-002 :: I write to server")
}
