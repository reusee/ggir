package main

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	funcsOutput io.Writer
}

func Gen(outputDir, girFilePath string) {
	// read gir file contents
	contents, err := ioutil.ReadFile(girFilePath)
	checkError(err)

	// parse
	generator := new(Generator)
	repo := generator.Parse(contents)

	// generate
	ns := repo.Namespace
	goPackageName := strings.ToLower(ns.Name)

	// functions
	funcsOutput := new(bytes.Buffer)
	generator.funcsOutput = funcsOutput
	w(funcsOutput, "package %s\n\n", goPackageName)
	w(funcsOutput, "/*\n")
	if repo.CInclude != nil {
		w(funcsOutput, "#include <%s>\n", repo.CInclude.Name)
	}
	if repo.Package != nil {
		w(funcsOutput, "#cgo pkg-config: %s\n", repo.Package.Name)
	}
	w(funcsOutput, `#cgo linux CFLAGS: -DLINUX
#ifdef LINUX
	#include <glib-unix.h>
#endif
#include <stdlib.h>
#include <stdio.h>
#include <glib-object.h>
#include <glib/gstdio.h>
`)
	w(funcsOutput, "*/\n")
	w(funcsOutput, "import \"C\"\n\n")
	w(funcsOutput, `import (
	"unsafe"
)
func init() {
	var _ unsafe.Pointer
}
`)
	for _, fn := range ns.Functions {
		generator.GenFunction(fn)
	}
	for typeSpec, funcs := range typeStat {
		p("===fixme=== %s TYPE NOT MAPPED => %s\n", goPackageName, typeSpec)
		p("%s\n\n", strings.Join(funcs, "\n"))
	}
	f, err := os.Create(filepath.Join(outputDir, goPackageName+"_functions.go"))
	formatted, err := format.Source(funcsOutput.Bytes())
	if err != nil {
		f.Write(funcsOutput.Bytes())
		f.Close()
		p("==> %s\n", girFilePath)
		checkError(err)
	}
	checkError(err)
	f.Write(formatted)
	f.Close()
}
