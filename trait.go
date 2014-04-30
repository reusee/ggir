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
	w(output, "import \"unsafe\"\n")
	w(output, "import \"reflect\"\n")
	w(output, "import \"errors\"\n")
	w(output, `func init() {
		_ = unsafe.Pointer(nil)
		_ = reflect.ValueOf(nil)
		_ = errors.New("")
	}
	`)

	// classes
	for _, c := range ns.Classes {
		goType := cTypeToGoType(c.CType)
		w(output, "type _Trait%s struct { CPointer *%s }\n", c.Name, goType)
		for _, m := range c.Methods {
			self.GenFunction(m, output, c)
		}
		//p("%v\n", c.VirtualMethods) FIXME
		//p("%v\n", c.Methods) FIXME
		//p("%v\n", c.Functions) FIXME

		w(output, "\n")
	}
	// interfaces FIXME
	//p("%v\n", c.VirtualMethods) FIXME
	//p("%v\n", c.Methods) FIXME
	//p("%v\n", c.Functions) FIXME

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
