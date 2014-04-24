package main

import (
	"fmt"
	"log"
)

var (
	p = fmt.Printf
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
