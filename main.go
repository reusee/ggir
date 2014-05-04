package main

import (
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		p("usage: %s [output dir]\n", os.Args[0])
		return
	}
	outputDir, err := filepath.Abs(os.Args[1])
	checkError(err)
	_, err = os.Stat(outputDir)
	if err != nil {
		err = os.Mkdir(outputDir, 0755)
		checkError(err)
	}
	Gen(outputDir)
}
