package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		p("usage: %s [output dir] [gir file path]\n", os.Args[0])
		return
	}
	outputDir, err := filepath.Abs(os.Args[1])
	checkError(err)
	_, err = os.Stat(outputDir)
	if err != nil {
		err = os.Mkdir(outputDir, 0755)
		checkError(err)
	}
	girFilePath, err := filepath.Abs(os.Args[2])
	checkError(err)

	Gen(outputDir, girFilePath)
}

type Generator struct {
}

func Gen(outputDir, girFilePath string) {
	// read gir file contents
	contents, err := ioutil.ReadFile(girFilePath)
	checkError(err)

	generator := new(Generator)
	repo := generator.Parse(contents)
	_ = repo
}
