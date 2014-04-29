package main

import (
	"bytes"
	"encoding/xml"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	funcsOutput io.Writer
	outputDir   string

	PackageName            string   `xml:"package-name"`
	GirPath                string   `xml:"gir-path"`
	Includes               []string `xml:"include"`
	PkgConfigs             []string `xml:"pkg-config"`
	FunctionIgnorePatterns []string `xml:"function-ignore-patterns>entry"`
	FunctionDeprecated     []string `xml:"function-deprecated>entry"`
	ConstantIgnorePatterns []string `xml:"constant-ignore-patterns>entry"`
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

	// functions
	funcsOutput := new(bytes.Buffer)
	generator.funcsOutput = funcsOutput
	w(funcsOutput, "package %s\n\n", generator.PackageName)
	w(funcsOutput, "/*\n")
	for _, include := range generator.Includes {
		w(funcsOutput, "#include <%s>\n", include)
	}
	if len(generator.PkgConfigs) > 0 {
		w(funcsOutput, "#cgo pkg-config: %s\n", strings.Join(generator.PkgConfigs, " "))
	}
	w(funcsOutput, "*/\n")
	w(funcsOutput, "import \"C\"\n\n")
	w(funcsOutput, `import ("unsafe";	"reflect"; "errors")
func init() {
	var _ unsafe.Pointer
	var _ reflect.SliceHeader
	_ = errors.New("")
}
`)
	for _, fn := range ns.Functions {
		generator.GenFunction(fn)
	}
	for typeSpec, funcs := range typeStat {
		p("===fixme=== %s TYPE NOT MAPPED => %s\n", generator.PackageName, typeSpec)
		p("%s\n\n", strings.Join(funcs, "\n"))
	}
	f, err := os.Create(filepath.Join(outputDir, generator.PackageName+"_functions.go"))
	checkError(err)
	formatted, err := format.Source(funcsOutput.Bytes())
	if err != nil {
		f.Write(funcsOutput.Bytes())
		f.Close()
		p("==> %s\n", generator.GirPath)
		checkError(err)
	}
	f.Write(formatted)
	f.Close()

	// constants
	generator.GenConstants(ns)

	// flags
	generator.GenFlags(ns)

	// enums
	generator.GenEnums(ns)
}
