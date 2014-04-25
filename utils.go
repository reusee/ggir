package main

import (
	"fmt"
	"log"
)

var (
	p = fmt.Printf
	w = fmt.Fprintf
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//TODO builtins?
func isGoReservedWord(s string) bool {
	return map[string]bool{
		"break":       true,
		"default":     true,
		"func":        true,
		"interface":   true,
		"select":      true,
		"case":        true,
		"defer":       true,
		"go":          true,
		"map":         true,
		"struct":      true,
		"chan":        true,
		"else":        true,
		"goto":        true,
		"package":     true,
		"switch":      true,
		"const":       true,
		"fallthrough": true,
		"if":          true,
		"range":       true,
		"type":        true,
		"continue":    true,
		"for":         true,
		"import":      true,
		"return":      true,
		"var":         true,
	}[s]
}
