package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func (self *Generator) GenFlags(ns *Namespace) {
	output := new(bytes.Buffer)
	w(output, "package %s\n\n", self.PackageName)
	w(output, "/*\n")
	for _, include := range self.Includes {
		w(output, "#include <%s>\n", include)
	}
	w(output, "*/\n")
	w(output, "import \"C\"\n")

	w(output, "var (\n")
	for _, f := range ns.Bitfields {
		w(output, "\t// %s\n", f.Name)
		for _, m := range f.Members {
			parts := strings.Split(m.CIdentifier, "_")
			goName := strings.Join(parts[1:], "_")
			w(output, "%s = C.%s\n", goName, m.CIdentifier)
		}
		w(output, "\n")
	}
	w(output, ")\n")

	f, err := os.Create(filepath.Join(self.outputDir, self.PackageName+"_flags.go"))
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
