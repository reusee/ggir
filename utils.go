package main

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

var (
	p  = fmt.Printf
	w  = fmt.Fprintf
	fs = fmt.Sprintf
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isGoReservedWord(s string) bool {
	return map[string]bool{
		// key words
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
		// builtin functions
		"append":  true,
		"cap":     true,
		"close":   true,
		"complex": true,
		"copy":    true,
		"delete":  true,
		"imag":    true,
		"len":     true,
		"make":    true,
		"new":     true,
		"panic":   true,
		"print":   true,
		"println": true,
		"real":    true,
		"recover": true,
		// builtin types
		"bool":       true,
		"byte":       true,
		"complex128": true,
		"complex64":  true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"int8":       true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uint8":      true,
		"uintptr":    true,
	}[s]
}

func (self *Generator) formatAndOutput(name string, source []byte) {
	f, err := os.Create(filepath.Join(self.outputDir, self.PackageName+"_"+name+".go"))
	checkError(err)
	formatted, err := format.Source(source)
	if err != nil {
		f.Write(source)
		f.Close()
		checkError(err)
	}
	f.Write(formatted)
	f.Close()
}

func replace(t string, m map[string]string) string {
	for key, value := range m {
		t = strings.Replace(t, fs("$$%s", key), value, -1)
	}
	return t
}
