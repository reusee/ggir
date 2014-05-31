package main

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

type Generator struct {
	outputDir string

	PackageName             string             `xml:"package-name"`
	GirPath                 string             `xml:"gir-path"`
	ExternalPackages        []*ExternalPackage `xml:"external-package"`
	Includes                []string           `xml:"include"`
	PkgConfigs              []string           `xml:"pkg-config"`
	FunctionIgnorePatterns  []string           `xml:"function-ignore-patterns>entry"`
	FunctionRenames         []*Rename          `xml:"function-rename>rename"`
	ConstantIgnorePatterns  []string           `xml:"constant-ignore-patterns>entry"`
	TypesIgnorePatterns     []string           `xml:"type-ignore-patterns>entry"`
	ParamDirectionEntries   []*Entry           `xml:"param-direction>entry"`
	ParamDirections         map[string]string
	TypeMappingEntries      []*Entry `xml:"type-mapping>entry"`
	InParamMarshalEntries   []*Entry `xml:"in-param-marshal>entry"`
	OutParamMarshalEntries  []*Entry `xml:"out-param-marshal>entry"`
	InParamMarshals         map[string]*Entry
	OutParamMarshals        map[string]*Entry
	SignalParamTypeMappings map[string]string
}

type ExternalPackage struct {
	Name    string `xml:"name,attr"`
	GirPath string `xml:"gir-path"`
	Import  string `xml:"import"`
	Repo    *Repository
}

type Rename struct {
	From string `xml:"from,attr"`
	To   string `xml:",chardata"`
}

type Entry struct {
	Text   string `xml:",chardata"`
	Spec   string `xml:"spec,attr"`
	Param  string `xml:"param"`
	Before string `xml:"before"`
	After  string `xml:"after"`
}

func Gen(outputDir string) {
	// read build file
	buildFileContent, err := ioutil.ReadFile(filepath.Join(outputDir, "build.xml"))
	checkError(err)
	generator := &Generator{
		ParamDirections:         make(map[string]string),
		InParamMarshals:         make(map[string]*Entry),
		OutParamMarshals:        make(map[string]*Entry),
		SignalParamTypeMappings: make(map[string]string),
	}
	err = xml.Unmarshal(buildFileContent, generator)
	checkError(err)
	generator.outputDir = outputDir

	// prepares
	for _, entry := range generator.ParamDirectionEntries {
		generator.ParamDirections[entry.Spec] = entry.Text
	}
	for _, entry := range generator.TypeMappingEntries {
		TypeMapping[entry.Spec] = entry.Text
	}
	for _, entry := range generator.InParamMarshalEntries {
		generator.InParamMarshals[entry.Spec] = entry
	}
	for _, entry := range generator.OutParamMarshalEntries {
		generator.OutParamMarshals[entry.Spec] = entry
	}

	// parse gir
	contents, err := ioutil.ReadFile(generator.GirPath)
	checkError(err)
	repo := generator.Parse(contents)

	// parse external gir
	for _, ext := range generator.ExternalPackages {
		contents, err := ioutil.ReadFile(ext.GirPath)
		checkError(err)
		ext.Repo = generator.Parse(contents)
	}

	// add includes
	generator.Includes = append(generator.Includes, "stdlib.h")

	// namespace
	ns := repo.Namespace
	p("=> %s\n", ns.Name)

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

	// class constructors
	generator.GenClassConstructors(ns)

	// signals
	// generator.GenSignals(ns)
}
