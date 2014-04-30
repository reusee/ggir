package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
)

func (self *Generator) GenTraits(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")

	// virtual methods methods functions
	// classes
	for _, c := range ns.Classes {
		goType := cTypeToGoType(c.CType)
		w(output, "type _Trait%s struct { CPointer *%s }\n", c.Name, goType)
		//p("%v\n", c.VirtualMethods)
		//p("%v\n", c.Methods)
		//p("%v\n", c.Functions)

		w(output, "\n")
	}
	// interfaces

	f, err := os.Create(filepath.Join(self.outputDir, self.PackageName+"_traits.go"))
	checkError(err)
	formatted, err := format.Source(output.Bytes())
	if err != nil {
		f.Write(output.Bytes())
		f.Close()
		checkError(err)
	}
	f.Write(formatted)
	f.Close()
}
