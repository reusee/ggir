package main

import "io/ioutil"

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
