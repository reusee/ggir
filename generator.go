package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Generator struct {
	outputDir string

	PackageName            string    `xml:"package-name"`
	GirPath                string    `xml:"gir-path"`
	Includes               []string  `xml:"include"`
	PkgConfigs             []string  `xml:"pkg-config"`
	FunctionIgnorePatterns []string  `xml:"function-ignore-patterns>entry"`
	FunctionRenames        []*Rename `xml:"function-rename>rename"`
	ConstantIgnorePatterns []string  `xml:"constant-ignore-patterns>entry"`
	TypesIgnorePatterns    []string  `xml:"type-ignore-patterns>entry"`
}

type Rename struct {
	From string `xml:"from,attr"`
	To   string `xml:",chardata"`
}

func Gen(outputDir, buildFilePath string) {
	// read build file
	buildFileContent, err := ioutil.ReadFile(buildFilePath)
	checkError(err)
	generator := new(Generator)
	err = xml.Unmarshal(buildFileContent, generator)
	checkError(err)
	generator.outputDir = outputDir

	// read gir file contents
	contents, err := ioutil.ReadFile(generator.GirPath)
	checkError(err)

	// parse
	repo := generator.Parse(contents)

	// namespace
	ns := repo.Namespace

	// struct types
	generator.GenStructTypes(ns)

	// class types
	generator.GenClassTypes(ns)

	// functions
	generator.GenFunctions(ns)

	// constants
	generator.GenConstants(ns)

	// flags
	generator.GenFlags(ns)

	// enums
	generator.GenEnums(ns)

	// structs
	generator.GenStructs(ns)

	// class traits
	generator.GenTraits(ns)

	// classes
	generator.GenClasses(ns)
}
