package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := " Lorem ipsum dolor sit amet consectetur, adipisicing elit. Dicta recusandae voluptates reiciendis ipsum ratione quos ad nihil ullam incidunt illum officia necessitatibus accusantium exercitationem porro, suscipit sequi, ab itaque inventore?"

	encodeStd := "ABCDEFGHIJKLMOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	s64std := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s), s)
	fmt.Println(len(s64), s64)
	fmt.Println(len(s64std), s64std)

	bs64, err := base64.NewEncoding(encodeStd).DecodeString(s64)
	if err != nil {
		log.Panicln(err)
	}

	bs64std, err := base64.StdEncoding.DecodeString(s64std)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(len(bs64), string(bs64))
	fmt.Println(len(bs64std), string(bs64std))
}
